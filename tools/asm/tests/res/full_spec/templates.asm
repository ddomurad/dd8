.org TMPL_TEST_LOC

.tmpl some_template (i,j,k) { ; this defines a template with 3 arguments
  ldai i ; arguments are actually just local definitions
  sta j
  .db k
  ldai CONST_1 ; of course, you can use golo constants
}

.tmpl other_template (var) {
  .db var
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


.tmpl nested_template_with_labels (arg) {
  .tmpl sub_template (arg2) {
    loop: ; this label is not visible outside the sub template
      ldai arg2
      jmp loop ; infinite loop, just and example
  }

  loop: ; this label is not visible outside the template
    @sub_template (arg*2) ; use sub template
    jmp loop ; infinite loop
}

@nested_template_with_labels (0x10) ; use the template


.rep i, 3, 8 {
  .db i, i*2
  ldai i
}
