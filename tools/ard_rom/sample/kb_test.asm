; Simple test of the numeric keyboard 
; and LCD display 

; ---------------------------------------
;   Definitions 
; ---------------------------------------

.def (
  ZPAGE := 0x0000
  RST_VTR := 0xfffc
  ENTER   := 0xf000

  LCD_CMD_ADDR   := 0x8000 
  LCD_DATA_ADDR  := 0x8001 

  LCD_CMD_ENABLE := 0b00001100
  LCD_CMD_RSET_DDRAM := 0x80
  LCD_CMD_CLR_DISPLAY := 0x01

  KBR_ADDR := 0xc000
)

.org ZPAGE
key_buffer: .byte 1

; ---------------------------------------
;   Routine 
; ---------------------------------------
.org ENTER 
  ldai  LCD_CMD_ENABLE     ; enable
  sta LCD_CMD_ADDR 
  ldai LCD_CMD_RSET_DDRAM  ; set DDRAM = 0x00
  sta LCD_CMD_ADDR 
  ldai LCD_CMD_CLR_DISPLAY ; clear display
  sta LCD_CMD_ADDR 

  ldai key_buffer
inf_loop:
  jsr read_key
  clc
  lda key_buffer 
  bne inf_loop
  sta LCD_DATA_ADDR
  bra inf_loop

read_key:
  clc
  lda KBR_ADDR + 0
  bne store_key_and_return
  lda KBR_ADDR + 1
  bne store_key_and_return
  lda KBR_ADDR + 2
  bne store_key_and_return
  lda KBR_ADDR + 3
  bne store_key_and_return
store_key_and_return:
  sta key_buffer
  rts 
  

.org RST_VTR
.db ENTER.l, ENTER.h
