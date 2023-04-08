package main

import (
	"fmt"
	"github.com/rendicott/ldaptreevis"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s []string
	// build DN slice from STDIN
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}
	// pass into the tree builder
	_, vis, err := ldaptreevis.BuildTree(s)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// print the visualization
	fmt.Println(vis)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
