
@256
D=A
@SP
M=D
	
@17
D=A
@SP
A=M
M=D
@SP
M=M+1

@18
D=A
@SP
A=M
M=D
@SP
M=M+1

@LCL
D=M         
@0
D=D+A       
@R13       
M=D        

@SP
AM=M-1     
D=M        

@R13         
A=M        
M=D

@LCL
D=M         
@1
D=D+A       
@R13       
M=D        

@SP
AM=M-1     
D=M        

@R13         
A=M        
M=D

@LCL        
D=M           
@0          
A=D+A        

D=M          

@SP         
A=M           
M=D 
@SP
M=M+1

@LCL        
D=M           
@0          
A=D+A        
D=M          

@SP         
A=M           
M=D 
@SP
M=M+1

@LCL        
D=M           
@0          
A=D+A        
D=M          

@SP         
A=M           
M=D 
@SP
M=M+1

(END)
0;JMP
