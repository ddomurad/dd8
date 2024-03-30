.inc "bios.asm"

.def (
  ENTRY := 0x1000
  DATA := 0x2000
  MSG := "Hello, Papupa!\n"
)

.org ENTRY
jsr serial_init

init_send:
ldxi 0x00 

send_loop:
lda MSG_DATA, x
clc
beq init_send
sta TMP_VAR_0
jsr serial_write_char
inx
jmp send_loop


.org DATA 
MSG_DATA: .db MSG, 0x00


; .org 0x9000
; serial_init:
; serial_write_char:
; rts
