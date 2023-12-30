package asm

import (
	"fmt"
	"strings"
)

type CodeListing struct {
	Address int16
	Bytes   []int8
}

type FileListing struct {
	FileNumber  int
	SourceLines []string
	CodeLines   map[int]CodeListing
}

type SourceListing struct {
	FileListings map[string]FileListing
}

func NewSourceListing() *SourceListing {
	return &SourceListing{
		FileListings: map[string]FileListing{},
	}
}

func (l *SourceListing) AddSource(srcName string, src string) {
	fl := FileListing{
		FileNumber:  len(l.FileListings),
		SourceLines: strings.Split(src, "\n"),
		CodeLines:   map[int]CodeListing{},
	}
	l.FileListings[srcName] = fl
}

func (l SourceListing) PrintFileIndex() string {
	var findex string
	findex += "Fl# Name\n"
	for name, listing := range l.FileListings {
		findex += fmt.Sprintf("%3d %s\n", listing.FileNumber, name)
	}
	return findex
}

func (l SourceListing) ListFile() string {

	return ""
}

func (l SourceListing) String() string {
	var listing string

	listing += l.PrintFileIndex()
	listing += "\n"
	listing += l.ListFile() //todo: rename

	return listing
}
