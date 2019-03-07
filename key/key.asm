;------------------------------------------------------------------------------
; KEY is a program that reads in input from the standard input device
; (keyboard). It accepts four kinds of input from the user: a space character,
; an upper-case alphabetic character, a lower-case alphabetic character,
; and a termination character key (the period key in this case). The following
; occurs during execution of the program:
; * If the character is a space or upper-case alphabetic character, it is
;	printed directly to the screen.
; * If it is a lower-case alphabetic character, the program translates it into
;	an upper-case alphabetic character representing the same letter in the
;	English language. It then prints to the screen. Some examples are:
;
;	 Input				Internal 					Output As
;	 Character			Representation				Character
;	-------------------------------------------------------------------------
;	 a					61h							A
;	 c					63h							C
;    z					7Ah						    Z
;
; * Finally, if the character is a period, the character is printed to the
;	screen and the program terminates execution.
;------------------------------------------------------------------------------
	.model small							; This program uses up to
											; 64k of data and 64k of code.
	.8086									; This program runs 8086
											; instructions.
	.stack 256								; The stack size is 256 bytes.
;------------------------------------------------------------------------------
;  This is the start of the code segment
;------------------------------------------------------------------------------
	.code
;------------------------------------------------------------------------------
; This is the start of the main program loop. It will execute instructions in
; the following pattern:
;	- Read in the character into the al register
; 	- translate the character in the al register into a character to be printed
;	- check to see if the character should be printed, if not loop back to the
;	  start of this loop.
;	- if the character should be printed, then print it.
;	- check to see if the character is a termination character, if not loop
;	  back to the start of the program
;	- If the character is the termination character, then execute the
;	  termination instructions and exit without error.
;------------------------------------------------------------------------------
getloop:									;
		mov 	BX, AX						; Move transition code into bx
		add 	BX, BX						; double for word data
		mov 	AX, CS:[trans + BX]			; find new read/write/exit code
		mov		DL, AL						; prepare to write a char (DH=char)
		int 	21h							; DOS interrupt issued.
		jmp 	getloop						; jump to the top of the loop.
;------------------------------------------------------------------------------
; Translation Table
;------------------------------------------------------------------------------
trans	dw 	022Eh dup (0800h)				; Read a new character, no print required (022Eh)
		dw 	4C00h							; Period Printed (022Eh), now exit
		dw 	05D1h dup (0800h)				; Read a new character, no print required (0800h)
		dw	20h dup (0800h)					; Read a new character, no print required (space)
		dw	0220h							; print 'space'
		dw 	0Dh dup (0800h)					; Read a new character, no print required (period)
		dw	022Eh							; print period
		dw	18 dup (0800h) 					; Read a new character, no print required ('A')
											; Print:
		dw 	0241h,0242h,0243h,0244h,0245h	; A-E
		dw	0246h,0247h,0248h,0249h,024Ah	; F-J
		dw 	024Bh,024Ch,024Dh,024Eh,024Fh	; K-O
		dw	0250h,0251h,0252h,0253h,0254h	; P-T
		dw	0255h,0256h,0257h,0258h,0259h	; U-Y
		dw	025Ah							; Z
		dw	6 dup (0800h) 				    ; Read a new character, no print required ('a')
											; Print and Change:
		dw	0241h,0242h,0243h,0244h,0245h	; a-e to A-E
		dw	0246h,0247h,0248h,0249h,024Ah	; f-j to F-J
		dw	024Bh,024Ch,024Dh,024Eh,024Fh	; k-o to K-O
		dw	0250h,0251h,0252h,0253h,0254h	; p-t to P-T
		dw	0255h,0256h,0257h,0258h,0259h	; u-y to U-Y
		dw	025Ah							; z to Z
		dw	133 dup (0800h) 				; Read a new character, no print required (remaining)
end 	 									; end of program marker
;------------------------------------------------------------------------------
; END OF PROGRAM KEY
;------------------------------------------------------------------------------
