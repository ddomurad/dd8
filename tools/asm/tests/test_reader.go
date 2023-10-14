package tests

import (
	"fmt"
)

type TestSourceReader struct {
	testSources map[string]string
}

func newTestSourceReader() *TestSourceReader {
	return &TestSourceReader{
		testSources: map[string]string{},
	}
}

func (r *TestSourceReader) ReadSourceFile(srcName string) (string, error) {
	src, ok := r.testSources[srcName]
	if !ok {
		panic(fmt.Sprint("missing test source: ", srcName))
	}

	return src, nil
}

func (r *TestSourceReader) setSourceFile(srcName string, src string) {
	r.testSources[srcName] = src
}
