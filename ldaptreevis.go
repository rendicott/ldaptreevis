package ldaptreevis

import (
	"fmt"
	"strings"
	"errors"
	"sort"
	"github.com/google/uuid"
)

type Collection struct {
	Nodes []*Node
}

func (c *Collection) AddNodeIfNotExist(node *Node) () {
	_, found := c.FindNodeExactLineage(node.Value, node.Lineage, node.Depth)
	if !found {
		if len(c.Nodes) == 0 {
			c.Nodes = []*Node{}
		}
		c.Nodes = append(c.Nodes, node)
	}
}

func (c *Collection) NewNode (value string, depth int) (*Node) {
	node := &Node{
		Value: value,
		Depth: depth,
	}
	node.Uid = uuid.New()
	return node
}

func (c *Collection) FindNode(value string) (result *Node, ok bool) {
	for _, node := range(c.Nodes) {
		if node.Value == value {
			result = node
			ok = true
		}
	}
	return result, ok
}

func (c *Collection) FindNodeExact(value string, depth int) (result *Node, ok bool) {
	for _, node := range(c.Nodes) {
		if (node.Value == value) && (node.Depth == depth) {
			result = node
			ok = true
		}
	}
	return result, ok
}

func (c *Collection) FindNodeExactLineage(value, lineage string, depth int) (result *Node, ok bool) {
	for _, node := range(c.Nodes) {
		if (node.Value == value && node.Depth == depth) {
			if (node.Lineage == lineage) {
				result = node
				ok = true
			}
		}
	}
	return result, ok
}

type Node struct {
	Value string `json:"label"`
	Children []*Node `json:"children,omitempty"`
	Parent *Node `json:"-"`
	Uid uuid.UUID `json:"-"`
	Depth int `json:"depth"`
	Lineage string `json:"-"`
}

func (n *Node) Dump(start string) (string) {
	padding := strings.Repeat("  ", n.Depth)
	start += fmt.Sprintf("%s%s\n", padding, n.Value)
	for _, node := range(n.Children) {
		start = node.Dump(start)
	}
	return start
}

func (n *Node) HasChild(value string) (found bool) {
	for _, child := range(n.Children) {
		if child.Value == value {
			found = true
		}
	}
	return found
}

func (n *Node) AddChild(child *Node) {
	if len(n.Children) == 0 {
		n.Children = []*Node{}
	}
	child.AddParent(n)
	n.Children = append(n.Children, child)
}

func (n *Node) AddParent(parent *Node) {
	n.Parent = parent
}

func ParseDNs(input []string) (root *Node, vis string, err error) {
	maxDepth := 0
	var results []map[int]string
	for _, i := range(input) {
		result, err := ParseDN(i)
		if err != nil {
			return root, vis, err
		}
		if len(result) > maxDepth {
			maxDepth = len(result)
		}
		results = append(results, result)
	}
	// start building a tree
	col := Collection{}
	root = col.NewNode("root", 0)
	root.Lineage = "root "
	col.AddNodeIfNotExist(root)
	for _, bore := range(results) {
		// first order the map
		keys := make([]int, 0)
		for k, _ := range bore{
			keys = append(keys, k)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		// loop through keys in order
		for i, level := range(keys) {
			unit := bore[level]
			cur := col.NewNode(unit, i+1)
			parentName := bore[level + 1]
			expectedParentDepth := i
			lineage := ""
			parentLineage := ""
			for j := level+1; j < len(bore); j++ {
				parentLineage = fmt.Sprintf("%s ",bore[j]) + parentLineage
			}
			for j := level; j < len(bore); j++ {
				lineage = fmt.Sprintf("%s ",bore[j]) + lineage
			}
			if parentName == "" {
				parentName = "root"
			}
			lineage = "root " + lineage
			parentLineage = "root " + parentLineage
			cur.Lineage = lineage
			if parentNode, ok := col.FindNodeExactLineage(
					parentName,
					parentLineage,
					expectedParentDepth,
					); ok {
				if !parentNode.HasChild(cur.Value) {
					col.AddNodeIfNotExist(cur)
					parentNode.AddChild(cur)
				}
			}
		}
	}
	vis = root.Dump("")
	return root, vis, err
}

func ParseDN(input string) (out map[int]string, err error) {
	out = make(map[int]string)
	chunks := strings.Split(input, ",")
	if len(chunks) > 0 {
		for i, s := range(chunks) {
			lchunks := strings.Split(s, "=")
			if len(lchunks) == 2 {
				out[i] = lchunks[1]
			}
		}
	} else {
		err = errors.New("non csv input")
	}
	return out, err
}
