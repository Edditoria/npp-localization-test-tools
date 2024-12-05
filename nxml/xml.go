package nxml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Doc struct {
	// Just normal node but parent is nil.
	Root *Node
}

type Node struct {
	Name     xml.Name
	Attrs    []xml.Attr
	Children []*Node
	Parent   *Node // Cache. May need to update when the etree changes.
	// Path     string
	// Value    xml.CharData
}

func (n *Node) Xdir() string {
	pos := n.Xposition()
	return n.Name.Local + "[" + strconv.FormatInt(pos, 10) + "]"
}

// Get position under the same element name. Compatible to XPath.
// According to W3C, position starts from 1.
func (n *Node) Xposition() int64 {
	var count int64 = 0
	if n.Parent == nil {
		return 1
	}
	for _, elm := range n.Children {
		if elm == n {
			count++
			break
		}
		if elm.Name.Local == n.Name.Local {
			count++
		}
	}
	return count
}

func NewFromFile(filepath string) (Doc, error) {
	var doc Doc
	file, err := os.Open(filepath)
	if err != nil {
		return doc, err
	}
	defer file.Close()

	// For each element, write to `nodes`.
	decoder := xml.NewDecoder(file)
	var parent *Node
	var position int
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
				Name:  t.Name,
				Attrs: t.Attr,
			}
			if parent == nil {
				// fmt.Println("I am root. I only occur once.")
				thisNode.Parent = Dir{Node: nil, Position: 0}
				doc.Root = &thisNode
				parent = &thisNode
				position = 0
			} else {
				thisNode.Parent = Dir{Node: parent, Position: position}
				parent.Children = append(parent.Children, &thisNode)
				parent = &thisNode
				position = 0
			}
		case xml.EndElement:
			parent = parent.Parent.Node
			if parent != nil {
				position = len(parent.Children)
			}
		}
	}

	return doc, nil
}
