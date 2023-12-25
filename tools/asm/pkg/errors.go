package internal

import "fmt"

type SourceErrorType string

const (
	SourceErrorTypeSyntaxError              SourceErrorType = "syntax_err"
	SourceErrorTypeUnexpectedStructureError SourceErrorType = "struct_err"
	SourceErrorTypeEvalError                SourceErrorType = "eval_err"
	SourceErrorTypeAssemblyError            SourceErrorType = "assembly_err"
)

type SourceError struct {
	Type    SourceErrorType
	SrcName string
	Line    int
	Col     int
	Msg     string
}

func (r SourceError) Error() string {
	if r.Type == SourceErrorTypeEvalError {
		return fmt.Sprintf("[%s] %s:%d %s", r.Type, r.SrcName, r.Line, r.Msg)
	}

	if r.Type == SourceErrorTypeUnexpectedStructureError {
		return fmt.Sprintf("[%s] %s:%d %s", r.Type, r.SrcName, r.Line, r.Msg)
	}

	return fmt.Sprintf("[%s] %s:%d:%d %s", r.Type, r.SrcName, r.Line, r.Col, r.Msg)
}

type AssemblerError struct {
	Errors []SourceError
}

func NewAssemblerError() AssemblerError {
	return AssemblerError{
		Errors: []SourceError{},
	}
}

func (r AssemblerError) Error() string {
	outStr := ""
	for _, es := range r.Errors {
		outStr += es.Error() + "\n"
	}
	return outStr
}

func (r *AssemblerError) Append(e SourceError) {
	r.Errors = append(r.Errors, e)
}

func (r *AssemblerError) HasErrors() bool {
	return len(r.Errors) > 0
}
