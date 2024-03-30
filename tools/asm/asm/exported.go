package asm

import (
	"fmt"
	"regexp"
)

// todo: rename to something better
type Exported struct {
	Labels map[string]int
}

func (s *Exported) Write(fname string, writer SourceWriter, filter string) error {
	return writer.WriteSourceFile(fname, s.String(filter))
}

func (s *Exported) String(filter string) string {
	if filter == "" {
		filter = ".*"
	}
	re := regexp.MustCompile(filter)

	out := "; exported labels\n"
	out += ".def (\n"
	for lname, laddr := range s.Labels {
		if re.MatchString(lname) {
			out += fmt.Sprintf("\t%s := 0x%04X\n", lname, laddr)
		}
	}
	out += ")"
	return out
}
