package parser

import (
	"fmt"

	"github.com/ddomurad/dd8/tools/asm/internal/models"
)

type RawSourceFile struct {
	Source models.SourceContent
	Uri    models.SourceUri
}

func NewRawSourceFile(source models.SourceContent, uri models.SourceUri) RawSourceFile {
	return RawSourceFile{
		Source: source,
		Uri:    uri,
	}
}

type CodePointer struct {
	Uri    string
	Line   int
	Column int
}

func (p CodePointer) String() string {
	return fmt.Sprintf("%s:%d:%d", p.Uri, p.Line, p.Column)
}

type SourceEntryType string

const (
	SourceEntryTypeOpCode SourceEntryType = "op_code"
)

type SourceEntry struct {
	Type     SourceEntryType
	OpCode   string
	NumParam int64
}

type Source struct {
	Entries []SourceEntry
}

// func (s SourceTree) String() string {
// 	outStr := ""
// 	for _, oc := range s.OpCodes {
// 		outStr += oc + "\n"
// 	}
// 	return outStr
// }
