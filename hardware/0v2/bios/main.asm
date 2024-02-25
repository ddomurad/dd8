.def (
  ENTER := 0xe000
  RST_VEC := 0xfffc

  UART_ADDR := 0xc000
)

.org ENTER 
; todo: check interup request setup

initialize_uart:
  ldai 0x00 
  sta UART_ADDR + 0x01 ; status register, soft reset 

  ldai 0b00011110 ; 1 stop bit, 8 bits, 9600 baud
  sta UART_ADDR + 0x03 ; ctrl reg

  ldai 0b00001011 ; DTR = 0;
  sta UART_ADDR + 0x02 ; cmd register

send_loop:
  ldai 0x34 
  sta UART_ADDR 
jmp send_loop

.org RST_VEC
.db ENTER.l, ENTER.h
