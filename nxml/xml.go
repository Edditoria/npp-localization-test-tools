package nxml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Cfg struct {
	KeysHaveEqualVal []string
}

type Doc struct {
	Root *Node
}

func ReadFromFile(filepath string) (Doc, error) {
	var doc Doc
	file, err := os.Open(filepath)
	if err != nil {
		return doc, err
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

	return doc, nil
}

type Node struct {
	Name     xml.Name
	Attrs    []xml.Attr
	Children []*Node
	Parent   *Node // For reference only. May need to update if the etree changes.
	// Xpath    string
	// Value    xml.CharData
}

type DiffType int

const (
	DiffPassed DiffType = iota
	DiffAdded
	DiffMissing
	DiffAttr
)

type DiffRecord struct {
	Type      DiffType
	Dirs      Dirs
	NodeLeft  *Node
	NodeRight *Node
}

type Dir struct {
	Node     *Node
	Position int
}

func (d Dir) AsString() string {
	return fmt.Sprintf("%s[%v]", d.Node.Name.Local, d.Position)
}

type Dirs []Dir

func (d Dirs) Xpath() string {
	o := "//"
	for _, dir := range d {
		o = o + dir.Node.Name.Local + "[" + strconv.Itoa(dir.Position) + "]"
	}
	return o
}

func Compare(baseDoc, userDoc *Doc, cfg Cfg) []DiffRecord {
	var diffs []DiffRecord
	// todo...
	return diffs
}
