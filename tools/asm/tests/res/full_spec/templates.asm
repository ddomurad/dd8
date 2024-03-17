.org TMPL_TEST_LOC
; there are different types of templates 
; .tmpl - act as custom blueprints, that can be later executed or "expanded"
; .rep - repeats the template body defined number of times 

.tmpl some_template (i,j,k) { ; this defines a template with 3 arguments
  ldai i ; arguments are actually just local definitions
  sta j
  .db k
  ldai CONST_1 ; you can also use global constants
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

; you cal also use expressions as template arguments
@some_template (TEST_NUMBER_1*2, RST_VEC_ORG.l, TEST_STR_1) 
; this will be expanded to:
; ldai 200 
; sta 0xfc
; .db "this is\n\ta test \"string\"" 
; ldai 0x10  ; CONST_1


; templates can be nested, however the sub_template is not visible outside 
.tmpl nested_template_with_labels (arg) {
  .tmpl sub_template (arg2) {
    loop: ; this label is not visible outside the sub_template
      ldai arg2
      jmp loop ; infinite loop, just an example
  }

  loop: ; this label is not visible outside this template, and not visible inside the sub_template
    @sub_template (arg*2) ; use sub template
    jmp loop ; infinite loop, just an example
}

@nested_template_with_labels (0x10) ; use the template

; this is a repeat template 
; i is the iterator variable
; 3 is the start value
; and 8 is the end value (included)
; currently the repeat template will always count from start to end, 
; incrementing the iterator variable by +1 or -1 (if start < end)
.rep i, 3, 8 { 
  .db i, i*2
  ldai i
}
; this will be expanded to:
;.db 3, 6
;ldai 3
;.db 4, 8
;ldai 4
;.db 5, 10
;ldai 5
;.db 6, 12
;ldai 6
;.db 7, 14
;ldai 7
;.db 8, 16
;ldai 8

.rep i, 5, 0 {
  ldai i*2
}
; this will expand to:
; ldai 10 
; ldai 8
; ldai 6
; ldai 4
; ldai 2 
; ldai 0

; templates can be also mixed 
.tmpl nested_template_with_repeat_and_labels (st, en) {
  .tmpl sub_template(s1) {
    loop:
    ldai s1 
    jmp loop
  }

  .rep i, st, en {
    @sub_template (i)
  }
}

@nested_template_with_repeat_and_labels (0, 2)

