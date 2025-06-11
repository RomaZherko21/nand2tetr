package code_writer

var add = `
@SP
AM=M-1
D=M
M=0
A=A-1
M=D+M
`

var sub = `
@SP
AM=M-1
D=M
M=0
A=A-1
M=M-D
`

var neg = `
@SP
A=M-1
M=-M
`

var eq = `
@SP
AM=M-1
D=M
M=0
A=A-1
M=M-D
D=M
@TRUE_{i}
D;JEQ
@FALSE_{i}
0;JMP

(TRUE_{i})
@SP
A=M-1
M=-1
@END_{i}
0;JMP

(FALSE_{i})
@SP
A=M-1
M=0

(END_{i})
`

var gt = `
@SP
AM=M-1
D=M
M=0
A=A-1
M=M-D
D=M
@TRUE_{i}
D;JGT
@FALSE_{i}
0;JMP

(TRUE_{i})
@SP
A=M-1
M=-1
@END_{i}
0;JMP

(FALSE_{i})
@SP
A=M-1
M=0

(END_{i})
`

var lt = `
@SP
AM=M-1
D=M
M=0
A=A-1
M=M-D
D=M
@TRUE_{i}
D;JLT
@FALSE_{i}
0;JMP

(TRUE_{i})
@SP
A=M-1
M=-1
@END_{i}
0;JMP

(FALSE_{i})
@SP
A=M-1
M=0

(END_{i})
`

var and = `
@SP
AM=M-1
D=M
M=0
A=A-1
M=D&M
`

var or = `
@SP
AM=M-1
D=M
M=0
A=A-1
M=D|M
`

var not = `
@SP
A=M-1
M=!M
`
