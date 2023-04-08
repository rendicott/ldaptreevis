package ldaptreevis

import (
    "testing"
)

//TODO: Add tests for safety of modifying the node tree using the exported methods

// TestBuildTreeStandard passes a large LDAP slice structure and checks
// for the proper formatting of the output
func TestBuildTreeStandard(t *testing.T) {
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
	want := testBuildTreeStandardWant
	_, result, err := BuildTree(s)
	got := result
	if want != got || err != nil {
		t.Fatalf(`BuildTree(%s) = %s, %v, want match for %#q, nil`,
			s, got, err, want)
	}
}

// TestBuildTreeSDupe passes a simple LDAP slice structure with 
// known duplicate and checks for the proper formatting of the output
func TestBuildTreeDupe(t *testing.T) {
	s := []string{
		"CN=SPSAdmins,OU=Groups,OU=MYTOWN,OU=Germany,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
		"CN=SPSAdmins,OU=Groups,OU=MYTOWN,OU=Germany,OU=MYCOMPANY,DC=MYTOWN,DC=MYCOMPANY,DC=com",
	}
	want := testBuildTreeDupeWant
	_, result, err := BuildTree(s)
	got := result
	if want != got || err != nil {
		t.Fatalf(`BuildTree(%s) = %s, %v, want match for %#q, nil`,
			s, got, err, want)
	}
}

var testBuildTreeDupeWant = `root
  com
    MYCOMPANY
      MYTOWN
        MYCOMPANY
          Germany
            MYTOWN
              Groups
                SPSAdmins
`

var testBuildTreeStandardWant = `root
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
`
