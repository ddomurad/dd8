package asm

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gotidy/ptr"
)

type CodeListing struct {
	Address        uint16
	Bytes          []byte
	ProxyLine      *int
	ProxyFile      *string
	VirtualSrcLine string
}

func (c CodeListing) IsProxied() bool {
	return c.ProxyLine != nil
}

func (c CodeListing) IsVirtualLine() bool {
	return c.VirtualSrcLine != ""
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

func (l *SourceListing) InsertVirtualLine(srcName string, line int, src string) {
	if l.overide {
		srcName = l.overideSource
		line = l.overideLine
	}

	if _, ok := l.FileListings[srcName].CodeLines[line]; !ok {
		l.FileListings[srcName].CodeLines[line] = []CodeListing{}
	}

	l.FileListings[srcName].CodeLines[line] = append(
		l.FileListings[srcName].CodeLines[line],
		CodeListing{
			VirtualSrcLine: src,
		})
}

func (l *SourceListing) AddCode(srcName string, line int, location int, codes ...byte) {
	var proxyFile *string = nil
	var proxyLine *int = nil
	if l.overide {
		proxyFile = ptr.String(srcName)
		proxyLine = ptr.Int(line)
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
				Address:   uint16(location + i),
				Bytes:     codes[i:min(i+4, len(codes))],
				ProxyFile: proxyFile,
				ProxyLine: proxyLine,
			})
	}
}

func (s *SourceListing) Write(writer SourceWriter) error {
	for name, listing := range s.FileListings {
		err := writer.WriteSourceFile(
			strings.ReplaceAll(name, ".asm", ".list"),
			listing.String(s.FileListings),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (l FileListing) String(listings map[string]FileListing) string {
	listing := "Line  Loc   Code         Source\n"

	for i, srcLine := range l.SourceLines {
		prtLine := srcLine
		if codeLines, ok := l.CodeLines[i+1]; ok {
			prevProxyLine := -1
			for j, code := range codeLines {
				if code.IsVirtualLine() {
					listing += fmt.Sprintf("%5d                   %s\n", i+1, code.VirtualSrcLine)
					continue
				}

				if code.IsProxied() {
					proxyListing, ok := listings[*code.ProxyFile]
					if !ok {
						prtLine = "ERROR: Proxy file not found"
					} else {
						prtLine = proxyListing.SourceLines[*code.ProxyLine-1]
					}
				}
				skipPrtLine := j != 0
				if code.IsProxied() {
					if prevProxyLine != *code.ProxyLine {
						skipPrtLine = false
					}
					prevProxyLine = *code.ProxyLine
				}

				if !skipPrtLine {
					listing += fmt.Sprintf("%5d  %04x  %-8s  %s\n", i+1, code.Address, hex.EncodeToString(code.Bytes), prtLine)
				} else {
					listing += fmt.Sprintf("%5d  %04x  %-8s\n", i+1, code.Address, hex.EncodeToString(code.Bytes))
				}
			}
		} else {
			listing += fmt.Sprintf("%5d                   %s\n", i+1, prtLine)
		}
	}

	return listing
}
