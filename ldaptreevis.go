package ldaptreevis

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"sort"
	"strings"
)

type collection struct {
	nodes []*Node
}

func (c *collection) addNodeIfNotExist(node *Node) {
	_, found := c.findNodeExactLineage(node.Value, node.Lineage, node.Depth)
	if !found {
		if len(c.nodes) == 0 {
			c.nodes = []*Node{}
		}
		c.nodes = append(c.nodes, node)
	}
}

func (c *collection) newNode(value string, depth int) *Node {
	node := &Node{
		Value: value,
		Depth: depth,
	}
	node.Uid = uuid.New()
	return node
}

func (c *collection) findNode(value string) (result *Node, ok bool) {
	for _, node := range c.nodes {
		if node.Value == value {
			result = node
			ok = true
		}
	}
	return result, ok
}

func (c *collection) findNodeExact(value string, depth int) (result *Node, ok bool) {
	for _, node := range c.nodes {
		if (node.Value == value) && (node.Depth == depth) {
			result = node
			ok = true
		}
	}
	return result, ok
}

func (c *collection) findNodeExactLineage(value, lineage string, depth int) (result *Node, ok bool) {
	for _, node := range c.nodes {
		if node.Value == value && node.Depth == depth {
			if node.Lineage == lineage {
				result = node
				ok = true
			}
		}
	}
	return result, ok
}

// Node contains properties and methods to represent an object in the LDAP
// tree and to alter properties such as children, parent, value, etc.
type Node struct {
	Value     string    `json:"label"`
	Class     string    `json:"class"`
	Children  []*Node   `json:"children,omitempty"`
	Parent    *Node     `json:"-"`
	ParentUid string    `json:"parentUid"`
	Uid       uuid.UUID `json:"uid"`
	Depth     int       `json:"depth"`
	Lineage   string    `json:"lineage"`
}

// FmtTree returns a string formatted as a multiline tree representing
// thise Node and it's children.
func (n *Node) FmtTree(start string) string {
	padding := strings.Repeat("  ", n.Depth)
	start += fmt.Sprintf("%s%s\n", padding, n.Value)
	for _, node := range n.Children {
		start = node.FmtTree(start)
	}
	return start
}

// HasChild searches the Node's children for a child with 
// the requested value. 
func (n *Node) HasChild(value string) (found bool) {
	for _, child := range n.Children {
		if child.Value == value {
			found = true
		}
	}
	return found
}

// AddChild adds a Node to this node's children and 
// updates the provided Node's parent property as well.
func (n *Node) AddChild(child *Node) {
	if len(n.Children) == 0 {
		n.Children = []*Node{}
	}
	n.Children = append(n.Children, child)
	child.AddParent(n)
}

// AddParent updates the parent of the current node
// to the node provided in the argument. Somewhat safe.
func (n *Node) AddParent(parent *Node) {
	if n.Uid == parent.Uid {
		return
	}
	n.Parent = parent
	n.ParentUid = parent.Uid.String()
	if !parent.HasChild(n.Value) {
		parent.AddChild(n)
	}
}

// BuildTree takes a slice of LDAP Distinguished Name strings
// and attempts to build a node tree that represents all of their
// relationships (if any) under a generic parent "root" node. 
//
// It will return the root node object and the visualization string
// and any errors.
func BuildTree(input []string) (root *Node, vis string, err error) {
	// first build a map of the strings so they can be processed
	// in an ordered map
	var results []map[int]pair
	for _, i := range input {
		result, err := parseDN(i)
		if err != nil {
			return root, vis, err
		}
		results = append(results, result)
	}
	// start building a tree. Start by building a collection
	// that we can use to cache unique unstructured results for the purposes
	// of searching, etc.
	col := collection{}
	// build the initial root node
	root = col.newNode("root", 0)
	seed := "root "
	root.Lineage = seed
	root.Class = "root"
	col.addNodeIfNotExist(root)
	for _, bore := range results {
		// first order the map
		keys := make([]int, 0)
		for k, _ := range bore {
			keys = append(keys, k)
		}
		// sort reverse so we're processing from parents to children
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		// loop through keys in order
		for i, level := range keys {
			unit := bore[level].value
			cur := col.newNode(unit, i+1)
			cur.Class = bore[level].class
			// we know our parent is always adjacent in the slice and if
			// it comes up blank we're a child of root
			parentName := bore[level+1].value
			if parentName == "" {
				parentName = strings.TrimSpace(seed)
			}
			// we'll construct the known parent depth, known parent lineage, and known
			// parent name so we can find the correct parent in the collection
			expectedParentDepth := i
			// both the parent and the current node's linage is contained
			// within the LDAP string so we'll go ahead and construct
			// for both
			parentLineage := ""
			for j := level + 1; j < len(bore); j++ {
				parentLineage = fmt.Sprintf("%s ", bore[j].value) + parentLineage
			}
			// have to add the known root string to the lineage as it's
			// not contained in the LDAP string
			parentLineage = seed + parentLineage
			lineage := ""
			for j := level; j < len(bore); j++ {
				lineage = fmt.Sprintf("%s ", bore[j].value) + lineage
			}
			lineage = seed + lineage
			cur.Lineage = lineage
			// now we have to find the parent Node object in the collection
			// in the case of a child of root it will always already exist
			// as we explicitly created it. 
			if parentNode, ok := col.findNodeExactLineage(
				parentName,
				parentLineage,
				expectedParentDepth,
			); ok {
				// once the parent is found we can make sure we're not
				// already an existing child (prevents duplication of 
				// repeat nodes)
				if !parentNode.HasChild(cur.Value) {
					// Then we add ourselves
					// to the collection (if we don't exist) and inform
					// the parent it has a new child
					col.addNodeIfNotExist(cur)
					parentNode.AddChild(cur)
				}
			}
		}
	}
	vis = root.FmtTree("")
	return root, vis, err
}

type pair struct {
	class string
	value string
}

func parseDN(input string) (out map[int]pair, err error) {
	out = make(map[int]pair)
	chunks := strings.Split(input, ",")
	if len(chunks) > 0 {
		for i, s := range chunks {
			lchunks := strings.Split(s, "=")
			if len(lchunks) == 2 {
				p := pair{
					class: lchunks[0],
					value: lchunks[1],
				}
				out[i] = p
			}
		}
	} else {
		err = errors.New("non csv input")
	}
	return out, err
}
