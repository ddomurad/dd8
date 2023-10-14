package main

import (
	"flag"
	"os"
	"path"

	"github.com/ddomurad/dd8/tools/asm/internal"
	"github.com/ddomurad/dd8/tools/asm/internal/assemblers"
	"github.com/ddomurad/dd8/tools/asm/internal/output"
)

func main() {
	outFile := flag.String("out", "", "output file")

	flag.Parse()
	tail := flag.Args()
	if len(tail) != 1 {
		panic("lol") //todo: ??
	}

	if *outFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	worksparceDir := path.Dir(tail[0])
	intpuFile := path.Base(tail[0])

	reader := internal.NewFileSourceReader("./" + worksparceDir)
	ast, err := internal.CompileAST(intpuFile, reader)
	if err != nil {
		panic(err)
	}

	byteCode, err := internal.Assemble(ast, assemblers.OpcodeAssemblerW65C02S)
	if err != nil {
		panic(err)
	}

	hexOutput := output.GetIntelHEX(byteCode)
	err = os.WriteFile(*outFile, hexOutput, 0644)
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
