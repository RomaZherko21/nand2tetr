package code_writer

var label = `
(%s)
`

var gotoLabel = `
@%s
0;JMP
`

var ifGotoLabel = `
@SP
AM=M-1
D=M
@%s
D;JNE
`

var functionLabel = `
(%s)
`

var callLabel = `
// сохраняем адрес возврата
@RETURN_%s
D=A
@SP
A=M
M=D
@SP
M=M+1

// сохраняем LCL, ARG, THIS, THAT
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1

@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1

@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1

@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1

// устанавливаем ARG = SP-5-n
@SP
D=M
@5
D=D-A
@%d    // количество аргументов
D=D-A
@ARG
M=D

// устанавливаем LCL = SP
@SP
D=M
@LCL
M=D

// переходим к функции
@%s
0;JMP
(RETURN_%s)
`

// Код для возврата из функции
var returnLabel = `
// FRAME = LCL
@LCL
D=M
@R13        // R13 = FRAME
M=D

// RET = *(FRAME-5)
@5
A=D-A
D=M
@R14        // R14 = RET (адрес возврата)
M=D

// *ARG = pop() - помещаем возвращаемое значение
@SP
AM=M-1
D=M         // D = возвращаемое значение
@ARG
A=M
M=D         // *ARG = возвращаемое значение

// SP = ARG + 1
@ARG
D=M+1
@SP
M=D         // восстанавливаем SP вызывающей функции

// восстанавливаем сегменты вызывающей функции
@R13        // FRAME
AM=M-1
D=M
@THAT
M=D         // THAT = *(FRAME-1)

@R13
AM=M-1
D=M
@THIS
M=D         // THIS = *(FRAME-2)

@R13
AM=M-1
D=M
@ARG
M=D         // ARG = *(FRAME-3)

@R13
AM=M-1
D=M
@LCL
M=D         // LCL = *(FRAME-4)

// goto RET
@R14
A=M
0;JMP
`
