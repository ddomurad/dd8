; simple program to test the assembly language

.org 0x8000 ; b1000_0000
  clc
  ldai 0x03 
  sta 0x00 
  ldai 0x0b 
  sta 0x01
  stz 0x02

loop:
  ldai 0x00
  clv 
  adc 0x01 
  adc 0x02 
  sta 0x02  

  dec 0x00 
  bne loop 
  
  lda 0x02 
  jsr write_msg

inf_loop:
  jmp inf_loop

write_msg:
  sta 0x9000  ;b1001_0000
  rts

.org 0xfffd
.db 0x00, 0x80
