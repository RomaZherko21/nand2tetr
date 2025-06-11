package parser

import "strings"

type Parser struct {
	lines []string
	idx   int
}

func NewParser(lines []string) *Parser {
	return &Parser{
		lines: lines,
		idx:   0,
	}
}

func (p *Parser) HasMoreLines() bool {
	return p.idx < len(p.lines)
}

// Advance read the next line
func (p *Parser) Advance() {
	p.idx++

	if !p.HasMoreLines() {
		return
	}

	for p.CheckComments() || p.CheckEmptyLine() {
		p.idx++
	}
}

func (p *Parser) CheckComments() bool {
	line := p.lines[p.idx]
	return strings.HasPrefix(line, "//")
}

func (p *Parser) CheckEmptyLine() bool {
	line := p.lines[p.idx]
	return line == ""
}

type InstructionType string

const (
	// A instruction, @xxx where xxx is a decimal number or a symbol
	InstructionTypeA InstructionType = "A"
	// C instruction, dest=comp;jump
	InstructionTypeC InstructionType = "C"
	// L instruction, (xxx) where xxx is a symbol
	InstructionTypeL InstructionType = "L"
)

// CurrentToken returns the current token
func (p *Parser) InstructionType() InstructionType {
	line := p.lines[p.idx]
	if strings.HasPrefix(line, "@") {
		return InstructionTypeA
	}
	if strings.HasPrefix(line, "(") && strings.HasSuffix(line, ")") {
		return InstructionTypeL
	}
	return InstructionTypeC
}

// Symbol returns the symbol of the instruction if it is an A or L instruction
func (p *Parser) Symbol() string {
	line := p.lines[p.idx]
	if p.InstructionType() == InstructionTypeA {
		return line[1:]
	}
	if p.InstructionType() == InstructionTypeL {
		return line[1 : len(line)-1]
	}
	return ""
}

// Dest returns the dest of the instruction if it is a C instruction
func (p *Parser) Dest() string {
	line := p.lines[p.idx]
	if p.InstructionType() == InstructionTypeC && strings.Contains(line, "=") {
		return line[:strings.Index(line, "=")]
	}
	return ""
}

// Comp returns the comp of the instruction if it is a C instruction
func (p *Parser) Comp() string {
	line := p.lines[p.idx]
	if p.InstructionType() == InstructionTypeC && strings.Contains(line, "=") {
		return line[strings.Index(line, "=")+1:]
	}
	return ""
}

// Jump returns the jump of the instruction if it is a C instruction
func (p *Parser) Jump() string {
	line := p.lines[p.idx]
	if p.InstructionType() == InstructionTypeC && strings.Contains(line, ";") {
		return line[strings.Index(line, ";")+1:]
	}
	return ""
}
