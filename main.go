package main

import (
	"fmt"
	"npptt/nxml"
)

func printAll_test(node *nxml.Node) {
	fmt.Print(node.Name.Space, node.Name.Local, node.Attrs, node.Parent.ToString())
	if len(node.Children) > 0 {
		fmt.Printf(", %d children\n", len(node.Children))
		for _, n := range node.Children {
			printAll_test(n)
		}
	} else {
		fmt.Print(", no child\n")
	}
}

func main() {
	userFile := "./testdata/HongKongCantonese.xml"
	baseFile := "./testdata/English.xml"
	cfg := nxml.Cfg{KeysHaveEqualVal: []string{"menuId"}}

	userDoc, err := nxml.NewFromFile(userFile)
	if err != nil {
		panic(err)
	}
	printAll_test(userDoc.Root)

	baseDoc, err := nxml.NewFromFile(baseFile)
	if err != nil {
		panic(err)
	}
	// printAll_test(baseDoc.Root)

	nxml.Compare(&baseDoc, &userDoc, cfg)
}
