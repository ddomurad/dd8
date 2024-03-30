.def (
  ZPAGE := 0x0000
  ENTRY := 0xe000
  LIB_CODE := 0xf000

  ARGS := <0x0000, 0x0001, 0x0002> ; function call args
)

.tmpl call_func (fname, args) {
  .rep i, 0, args.len-1 {
    ldai args[i]
    sta ARGS[i]
  }
  jsr fname
}

.org ZPAGE
; ARG_0: .byte
; ARG_1: .byte
; ARG_2: .byte
; ARG_3: .byte


.org ENTRY
@call_func (test_function, <0x10, 0x20>)

.org LIB_CODE

test_function: {
  ldx ARGS[0] 
  ldy ARGS[1]
  rts
}
