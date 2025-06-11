package code

import (
	"fmt"
	"strings"
)

type Code struct {
}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) Dest(dest string) string {
	const adm = "ADM"
	result := []byte("000")

	for _, v := range dest {
		if strings.Contains(adm, string(v)) {
			result[strings.Index(adm, string(v))] = '1'
		}
	}

	return string(result)
}

func (c *Code) Comp(comp string) string {
	compMap := map[string]string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",
		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}

	return compMap[comp]
}

func (c *Code) Jump(jump string) string {
	jmp := []string{"", "JGT", "JEQ", "JGE", "JLT", "JNE", "JLE", "JMP"}

	for i, v := range jmp {
		if v == jump {
			// Преобразует число i в трехзначное двоичное число с ведущими нулями
			// Например: 5 -> "101", 2 -> "010"
			return fmt.Sprintf("%03b", i)
		}
	}
	return ""
}
