; ---------------------------------------
;   Simulation const 
; ---------------------------------------

; .org 0x0000
; .db "3 4 +5 6 ++"
; test_expr_end:
; .org 0xc000
; .db 0x00, 0x01


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
  LCD_CMD_MOVE_CURSOR_LEFT := 0x10

  KBR_ADDR := 0xc000

  CMD_BS := 0xfe 
  CMD_ENTER := 0xff
)

; ---------------------------------------
;   Variables 
; ---------------------------------------
.org ZPAGE
exp_buff: .byte 16 ;100

kbr_key:       .byte 1
kbr_rd_offset: .byte 1
exp_buff_ptr: .byte 2

tmp1: .byte 1
; ---------------------------------------
;   Routine 
; ---------------------------------------
.org ENTER 
  jsr init_buffer
  jsr init_lcd

main_loop:
  jsr handle_keys
bra main_loop

init_buffer:
  ldai (exp_buff - 0x01).l
  sta exp_buff_ptr ; initialize buffer pointer
  ldai 0x00 
  sta exp_buff_ptr+1
rts 

init_lcd:
  ldai  LCD_CMD_ENABLE     ; enable
  sta LCD_CMD_ADDR 
  ldai LCD_CMD_RSET_DDRAM  ; set DDRAM = 0x00
  sta LCD_CMD_ADDR 
  ldai LCD_CMD_CLR_DISPLAY ; clear display
  sta LCD_CMD_ADDR 
rts

handle_keys:
  jsr read_keys 
  beq handle_keys_ret
check_cmd_key:
  ldai 0x08
  sec
  sbc kbr_key
  beq handle_back_space
push_key_on_buffer:
  lda kbr_key
  jsr map_key
  inc exp_buff_ptr
  sta [exp_buff_ptr]
  jsr print_to_lcd
  jsr delay_kbr
  rts
handle_back_space:
  ldai 0xff 
  sec 
  sbc exp_buff_ptr
  beq handle_keys_ret
  dec exp_buff_ptr
  jsr remove_last_char
  jsr delay_kbr
handle_keys_ret:
rts

print_to_lcd:
  lda [exp_buff_ptr]
  sta LCD_DATA_ADDR
rts

remove_last_char:
  ldai LCD_CMD_MOVE_CURSOR_LEFT
  sta LCD_CMD_ADDR 
  ldai 0x20
  sta LCD_DATA_ADDR
  ldai LCD_CMD_MOVE_CURSOR_LEFT
  sta LCD_CMD_ADDR 
rts 

map_key:
  sta tmp1 
map_key_1:
  sec
  ldai 0b0000_0001
  sbc tmp1 
  bne map_key_2 
  ldai 0b0011_0001 
  rts
map_key_2:
  sec
  ldai 0b0001_0001
  sbc tmp1 
  bne map_key_3 
  ldai 0b0011_0010 
  rts
map_key_3:
  sec
  ldai 0b0010_0001
  sbc tmp1 
  bne map_key_4 
  ldai 0b0011_0011 
  rts
map_key_4:
  sec
  ldai 0b0000_0010
  sbc tmp1 
  bne map_key_5 
  ldai 0b0011_0100 
  rts
map_key_5:
  sec
  ldai 0b0001_0010
  sbc tmp1 
  bne map_key_6 
  ldai 0b0011_0101 
  rts
map_key_6:
  sec
  ldai 0b0010_0010
  sbc tmp1 
  bne map_key_7 
  ldai 0b0011_0110 
  rts
map_key_7:
  sec
  ldai 0b0000_0100
  sbc tmp1 
  bne map_key_8 
  ldai 0b0011_0111
  rts
map_key_8:
  sec
  ldai 0b0001_0100
  sbc tmp1 
  bne map_key_9 
  ldai 0b0011_1000 
  rts
map_key_9:
  sec
  ldai 0b0010_0100
  sbc tmp1 
  bne map_key_0
  ldai 0b0011_1001 
  rts
map_key_0:
  sec
  ldai 0b0001_1000
  sbc tmp1 
  bne map_key_add
  ldai 0b0011_0000
  rts
map_key_add:
  sec
  ldai 0b0011_0001
  sbc tmp1 
  bne map_key_sub
  ldai 0b0010_1011 
  rts
map_key_sub:
  sec
  ldai 0b0011_0010
  sbc tmp1 
  bne map_key_mult
  ldai 0b0010_1101 
  rts
map_key_mult:
  sec
  ldai 0b0011_0100
  sbc tmp1 
  bne map_key_div
  ldai 0b0010_1010 
  rts
map_key_div:
  sec
  ldai 0b0011_1000
  sbc tmp1 
  bne map_key_space
  ldai 0b0010_1111 
  rts
map_key_space:
  sec
  ldai 0b0010_1000
  sbc tmp1 
  bne map_key_no_key
  ldai 0x20
  rts
map_key_no_key:
  ldai 0x00
rts

read_keys:
  stz kbr_rd_offset
  jsr read_key ;+0
  ; lda kbr_key 
  bne read_keys_ret

  inc kbr_rd_offset 
  jsr read_key ;+1
  ; lda kbr_key 
  bne read_keys_ret


  inc kbr_rd_offset 
  jsr read_key ;+2
  ; lda kbr_key 
  bne read_keys_ret

  inc kbr_rd_offset 
  jsr read_key ;+3
  ; lda kbr_key 
  bne read_keys_ret
read_keys_ret:
rts 

read_key:
  ldx kbr_rd_offset 
  lda KBR_ADDR,x
  sta kbr_key 
  beq read_key_ret
  clc
  txa 
  rol a 
  rol a
  rol a
  rol a
  ora kbr_key 
  sta kbr_key
read_key_ret:
rts

delay_kbr:
  ldai 0xff 
  sta tmp1
delay_kbr_loop:
  dec tmp1 
  bne delay_kbr_loop
rts


.org RST_VTR
.db ENTER.l, ENTER.h



