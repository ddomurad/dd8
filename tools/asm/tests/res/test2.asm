; this is an nonsense asm program
; used only for testing the final binary output

php
php ; some test comment

.def (
  DATA_START := 0x1077
  VECTOR_START := 0x2077

  ZERO := 0x00 
  ONE  := 0x01 
  FULL := 0xFF
  NULL := ZERO 
  RX := x
  R1 := RX
)

.org VECTOR_START
ldai ZERO
ldai FULL   
ldai ONE 

.org DATA_START
lda 0x1111
test_label_1:
lda 0x41, RX 
lda DATA_START, R1
lda test_label_1
