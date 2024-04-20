.org 0xa000
  jsr load_aa
  jsr load_bb

loop:
jmp loop


load_aa:
  ldai 0xaa 
  sta 0x8080
rts

load_bb:
  ldai 0xbb
  sta 0x8090
rts

.org 0xfffc
  .db 0x00, 0xa0
