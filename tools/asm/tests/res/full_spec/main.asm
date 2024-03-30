; this is the only comment type 
; only EOL will terminate a comment 

; all preprocessor directives start with a '.' - except labels, labels always end with ':'
; ... and template executions, those start with a '@'
; preprocessor directives have effect on the assembly source code

; including a source file
.inc "defs.asm" ; the effect of this line is as if the whole file would replace this line 
.inc "templates.asm"

.org 0xff00           ; set location counter to 0xff00
start:                ; define a label at location 0xff00
nop                   ; opcode 
start_1: nop          ; labels can be also inline 
start_2: start_3: nop ; this is also supported


.org RST_VEC_ORG ; set location counter to the value of RST_VEC_ORG (defined in defs.asm)
jmp start        ; jump instruction, jump to location of 'start' label

; to define explicit data use .db (data byte) and .dw (data word)
.org CONST_DATA_LOC ; set location counter to the value of CONST_DATA_LOC
.db 0xaa   ; this will write 0xaa in the ROM at location   CONST_DATA_LOC+0
.db 0x10   ; this will write 0x  10 in the ROM at location CONST_DATA_LOC+1 
.dw 0x10   ; this will write 0x0010 in the ROM at location CONST_DATA_LOC+2 
.dw 0xaaaa ; this will write 0xaaaa in the ROM at location CONST_DATA_LOC+4
; db writes 1 byte and increments the location counter by +1 
; dw writes 2 bytes and increments the location counter by +2 

; more than one value can be written in one line 
.db 0x10, 0x20, 0x30    ; this will increment the location counter by +3
.dw 0x10, 0x100, 0x1000 ; this will increment the location counter by +6

; data can be also provided in string format 
.db "test string" ; this will write the ascii codes, 0x74, 0x65, ... etc.
; the upper line will also increment the location counter by the number of characters 

.dw "123" ; will write 0x31, 0x32, 0x33, 0x00  - note the alignment to full words
; the upper line will increment the location counter by full words, in this case 4 bytes

; data can be defined also in blocks 
.db (
  "some test string", 0x00 
  "other\nstring", 0x00
)

.org TEST_INSTR_LOC
  plp                              ; opcode only instruction
  rol a                            ; opcode with register operand
  ldai 0x11                        ; load immediate  alt to lda #$11
  lda 0xaa                         ; zero page addressing mode
  lda 0xaaaa                       ; absolute addressing mode
  lda [0xaa,  x]                   ; zero page indexed indirect mode
  lda [0xaa], y                    ; zero page indirect indexed with Y mode
  lda [0xbb]                       ; zero page indirect mode
  lda 0xcc,   x                    ; zero page indexed with X mode
  lda 0xcc,   y                    ; zero page indexed with Y mode
  lda 0xccaa, x                    ; absolute indexed with X mode
  lda 0xccaa, y                    ; absolute indexed with Y mode
  bbr 0,      0x10, TEST_INSTR_LOC ; branch to TEST_INSTR_LOC if bit 0 of zero page location 0x10 is cleared
  smb 1,      0x20                 ; set bit 1 of zero page location 0x20

.org VAR_DEF_LOC  
; variables can be defined with the combination of labels and .byte / .word directives 
; a .byte or .byte n directive just increments the location counter 
my_var_1: .byte   ; my_var_1 points to addr VAR_DEF_LOC
my_var_2: .byte 2 ; my_var_2 points to addr VAR_DEF_LOC+1
my_var_3: .byte   ; my_var_3 points to addr VAR_DEF_LOC+3

my_other_var_1: .word   ; my_other_var_1 points to addr VAR_DEF_LOC+4
my_other_var_2: .word 2 ; my_other_var_2 points to addr VAR_DEF_LOC+6 , note that .word 2 has a size of 4
my_other_var_3: .word   ; my_other_var_3 points to addr VAR_DEF_LOC+10

; and variables can be used in the code 
.org 0x8080
  ldai 0xaa 
  sta my_var_1
  ; and event jumped to !!! 
  jmp my_var_1 ; however this should be rarely useful


; math expressions can be embedded into any place that accepts a number or (tbd string)
.org 0x8100
.db 0xa0 + 0x0a ; = 0xaa 
.db 0xaa - 0xa0 ; = 0x0a 
.db (2 + 2)*2   ; = 8 
.db 2 + 2 * 2   ; = 6 
.db 2 << 2      ; = 8 
.db 10 >> 1     ; = 5 
.db 2*2/2       ; = 2 
.db 0xaa ^ 0xff ; = 0x55 
.db 0xaa | 0x55 ; = 0xff 
.db 0xaa & 0x55 ; = 0x00 
.db 0x1234.l    ; = 0x34 
.db 0x1234.h    ; = 0x12 
.dw ~0xaaaa     ; = 0x5555 
.dw ~0x00       ; = 0xffff 
.db ~0xaa.l     ; = 0x55 
.db ~0xaa.h     ; = 0xff 
.db -(-10)      ; = 10


.db ~((0x11a9+(3>>1)).l>>1*2+1).l ; = 0xea

; definitions can refer to labels and other definitions
; definitions can also represent register - however this might not be practical
.def (
  DEF_WITH_LABEL := some_label + 0x10
  REG_DEF        := x 
  DEF_WITH_DEF   := DEF_WITH_LABEL - some_label
)

some_label:
  lda DEF_WITH_DEF, REG_DEF

; some string operation are also supported
.def (
  SOME_STR := "some" + " " + "str" ; will concatenate strings
  OTER_STR := SOME_STR + "."
  REP_STR := "\x00"*10 ; will repeat the 0x00 byte 10 times
)

.db "\x00"*20 ; repeat 0x00 byte 20 times

; using a array item is simple as 
.db TEST_ARRAY[0] ; will use the first item of the TEST_ARRAY
.db TEST_ARRAY_3[TEST_ARRAY[0]] ; item indexes can also be expressions
jmp TEST_ARRAY_3[0] ; this will jump to the address stored in the first item of the TEST_ARRAY_3
.db TEST_ARRAY_3.len ; the length of the array is also available
.db TEST_ARRAY_3[TEST_ARRAY_3.len -1] ; this will use the last item of the TEST_ARRAY_3

; execute a defined template
@other_template("test string")
