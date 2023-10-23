package main

import (
	"flag"
	"fmt"
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
	if len(tail) == 0 {
		os.Stderr.WriteString("command line error: missing source file")
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	if len(tail) > 1 {
		os.Stderr.WriteString(fmt.Sprintf("command line error: unexpected argument '%s'\n", tail[len(tail)-1]))
		os.Exit(1)
	}

	if *outFile == "" {
		os.Stderr.WriteString("command line error: missing output file")
		os.Stderr.WriteString("\n")
		flag.PrintDefaults()
		os.Exit(1)
		return
	}

	worksparceDir := path.Dir(tail[0])
	intpuFile := path.Base(tail[0])

	reader := internal.NewFileSourceReader("./" + worksparceDir)
	byteCode, err := internal.AssembleSrc(intpuFile, reader, assemblers.OpcodeAssemblerW65C02S)

	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}

	hexOutput := output.GetIntelHEX(byteCode)
	err = os.WriteFile(*outFile, hexOutput, 0644)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}

	os.Exit(0)
}
