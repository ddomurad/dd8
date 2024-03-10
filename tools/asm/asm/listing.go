package asm

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type CodeListing struct {
	Address uint16
	Bytes   []byte
}

type FileListing struct {
	SourceLines []string
	CodeLines   map[int][]CodeListing
}

type SourceListing struct {
	FileListings  map[string]FileListing
	overide       bool
	overideLine   int
	overideSource string
}

func NewSourceListing() *SourceListing {
	return &SourceListing{
		FileListings: map[string]FileListing{},
	}
}

func (l *SourceListing) OverrideLine(source string, line int) bool {
	if l.overide {
		return false
	}

	l.overide = true
	l.overideSource = source
	l.overideLine = line
	return true
}

func (l *SourceListing) ClearOverride() {
	l.overide = false
}

func (l *SourceListing) AddSource(srcName string, src string) {
	fl := FileListing{
		SourceLines: strings.Split(src, "\n"),
		CodeLines:   map[int][]CodeListing{},
	}
	l.FileListings[srcName] = fl
}

func (l *SourceListing) ClearCodePass() {
	for k, v := range l.FileListings {
		v.CodeLines = map[int][]CodeListing{}
		l.FileListings[k] = v
	}
}

func (l *SourceListing) AddCode(srcName string, line int, location int, codes ...byte) {
	if l.overide {
		srcName = l.overideSource
		line = l.overideLine
	}

	if _, ok := l.FileListings[srcName].CodeLines[line]; !ok {
		l.FileListings[srcName].CodeLines[line] = []CodeListing{}
	}

	for i := 0; i < len(codes); i += 4 {
		l.FileListings[srcName].CodeLines[line] = append(
			l.FileListings[srcName].CodeLines[line],
			CodeListing{
				Address: uint16(location + i),
				Bytes:   codes[i:min(i+4, len(codes))],
			})
	}
}

func (s *SourceListing) Write(writer SourceWriter) error {
	for name, listing := range s.FileListings {
		err := writer.WriteSourceFile(strings.ReplaceAll(name, ".asm", ".list"), listing.String())
		if err != nil {
			return err
		}
	}

	return nil
}

func (l FileListing) String() string {
	listing := "Line  Loc   Code         Source\n"

	for i, line := range l.SourceLines {
		if codeLines, ok := l.CodeLines[i+1]; ok {
			for j, code := range codeLines {
				if j == 0 {
					listing += fmt.Sprintf("%5d  %04x  %-8s  %s\n", i+1, code.Address, hex.EncodeToString(code.Bytes), line)
				} else {
					listing += fmt.Sprintf("%5d  %04x  %-8s\n", i+1, code.Address, hex.EncodeToString(code.Bytes))
				}
			}
		} else {
			listing += fmt.Sprintf("%5d                   %s\n", i+1, line)
		}
	}

	return listing
}
