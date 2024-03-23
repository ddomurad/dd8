.def (
  ENTER   := 0xf000
  RST_VEC := 0xfffc

  UART_ADDR := 0x8000
  UART_STS  := UART_ADDR + 0x01
  UART_CMD  := UART_ADDR + 0x02
  UART_CTR  := UART_ADDR + 0x03

  VARIABLES := 0x0200
)

; .org 0x0000
; jmp 0xf000

.org VARIABLES
DATA_BUFFER: .byte 0xff 
DATA_POINTER: .byte 
LAST_CHAR: .byte 

.org ENTER 
initialize_uart:
  ldai 0x00 
  sta UART_STS    ; status register, soft reset 

  ldai 0b00011110 ; 1 stop bit, 8 bits, 9600 baud
  sta UART_CTR    ; ctrl reg

  ldai 0b00001011 ; DTR = 0;
  sta UART_CMD    ; cmd register

init_ram:
  stz DATA_POINTER
; init_test_ram:
;   ldxi 0x00
;   ldai 0x64
;   sta DATA_BUFFER,x
;   inx
;   sta DATA_BUFFER,x
;   inx
;   sta DATA_BUFFER,x
;   inx
;   sta DATA_BUFFER,x
;   inx
;   ldai 0x0a 
;   sta DATA_BUFFER,x
;   ldai 0x0a
;   sta LAST_CHAR 
;   ldai 0x05
;   sta DATA_POINTER

loop:
  jsr read_serial_st
  ldai 0x0a ; new line 
  sec
  sbc  LAST_CHAR
  beq call_send
jmp loop
call_send:
  jsr send_serial_st
jmp loop

read_serial_st:
  ldai 0x08 ; uart receive bit
  and UART_STS
  beq read_serial_ret ; if not set just return 
  clc 
  lda UART_ADDR 
  sta LAST_CHAR
  ldx DATA_POINTER 
  sta DATA_BUFFER, x
  inx 
  stx DATA_POINTER
read_serial_ret:
rts

send_serial_st:
  ldxi 0x00 
send_serial_st_loop: 
  txa 
  inc a
  inc a
  clc
  sbc DATA_POINTER
  beq send_serial_ret
  lda DATA_BUFFER, x
  sta UART_ADDR
  inx 
  jmp send_serial_st_loop
send_serial_ret:
rts

delay_sr:
rts

.org RST_VEC
.db ENTER.l, ENTER.h


