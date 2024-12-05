package nxml

type Cfg struct {
	KeysHaveEqualVal []string
	KeysHaveDiffVal  []string
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
	Dirs      []Dir
	NodeLeft  *Node
	NodeRight *Node
}

func Compare(baseDoc, userDoc *Doc, cfg Cfg) []DiffRecord {
	var diffs []DiffRecord
	// todo...
	return diffs
}
