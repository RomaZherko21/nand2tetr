package vm_translator

import (
	"vm/code_writer"
	"vm/parser"
)

func Translate(parser *parser.Parser) string {
	codeWriter := code_writer.NewCodeWriter()
	result := codeWriter.WriteInit()

	for parser.HasMoreLines() {
		commandType := parser.CommandType()
		switch commandType {
		case "C_ARITHMETIC":
			result += codeWriter.WriteArithmetic(parser.Arg1())
		case "C_PUSH", "C_POP":
			result += codeWriter.WritePushPop(parser.Command(), parser.Arg1(), parser.Arg2())
		case "C_LABEL":
			result += codeWriter.WriteLabel(parser.Arg1())
		case "C_GOTO":
			result += codeWriter.WriteGoto(parser.Arg1())
		case "C_IF":
			result += codeWriter.WriteIf(parser.Arg1())
		}

		parser.Advance()
	}

	result += codeWriter.Close()
	return result
}
