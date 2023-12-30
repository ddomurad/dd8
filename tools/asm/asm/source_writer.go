package asm

import (
	"os"
	"path"
)

type SourceWriter interface {
	WriteSourceFile(srcUri string, src string) error
}

type FileSourceWriter struct {
	rootDir string
}

func NewFileSourceWriter(rootDir string) *FileSourceWriter {
	return &FileSourceWriter{
		rootDir: rootDir,
	}
}

func (w *FileSourceWriter) WriteSourceFile(srcName string, src string) error {
	srcPath := path.Join(w.rootDir, srcName)
	return os.WriteFile(srcPath, []byte(src), 0644)
}
