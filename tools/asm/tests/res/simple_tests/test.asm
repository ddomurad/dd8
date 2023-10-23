; this is an nonsense asm program
; used only for testing the final binary output

php
php ; some test comment

ldai 0x10         ; load 16
ldai 16           ; load 16
ldai 1_00       ; load 16
ldai 0b00010000   ; load 16
ldai 0b0001_0000  ; load 16

lda 0x55   ; zerop page load 
lda 0x1111 ; absolute loade 
lda 0x41, x 
lda 0x4100, x 
lda 0x1133, y

