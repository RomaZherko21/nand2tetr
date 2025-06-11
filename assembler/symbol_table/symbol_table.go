package symbol_table

type SymbolTable struct {
	table   map[string]int
	address int
}

func NewSymbolTable() *SymbolTable {
	table := make(map[string]int)

	table["SP"] = 0
	table["LCL"] = 1
	table["ARG"] = 2
	table["THIS"] = 3
	table["THAT"] = 4
	table["R0"] = 0
	table["R1"] = 1
	table["R2"] = 2
	table["R3"] = 3
	table["R4"] = 4
	table["R5"] = 5
	table["R6"] = 6
	table["R7"] = 7
	table["R8"] = 8
	table["R9"] = 9
	table["R10"] = 10
	table["R11"] = 11
	table["R12"] = 12
	table["R13"] = 13
	table["R14"] = 14
	table["R15"] = 15

	return &SymbolTable{
		table:   table,
		address: 16,
	}
}

func (s *SymbolTable) AddEntry(symbol string, address int) {
	if address == -1 {
		s.table[symbol] = s.address
		s.address++
	} else {
		s.table[symbol] = address
	}
}

func (s *SymbolTable) Contains(symbol string) bool {
	_, ok := s.table[symbol]
	return ok
}

func (s *SymbolTable) GetAddress(symbol string) int {
	return s.table[symbol]
}
