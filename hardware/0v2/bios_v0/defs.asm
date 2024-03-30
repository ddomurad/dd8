.def (
  ; memory regions
  PRG_ENTER := 0xf000
  ZERO_PAGE := 0x0000

  ; io  
  UART      := 0x8000
  UART_STS  := UART + 0x01
  UART_CMD  := UART + 0x02
  UART_CTR  := UART + 0x03

  ; int. vectors
  RST_VEC := 0xfffc
)

.org ZERO_PAGE
TMP_VAR_0: .byte
TMP_VAR_1: .byte
TMP_VAR_2: .byte

LOADER_STATE: .byte ; 0 - read low addr, 1 - read high addr, 2 - read byte
LOAD_ADDR: .word
BYTES_LOADED: .byte 
NIBLE_LOAD: .byte
TMP_BYTE: .byte

