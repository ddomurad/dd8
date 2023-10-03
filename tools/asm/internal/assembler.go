package internal

import (
	"github.com/ddomurad/dd8/tools/asm/internal/models"
	"github.com/ddomurad/dd8/tools/asm/internal/parser"
)

func Assemble(srcName models.SourceName, reader SourceReader) error {
	srcUri, err := reader.GetSourceUri(srcName)
	if err != nil {
		return err
	}

	src, err := reader.ReadSourceFile(srcUri)
	if err != nil {
		return err
	}

	tree := parser.ParseSource(parser.NewRawSourceFile(src, srcUri))
	_ = tree
	return nil
}
