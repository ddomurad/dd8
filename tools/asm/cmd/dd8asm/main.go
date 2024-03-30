package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	pkg "github.com/ddomurad/dd8/tools/asm/asm"
	"github.com/ddomurad/dd8/tools/asm/asm/assemblers"
	"github.com/ddomurad/dd8/tools/asm/asm/output"
)

func main() {
	outFile := flag.String("out", "", "output file")
	listDir := flag.String("list", "", "source listing dir")
	exportFile := flag.String("eout", "", "export labels output file")
	exportFilter := flag.String("eflt", "", "export labels name regex filter")

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

	reader := pkg.NewFileSourceReader("./" + worksparceDir)

	byteCode, sourceListing, exported, aerr, err := pkg.AssembleSrc(intpuFile, reader, assemblers.OpcodeAssemblerW65C02S, *listDir != "", *exportFile != "")

	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}

	if aerr != nil {
		os.Stderr.WriteString(aerr.Error())
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

	if *listDir != "" {
		writer := pkg.NewFileSourceWriter("./" + *listDir)
		err := sourceListing.Write(writer)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Stderr.WriteString("\n")
			os.Exit(1)
		}
	}

	if *exportFile != "" {
		writer := pkg.NewFileSourceWriter("./" + worksparceDir)
		err := exported.Write(*exportFile, writer, *exportFilter)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Stderr.WriteString("\n")
			os.Exit(1)
		}
	}

	os.Exit(0)
}
