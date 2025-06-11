package assembler

import (
	"assembler/code"
	. "assembler/parser"
	aParser "assembler/parser"
	"assembler/symbol_table"
	"fmt"
	"strconv"
	"strings"
)

type Assembler struct {
	sourceCode []string
	binaryCode []string
}

func NewAssembler(lines []string) *Assembler {
	return &Assembler{
		sourceCode: lines,
	}
}

func (a *Assembler) Assemble() string {
	parser := aParser.NewParser(a.sourceCode)
	symbolTable := symbol_table.NewSymbolTable()

	for parser.HasMoreLines() {
		codeLine := 0
		if parser.InstructionType() == InstructionTypeL {
			symbol := parser.Symbol()
			symbolTable.AddEntry(symbol, codeLine+1)
			fmt.Println("L", symbol)
		}
		codeLine++
		parser.Advance()
	}

	parser = aParser.NewParser(a.sourceCode)
	code := code.NewCode()

	for parser.HasMoreLines() {

		switch parser.InstructionType() {
		case InstructionTypeA:
			symbol := parser.Symbol()

			num, isNotNumber := strconv.Atoi(symbol)

			if isNotNumber == nil {
				binary := strconv.FormatInt(int64(num), 2)
				binary = strings.Repeat("0", 16-len(binary)) + binary
				a.binaryCode = append(a.binaryCode, binary)
				fmt.Println("A", binary)
			} else {
				if symbolTable.Contains(symbol) {
					address := symbolTable.GetAddress(symbol)
					binary := strconv.FormatInt(int64(address), 2)
					binary = strings.Repeat("0", 16-len(binary)) + binary
					a.binaryCode = append(a.binaryCode, binary)
					fmt.Println("A", binary)
				} else {
					symbolTable.AddEntry(symbol, -1)
					address := symbolTable.GetAddress(symbol)
					binary := strconv.FormatInt(int64(address), 2)
					binary = strings.Repeat("0", 16-len(binary)) + binary
					a.binaryCode = append(a.binaryCode, binary)
					fmt.Println("A", binary)
				}
			}
		case InstructionTypeC:
			comp := parser.Comp()
			dest := parser.Dest()
			jump := parser.Jump()

			binary := "111" + code.Comp(comp) + code.Dest(dest) + code.Jump(jump)
			a.binaryCode = append(a.binaryCode, binary)
			fmt.Println("C", binary)
		}

		parser.Advance()
	}

	return strings.Join(a.binaryCode, "\n")
}
