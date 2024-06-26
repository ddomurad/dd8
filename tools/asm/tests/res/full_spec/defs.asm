;.def will define preprocessor constants 
; this constants are resolved during assembly of the source code 
; think of it as aliases to some values 

.def CONST_1    := 0x10               ; number constant
.def CONST_2    := 0x1010             ; number constant
.def STR_CONST  := "this is a test"   ; string constant

; constants can be also defined in blocks 
.def (
  OTHER_CONST_1   := 0x00
  OTHER_CONST_2   := 0xaa
  OTHER_CONST_STR := "other"
)

; numbers can be represented in decimal, hexadecimal or binary format 
; numbers can be also augmented with '_' signs to enhance the readability

.def (
  TEST_NUMBER_1 := 100         ; decimal form
  TEST_NUMBER_2 := 0xaa        ; hexadecimal
  TEST_NUMBER_3 := 0xaa_aa     ; hexadecimal with _ grouping
  TEST_NUMBER_4 := 0b1010_0011 ; binary with _ grouping
)

; strings can encode special characters or binary data 
.def (
  TEST_STR_1 := "this is\n\ta test \"string\""  ; \\, \", "\n, \t and \r can be used
  TEST_STR_2 := "this is a byte: \xaa"          ; \xNN represents a byte 
)

; a defined constant can be used during constant definition 
.def NEW_CONST := TEST_STR_2

; this constants are actually used in this test program 
.def RST_VEC_ORG    := 0xfffc
.def CONST_DATA_LOC := 0x8000
.def TEST_INSTR_LOC := 0xf000
.def VAR_DEF_LOC    := 0x0000
.def TMPL_TEST_LOC  := 0x1000


.def i := "some i value" ; preprocessor constants can also be single charas long

.def (
  TEST_ARRAY := <1, 2, 3> ; it is possible to define arrays of constants
  TEST_ARRAY_2 := <TEST_ARRAY[1], NEW_CONST*2> ; arrays can be used in other constant definitions
  TEST_ARRAY_3 := <my_var_1, my_var_2> ; arrays can also contain labels
)
