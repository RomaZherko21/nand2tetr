package code_writer

// constant
var pushConstant = `
@%s
D=A
@SP
A=M
M=D
@SP
M=M+1
`

// static
var pushStatic = `
@static_%s
D=M
@SP
A=M
M=D
@SP
M=M+1
`

var popStatic = `
@SP
AM=M-1
D=M
@static_%s
M=D
`

// LCL (local)
var pushLocal = `
@LCL        
D=M           
@%s          
A=D+A        
D=M          

@SP         
A=M           
M=D 
@SP
M=M+1
`

var popLocal = `
@LCL
D=M         
@%s
D=D+A       
@R13       
M=D        

@SP
AM=M-1     
D=M        

@R13         
A=M        
M=D
`

// ARG (argument)
var pushArgument = `
@ARG
D=M
@%s
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`

var popArgument = `
@ARG
D=M         
@%s
D=D+A       
@R13       
M=D        

@SP
AM=M-1     
D=M        

@R13         
A=M        
M=D
`

// THIS
var pushThis = `
@THIS
D=M
@%s
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`

var popThis = `
@THIS
D=M         
@%s
D=D+A       
@R13       
M=D        

@SP
AM=M-1     
D=M        

@R13         
A=M        
M=D
`

// THAT
var pushThat = `
@THAT
D=M
@%s
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`

var popThat = `
@THAT
D=M         
@%s
D=D+A       
@R13       
M=D        

@SP
AM=M-1     
D=M        

@R13         
A=M        
M=D
`

// TEMP
var pushTemp = `
@R5
D=M
@%s
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
`

var popTemp = `
@R5
D=M         
@%s
D=D+A       
@R13       
M=D        

@SP
AM=M-1     
D=M        

@R13         
A=M        
M=D
`

// POINTER 0
var pushPointer0 = `
@THIS
D=A
@SP
A=M
M=D
@SP
M=M+1
`

var popPointer0 = `
@SP
AM=M-1
D=M
@THIS
M=D
`

// POINTER 1
var pushPointer1 = `
@THAT
D=A
@SP
A=M
M=D
@SP
M=M+1
`

var popPointer1 = `
@SP
AM=M-1
D=M
@THAT
M=D
`
