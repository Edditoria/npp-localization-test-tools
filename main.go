package main

import (
	"fmt"
	"npptt/cmd"
)

func main() {
	baseFile := "./testdata/English.xml"
	var baseNodes []cmd.Node
	if err := cmd.ReadXml(baseFile, &baseNodes); err != nil {
		panic(err)
	}
	fmt.Println(baseNodes)
}
