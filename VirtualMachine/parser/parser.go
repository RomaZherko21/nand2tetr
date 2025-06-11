package parser

import (
	"strings"
)

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

type CommandType string

const (
	C_ARITHMETIC             = "C_ARITHMETIC"
	C_PUSH                   = "C_PUSH"
	C_POP                    = "C_POP"
	C_LABEL      CommandType = "C_LABEL"
	C_GOTO       CommandType = "C_GOTO"
	C_IF         CommandType = "C_IF"
	C_FUNCTION   CommandType = "C_FUNCTION"
	C_RETURN     CommandType = "C_RETURN"
	C_CALL       CommandType = "C_CALL"
)

// CurrentToken returns the current token
func (p *Parser) CommandType() CommandType {
	line := p.lines[p.idx]

	if strings.HasPrefix(line, "push") {
		return C_PUSH
	}

	if strings.HasPrefix(line, "pop") {
		return C_POP
	}

	if strings.HasPrefix(line, "label") {
		return C_LABEL
	}

	if strings.HasPrefix(line, "goto") {
		return C_GOTO
	}

	if strings.HasPrefix(line, "if-goto") {
		return C_IF
	}

	if strings.HasPrefix(line, "function") {
		return C_FUNCTION
	}

	if strings.HasPrefix(line, "return") {
		return C_RETURN
	}

	if strings.HasPrefix(line, "call") {
		return C_CALL
	}
	return C_ARITHMETIC
}

func (p *Parser) Arg1() string {
	if p.CommandType() == C_RETURN {
		return ""
	}

	line := p.lines[p.idx]

	if p.CommandType() == C_ARITHMETIC {
		return line
	}

	parts := strings.Split(line, " ")
	return parts[1]
}

func (p *Parser) Arg2() string {
	line := p.lines[p.idx]

	if p.CommandType() == C_PUSH || p.CommandType() == C_POP || p.CommandType() == C_FUNCTION || p.CommandType() == C_CALL {
		parts := strings.Split(line, " ")
		return parts[2]
	}

	return ""
}

func (p *Parser) Command() string {
	line := p.lines[p.idx]
	parts := strings.Split(line, " ")
	return parts[0]
}
