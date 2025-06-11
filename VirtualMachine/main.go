package main

import (
	"fmt"
	"os"
	"strings"
	"vm/parser"
	"vm/vm_translator"
)

func main() {
	lines, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	linesStr := strings.Split(string(lines), "\n")

	parser := parser.NewParser(linesStr)
	result := vm_translator.Translate(parser)

	os.WriteFile(os.Args[2], []byte(result), 0644)

	fmt.Println("Assembler finished.")
}
