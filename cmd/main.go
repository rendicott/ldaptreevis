package main

import (
	"fmt"
	"encoding/json"
	"github.com/rendicott/ldaptreevis"
)

func main() {
	s := []string{
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
	}
	root, _, err := ldaptreevis.BuildTree(s)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	//fmt.Println(vis)

	b, err := json.Marshal(root)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(string(b))
}
