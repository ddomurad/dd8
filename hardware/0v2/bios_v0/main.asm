.inc "./defs.asm"
.inc "./sim_setup.asm"

.org PRG_ENTER 
clc

jsr serial_init
jsr init_mem_loader

; ldai 0x22
; sta LOAD_ADDR
; ldai 0x33
; sta LOAD_ADDR+1
; ldai 0x02
; sta LOADER_STATE
; ldai '!'
; sta UART


loop: {
  jsr serial_read_char
  lda TMP_VAR_1
  beq loop
  ; read_chat read the char to TMP_VAR_0 
  ; write_char writes the char from TMP_VAR_0
  jsr serial_write_char ; just an software echo
  ; handle_mem_load will parse TMP_VAR_0
  jsr handle_mem_load
}
jmp loop

.inc "./serial.asm"
.inc "./mem_loader.asm"

.org RST_VEC
.db PRG_ENTER.l, PRG_ENTER.h

; .org 0x3322
; .db "\xea"*10
