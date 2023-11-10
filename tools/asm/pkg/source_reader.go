package internal

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
