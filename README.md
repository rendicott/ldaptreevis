# ldaptreevis
Structures LDAP Distinguished Name strings into a visualized tree. 

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

See example usage in `cmd/main.go`

You can also use the returned root node to do whatever you want, e.g., export to JSON

```
{
  "label": "root",
  "children": [
    {
      "label": "com",
      "children": [
        {
          "label": "MYCOMPANY",
          "children": [
            {
              "label": "MYTOWN",
              "children": [
                {
                  "label": "MYCOMPANY",
                  "children": [
                    {
                      "label": "Germany",
                      "children": [
                        {
                          "label": "MYTOWN",
                          "children": [
                            {
                              "label": "Groups",
                              "children": [
                                {
                                  "label": "SPSAdmins",
                                  "depth": 8
                                },
                                {
                                  "label": "FooAdmin",
                                  "depth": 8
                                },
                                {
                                  "label": "BarAdmin",
                                  "depth": 8
                                }
                              ],
                              "depth": 7
                            }
                          ],
                          "depth": 6
                        }
                      ],
                      "depth": 5
                    },
                    {
                      "label": "America",
                      "children": [
                        {
                          "label": "MYTOWN",
                          "children": [
                            {
                              "label": "Groups",
                              "children": [
                                {
                                  "label": "SPSAdmins",
                                  "depth": 8
                                },
                                {
                                  "label": "FooAdmin",
                                  "depth": 8
                                },
                                {
                                  "label": "BarAdmin",
                                  "depth": 8
                                }
                              ],
                              "depth": 7
                            }
                          ],
                          "depth": 6
                        },
                        {
                          "label": "YourTown",
                          "children": [
                            {
                              "label": "Groups",
                              "children": [
                                {
                                  "label": "SPSAdmins",
                                  "depth": 8
                                },
                                {
                                  "label": "FooAdmin",
                                  "depth": 8
                                },
                                {
                                  "label": "BarAdmin",
                                  "depth": 8
                                },
                                {
                                  "label": "gone",
                                  "children": [
                                    {
                                      "label": "Doug",
                                      "children": [
                                        {
                                          "label": "Bob",
                                          "children": [
                                            {
                                              "label": "LongAdmin",
                                              "depth": 11
                                            }
                                          ],
                                          "depth": 10
                                        }
                                      ],
                                      "depth": 9
                                    }
                                  ],
                                  "depth": 8
                                }
                              ],
                              "depth": 7
                            }
                          ],
                          "depth": 6
                        }
                      ],
                      "depth": 5
                    },
                    {
                      "label": "Groups",
                      "children": [
                        {
                          "label": "ShortAdmin",
                          "depth": 6
                        }
                      ],
                      "depth": 5
                    }
                  ],
                  "depth": 4
                }
              ],
              "depth": 3
            }
          ],
          "depth": 2
        }
      ],
      "depth": 1
    }
  ],
  "depth": 0
}
``` 
