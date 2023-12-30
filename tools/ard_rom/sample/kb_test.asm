

.org 0xf000 
  ldai 0b00001100 ; enable
  sta 0x8000
  ldai 0x80       ; set DDRAM = 0x00
  sta 0x8000 

  ldai 0x01       ; clear display
  sta 0x8000 
  ldai 0x00
inf_loop:
  jsr read_key
  clc
  lda 0x00 
  beq inf_loop
  sta 0x8001
  bra inf_loop

read_key:
  clc
  lda 0xc000
  bne store_key_and_return
  lda 0xc001
  bne store_key_and_return
  lda 0xc002
  bne store_key_and_return
  lda 0xc003
  bne store_key_and_return
store_key_and_return:
  sta 0x00
  rts 
  

.org 0xfffc
.db 0x00, 0xf0
