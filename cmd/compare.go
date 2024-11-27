package cmd

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Doc struct {
	Root *Node
}

func (doc *Doc) ReadXml(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// For each element, write to `nodes`.
	decoder := xml.NewDecoder(file)
	parentNode := doc.Root
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
			thisNode := Node{
				Name:   t.Name,
				Attrs:  t.Attr,
				Parent: parentNode,
			}
			if parentNode == nil {
				// fmt.Println("I am root. I only occur once.")
				doc.Root = &thisNode
				parentNode = &thisNode
			} else {
				parentNode.Children = append(parentNode.Children, &thisNode)
				parentNode = &thisNode
			}
		case xml.EndElement:
			parentNode = parentNode.Parent
		}
	}

	return nil
}

type Node struct {
	Name     xml.Name
	Attrs    []xml.Attr
	Children []*Node
	Parent   *Node
	// Xpath    string
	// Value    xml.CharData
}

type DiffType = int

const (
	DiffPassed DiffType = iota
	DiffAdded
	DiffMissing
	DiffAttr
)

type DiffRecord struct {
	Type      DiffType
	Xpath     string
	NodeLeft  *Node
	NodeRight *Node
}
