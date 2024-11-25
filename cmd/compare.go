package cmd

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type DiffType = int

const (
	DiffPassed DiffType = iota
	DiffAdded
	DiffMissing
	DiffAttr
)

type Node struct {
	Xpath string
	Name  xml.Name
	Attrs []xml.Attr
	// Value    xml.CharData
	// Children []*Node
}

type DiffRecord struct {
	Type      DiffType
	Xpath     string
	NodeLeft  *Node
	NodeRight *Node
}

func ReadXml(filepath string, nodes *[]Node) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// For each element, write to `nodes`.
	decoder := xml.NewDecoder(file)
	paths := []string{}
	for {
		token, err := decoder.Token()
		// Handling EOF before looping tokens:
		if err == io.EOF {
			// fmt.Println("(EOF)")
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		switch t := token.(type) {
		case xml.StartElement:
			paths = append(paths, t.Name.Local)
			thisXpath := "/" + strings.Join(paths, "/")
			thisNode := Node{
				Xpath: thisXpath,
				Name:  t.Name,
			}
			thisNode.Attrs = t.Attr
			*nodes = append(*nodes, thisNode)
		case xml.EndElement:
			paths = paths[:len(paths)-1]
		}
	}

	return nil
}
