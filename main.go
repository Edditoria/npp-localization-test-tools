package main

import (
	"fmt"
	"npptt/cmd"
)

func printAll_test(node *cmd.Node) {
	fmt.Print(node.Name.Space, node.Name.Local, node.Attrs)
	if len(node.Children) > 0 {
		fmt.Printf(" : %v\n", len(node.Children))
		for _, n := range node.Children {
			printAll_test(n)
		}
	} else {
		fmt.Print(" : 0\n")
	}
}

func main() {
	userFile := "./testdata/HongKongCantonese.xml"
	baseFile := "./testdata/English.xml"
	cfg := cmd.Cfg{KeysHaveEqualVal: []string{"menuId"}}

	userDoc, err := cmd.ReadFromFile(userFile)
	if err != nil {
		panic(err)
	}
	printAll_test(userDoc.Root)

	baseDoc, err := cmd.ReadFromFile(baseFile)
	if err != nil {
		panic(err)
	}
	printAll_test(baseDoc.Root)
	cmd.Compare(&baseDoc, &userDoc, cfg)
}
