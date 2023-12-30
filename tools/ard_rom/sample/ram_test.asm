
.org 0xf000 
  ldai 0x00 
  ldxi 0x00 

inf_loop:
  sta 0x00, x
  inx
  inc a
  beq read_loop
  bra inf_loop
read_loop: 
  lda 0x00, x
  sta 0x1000, x 
  inx
  bra read_loop

.org 0xfffc
.db 0x00, 0xf0
