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
	baseFile := "./testdata/English.xml"
	var baseDoc cmd.Doc
	if err := baseDoc.ReadXml(baseFile); err != nil {
		panic(err)
	}
	printAll_test(baseDoc.Root)

}
