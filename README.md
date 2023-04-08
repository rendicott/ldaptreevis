# ldaptreevis
Structures LDAP Distinguished Name strings into a visualized tree. 

[docs](#docs)

For example, providing a slice of strings like
```
"CN=SPSAdmins,OU=Groups,OU=MYTOWN,OU=Germany,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=FooAdmin,OU=Groups,OU=MYTOWN,OU=Germany,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=BarAdmin,OU=Groups,OU=MYTOWN,OU=Germany,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=SPSAdmins,OU=Groups,OU=MYTOWN,OU=America,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=FooAdmin,OU=Groups,OU=MYTOWN,OU=America,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=BarAdmin,OU=Groups,OU=MYTOWN,OU=America,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=SPSAdmins,OU=Groups,OU=YourTown,OU=America,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=FooAdmin,OU=Groups,OU=YourTown,OU=America,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=BarAdmin,OU=Groups,OU=YourTown,OU=America,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=LongAdmin,OU=Bob,OU=Doug,OU=gone,OU=Groups,OU=YourTown,OU=America,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
"CN=ShortAdmin,OU=Groups,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
```

You would get
```
root
  com
    MYCOMPANY
      MYTOWN
        MYCOMPANY
          Germany
            MYTOWN
              Groups
                SPSAdmins
                FooAdmin
                BarAdmin
          America
            MYTOWN
              Groups
                SPSAdmins
                FooAdmin
                BarAdmin
            YourTown
              Groups
                SPSAdmins
                FooAdmin
                BarAdmin
                gone
                  Doug
                    Bob
                      LongAdmin
          Groups
            ShortAdmin
```

See example usage in `cmd/main.go` of how to read DN's from STDIN and display a tree.

```
# https://www.forumsys.com/2022/05/10/online-ldap-test-server/

$ ldapsearch -x -H ldap://ldap.forumsys.com -D "cn=read-only-admin,dc=example,dc=com" -w password -b "dc=example,dc=com" | grep "dn:" | cut -d ' ' -f 2 | go run main.go
root
  com
    example
      admin
      newton
      einstein
      tesla
      galieleo
      euler
      gauss
      riemann
      euclid
      mathematicians
      scientists
        italians
      read-only-admin
      test
      chemists
      curie
      nobel
      boyle
      pasteur
      nogroup
```

You can also use the returned root node to do whatever you want, e.g., export to JSON

```
{
  "label": "root",
  "class": "root",
  "children": [
    {
      "label": "com",
      "class": "DC",
      "children": [
        {
          "label": "MYCOMPANY",
          "class": "DC",
          "children": [
            {
              "label": "MYTOWN",
              "class": "DC",
              "children": [
                {
                  "label": "MYCOMPANY",
                  "class": "OU",
                  "children": [
                    {
                      "label": "Germany",
                      "class": "OU",
                      "children": [
                        {
                          "label": "MYTOWN",
                          "class": "OU",
                          "children": [
                            {
                              "label": "Groups",
                              "class": "OU",
                              "children": [
                                {
                                  "label": "SPSAdmins",
                                  "class": "CN",
                                  "parentUid": "879a3643-276f-43e2-92c7-6a230f6a7edf",
                                  "uid": "7c650c2a-d70e-4fd8-b2b2-6062036b4fc9",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Germany MYTOWN Groups SPSAdmins "
                                },
                                {
                                  "label": "FooAdmin",
                                  "class": "CN",
                                  "parentUid": "879a3643-276f-43e2-92c7-6a230f6a7edf",
                                  "uid": "a39c4145-a01e-48de-a04f-a10f0119957e",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Germany MYTOWN Groups FooAdmin "
                                },
                                {
                                  "label": "BarAdmin",
                                  "class": "CN",
                                  "parentUid": "879a3643-276f-43e2-92c7-6a230f6a7edf",
                                  "uid": "6f546998-b67c-4777-a3c1-503e00f37536",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Germany MYTOWN Groups BarAdmin "
                                }
                              ],
                              "parentUid": "7b0482a6-ac19-400a-b807-e1f8c673c8d0",
                              "uid": "879a3643-276f-43e2-92c7-6a230f6a7edf",
                              "depth": 7,
                              "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Germany MYTOWN Groups "
                            }
                          ],
                          "parentUid": "efa2fd31-f714-44b2-94fd-f7fa547c4cc9",
                          "uid": "7b0482a6-ac19-400a-b807-e1f8c673c8d0",
                          "depth": 6,
                          "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Germany MYTOWN "
                        }
                      ],
                      "parentUid": "d75a6a1c-fed5-4876-96fd-d463d2d3636d",
                      "uid": "efa2fd31-f714-44b2-94fd-f7fa547c4cc9",
                      "depth": 5,
                      "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Germany "
                    },
                    {
                      "label": "America",
                      "class": "OU",
                      "children": [
                        {
                          "label": "MYTOWN",
                          "class": "OU",
                          "children": [
                            {
                              "label": "Groups",
                              "class": "OU",
                              "children": [
                                {
                                  "label": "SPSAdmins",
                                  "class": "CN",
                                  "parentUid": "90aec12d-bbb9-42ae-b880-4a131c2ed6c7",
                                  "uid": "eb17a562-1a2a-4508-9ff7-4d21a943330a",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America MYTOWN Groups SPSAdmins "
                                },
                                {
                                  "label": "FooAdmin",
                                  "class": "CN",
                                  "parentUid": "90aec12d-bbb9-42ae-b880-4a131c2ed6c7",
                                  "uid": "42a35d7a-0309-4dd9-80df-ccb1e73fb077",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America MYTOWN Groups FooAdmin "
                                },
                                {
                                  "label": "BarAdmin",
                                  "class": "CN",
                                  "parentUid": "90aec12d-bbb9-42ae-b880-4a131c2ed6c7",
                                  "uid": "f0d9ca18-1c9f-4c73-82e6-c91bc05241bf",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America MYTOWN Groups BarAdmin "
                                }
                              ],
                              "parentUid": "64f2d3b6-b2f6-4e44-b3e5-6453d1247ca0",
                              "uid": "90aec12d-bbb9-42ae-b880-4a131c2ed6c7",
                              "depth": 7,
                              "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America MYTOWN Groups "
                            }
                          ],
                          "parentUid": "f1639517-0471-492f-8af3-e6d8698115c6",
                          "uid": "64f2d3b6-b2f6-4e44-b3e5-6453d1247ca0",
                          "depth": 6,
                          "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America MYTOWN "
                        },
                        {
                          "label": "YourTown",
                          "class": "OU",
                          "children": [
                            {
                              "label": "Groups",
                              "class": "OU",
                              "children": [
                                {
                                  "label": "SPSAdmins",
                                  "class": "CN",
                                  "parentUid": "11c64711-348d-44f1-bf8f-9fb26f488eb1",
                                  "uid": "f691c718-83ee-4f96-ae39-4f2c71fca20f",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups SPSAdmins "
                                },
                                {
                                  "label": "FooAdmin",
                                  "class": "CN",
                                  "parentUid": "11c64711-348d-44f1-bf8f-9fb26f488eb1",
                                  "uid": "2d83033a-d39a-47a5-994d-ad066e260706",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups FooAdmin "
                                },
                                {
                                  "label": "BarAdmin",
                                  "class": "CN",
                                  "parentUid": "11c64711-348d-44f1-bf8f-9fb26f488eb1",
                                  "uid": "d021aff9-c065-45d1-b7ee-5fdb05256033",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups BarAdmin "
                                },
                                {
                                  "label": "gone",
                                  "class": "OU",
                                  "children": [
                                    {
                                      "label": "Doug",
                                      "class": "OU",
                                      "children": [
                                        {
                                          "label": "Bob",
                                          "class": "OU",
                                          "children": [
                                            {
                                              "label": "LongAdmin",
                                              "class": "CN",
                                              "parentUid": "f9bd61c3-eef4-4305-b122-67e4e5041951",
                                              "uid": "750d6999-2da3-4df7-8c2b-d43f2a7931c7",
                                              "depth": 11,
                                              "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups gone Doug Bob LongAdmin "
                                            }
                                          ],
                                          "parentUid": "b093805f-7b10-4f55-b63d-680348d4faed",
                                          "uid": "f9bd61c3-eef4-4305-b122-67e4e5041951",
                                          "depth": 10,
                                          "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups gone Doug Bob "
                                        }
                                      ],
                                      "parentUid": "325149f6-c917-4398-b67d-9d0811b8cb7a",
                                      "uid": "b093805f-7b10-4f55-b63d-680348d4faed",
                                      "depth": 9,
                                      "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups gone Doug "
                                    }
                                  ],
                                  "parentUid": "11c64711-348d-44f1-bf8f-9fb26f488eb1",
                                  "uid": "325149f6-c917-4398-b67d-9d0811b8cb7a",
                                  "depth": 8,
                                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups gone "
                                }
                              ],
                              "parentUid": "c28058a2-781e-4ca6-a7d0-0a66db180cc7",
                              "uid": "11c64711-348d-44f1-bf8f-9fb26f488eb1",
                              "depth": 7,
                              "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown Groups "
                            }
                          ],
                          "parentUid": "f1639517-0471-492f-8af3-e6d8698115c6",
                          "uid": "c28058a2-781e-4ca6-a7d0-0a66db180cc7",
                          "depth": 6,
                          "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America YourTown "
                        }
                      ],
                      "parentUid": "d75a6a1c-fed5-4876-96fd-d463d2d3636d",
                      "uid": "f1639517-0471-492f-8af3-e6d8698115c6",
                      "depth": 5,
                      "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY America "
                    },
                    {
                      "label": "Groups",
                      "class": "OU",
                      "children": [
                        {
                          "label": "ShortAdmin",
                          "class": "CN",
                          "parentUid": "16c15b19-8518-4e51-aabf-feffe030e90f",
                          "uid": "9e767c1d-9e96-4feb-a4fb-9069dc74396b",
                          "depth": 6,
                          "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Groups ShortAdmin "
                        }
                      ],
                      "parentUid": "d75a6a1c-fed5-4876-96fd-d463d2d3636d",
                      "uid": "16c15b19-8518-4e51-aabf-feffe030e90f",
                      "depth": 5,
                      "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY Groups "
                    }
                  ],
                  "parentUid": "2525a527-7bd5-4823-8b8b-6f8b541ac5d6",
                  "uid": "d75a6a1c-fed5-4876-96fd-d463d2d3636d",
                  "depth": 4,
                  "lineage": "root com MYCOMPANY MYTOWN MYCOMPANY "
                }
              ],
              "parentUid": "ee20c8bc-8b93-4993-8e50-22ace77b4fcb",
              "uid": "2525a527-7bd5-4823-8b8b-6f8b541ac5d6",
              "depth": 3,
              "lineage": "root com MYCOMPANY MYTOWN "
            }
          ],
          "parentUid": "a8342a50-ce67-4753-9921-517d7124eac3",
          "uid": "ee20c8bc-8b93-4993-8e50-22ace77b4fcb",
          "depth": 2,
          "lineage": "root com MYCOMPANY "
        }
      ],
      "parentUid": "1b67f514-fe87-483e-94e0-bc778d004310",
      "uid": "a8342a50-ce67-4753-9921-517d7124eac3",
      "depth": 1,
      "lineage": "root com "
    }
  ],
  "parentUid": "",
  "uid": "1b67f514-fe87-483e-94e0-bc778d004310",
  "depth": 0,
  "lineage": "root "
}
``` 

# Docs
```
TYPES

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
    Node contains properties and methods to represent an object in the LDAP tree
    and to alter properties such as children, parent, value, etc.

func BuildTree(input []string) (root *Node, vis string, err error)
    BuildTree takes a slice of LDAP Distinguished Name strings and attempts to
    build a node tree that represents all of their relationships (if any) under
    a generic parent "root" node.

    It will return the root node object and the visualization string and any
    errors.

func (n *Node) AddChild(child *Node)
    AddChild adds a Node to this node's children and updates the provided Node's
    parent property as well.

func (n *Node) AddParent(parent *Node)
    AddParent updates the parent of the current node to the node provided in the
    argument. Somewhat safe.

func (n *Node) FmtTree(start string) string
    FmtTree returns a string formatted as a multiline tree representing thise
    Node and it's children.

func (n *Node) HasChild(value string) (found bool)
    HasChild searches the Node's children for a child with the requested value.
```
