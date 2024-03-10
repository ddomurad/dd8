.org TMPL_TEST_LOC

.tmpl some_template (i,j,k) { ; this defines a template with 3 arguments
  ldai i ; arguments are actually just local definitions
  sta j
  .db k
  ldai CONST_1 ; of course, you can use golo constants
}


@some_template (1,2,3) ; this is how you use a template
; this will be expanded to:
; ldai 1 
; sta 2 
; .db 3 
; ldai 0x10  ; CONST_1

; you cal also use expiresion as template arguments
@some_template (TEST_NUMBER_1*2, RST_VEC_ORG.l, TEST_STR_1) 
; this will be expanded to:
; ldai 200 
; sta 0xfc
; .db "this is\n\ta test \"string\"" 
; ldai 0x10  ; CONST_1

