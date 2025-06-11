package main

import (
	"assembler/assembler"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting assembler...")

	// read input file from terminal
	lines, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	linesStr := strings.Split(string(lines), "\n")

	assembler := assembler.NewAssembler(linesStr)
	result := assembler.Assemble()

	// create and write result to output file
	os.WriteFile(os.Args[2], []byte(result), 0644)

	fmt.Println("Assembler finished.")
}
