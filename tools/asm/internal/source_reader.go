package internal

import (
	"os"

	"github.com/ddomurad/dd8/tools/asm/internal/models"
)

type SourceReader interface {
	GetSourceUri(srcName models.SourceName) (models.SourceUri, error)
	ReadSourceFile(srcUri models.SourceUri) (models.SourceContent, error)
}

type FileSourceReader struct {
	rootDir models.DirPath
}

func NewFileSourceReader(rootDir models.DirPath) *FileSourceReader {
	return &FileSourceReader{
		rootDir: rootDir,
	}
}

func (r *FileSourceReader) GetSourceUri(srcName models.SourceName) (models.SourceUri, error) {
	return models.SourceUri(string(r.rootDir) + string(srcName)), nil
}

func (r *FileSourceReader) ReadSourceFile(srcUri models.SourceUri) (models.SourceContent, error) {
	data, err := os.ReadFile(string(srcUri))
	return models.SourceContent(data), err
}
