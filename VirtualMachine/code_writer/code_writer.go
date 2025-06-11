package code_writer

import "fmt"

type CodeWriter struct {
	staticIndex int
	constIndex  int
}

func NewCodeWriter() *CodeWriter {
	return &CodeWriter{}
}

func (c *CodeWriter) WriteInit() string {
	return `
@256
D=A
@SP
M=D
	`
}

func (c *CodeWriter) WriteArithmetic(command string) string {
	switch command {
	case "add":
		return add
	case "sub":
		return sub
	case "neg":
		return neg
	case "eq":
		return eq
	case "gt":
		return gt
	case "lt":
		return lt
	case "and":
		return and
	case "or":
		return or
	case "not":
		return not
	}
	return ""
}

func (c *CodeWriter) WritePushPop(command string, segment string, index string) string {
	if command == "push" {
		if segment == "constant" {
			return fmt.Sprintf(pushConstant, index)
		}
		if segment == "static" {
			return fmt.Sprintf(pushStatic, index)
		}
		if segment == "local" {
			return fmt.Sprintf(pushLocal, index)
		}
		if segment == "argument" {
			return fmt.Sprintf(pushArgument, index)
		}
		if segment == "this" {
			return fmt.Sprintf(pushThis, index)
		}
		if segment == "that" {
			return fmt.Sprintf(pushThat, index)
		}
		if segment == "temp" {
			return fmt.Sprintf(pushTemp, index)
		}
		if segment == "pointer" {
			if index == "0" {
				return pushPointer0
			}
			if index == "1" {
				return pushPointer1
			}
		}

	}

	if command == "pop" {
		if segment == "static" {
			return fmt.Sprintf(popStatic, index)
		}
		if segment == "local" {
			return fmt.Sprintf(popLocal, index)
		}
		if segment == "argument" {
			return fmt.Sprintf(popArgument, index)
		}
		if segment == "this" {
			return fmt.Sprintf(popThis, index)
		}
		if segment == "that" {
			return fmt.Sprintf(popThat, index)
		}
		if segment == "temp" {
			return fmt.Sprintf(popTemp, index)
		}
		if segment == "pointer" {
			if index == "0" {
				return popPointer0
			}
			if index == "1" {
				return popPointer1
			}
		}
	}

	return ""
}

func (c *CodeWriter) WriteLabel(lbl string) string {
	return fmt.Sprintf(label, lbl)
}

func (c *CodeWriter) WriteGoto(label string) string {
	return fmt.Sprintf(gotoLabel, label)
}

func (c *CodeWriter) WriteIf(label string) string {
	return fmt.Sprintf(ifGotoLabel, label)
}

func (c *CodeWriter) Close() string {
	return `
(END)
0;JMP
`
}
