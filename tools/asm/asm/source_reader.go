package asm

import (
	"os"
	"path"
)

type SourceReader interface {
	ReadSourceFile(srcUri string) (string, error)
}

type FileSourceReader struct {
	rootDir string
}

func NewFileSourceReader(rootDir string) *FileSourceReader {
	return &FileSourceReader{
		rootDir: rootDir,
	}
}

func (r *FileSourceReader) ReadSourceFile(srcName string) (string, error) {
	srcPath := path.Join(r.rootDir, srcName)
	data, err := os.ReadFile(string(srcPath))
	return string(data), err
}

type MemorySourceReader struct {
	buffers map[string][]byte
}

func NewMemorySourceReader(buffers map[string][]byte) *MemorySourceReader {
	return &MemorySourceReader{
		buffers: buffers,
	}
}

func (r *MemorySourceReader) ReadSourceFile(srcName string) (string, error) {
	data, ok := r.buffers[srcName]
	if !ok {
		return "", os.ErrNotExist
	}
	return string(data), nil
}
