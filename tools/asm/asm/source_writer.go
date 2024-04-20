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

type MemorySourceWriter struct {
	buffers map[string][]byte
}

func NewMemorySourceWriter(buffers map[string][]byte) *MemorySourceWriter {
	return &MemorySourceWriter{
		buffers: buffers,
	}
}

func (w *MemorySourceWriter) WriteSourceFile(srcName string, src string) error {
	w.buffers[srcName] = []byte(src)
	return nil
}
