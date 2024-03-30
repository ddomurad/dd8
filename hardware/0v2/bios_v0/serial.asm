serial_init: {
  stz UART_STS    ; status register, soft reset 

  ldai 0b00011110 ; 1 stop bit, 8 bits, 9600 baud
  sta UART_CTR    ; ctrl reg

  ldai 0b00001011 ; DTR = 0;
  sta UART_CMD    ; cmd register
  rts
}

; write TMP_VAR_0 to serial 
serial_write_char: {
  lda TMP_VAR_0
  sta UART
  jsr _write_char_delay
  rts
}

_write_char_delay: {
  ; todo: delay 
  rts
}

; read TMP_VAR_0 from serial 
; set TMP_VAR_1 to 0xff when read, otherwise to 0x00
serial_read_char: {
  ldai 0x08
  and UART_STS
  beq no_data ; return if no data available
  clc 
  lda UART
  sta TMP_VAR_0
  ldai 0xff
  sta TMP_VAR_1
  jmp return
no_data:
  stz TMP_VAR_1
return:
  rts
}
