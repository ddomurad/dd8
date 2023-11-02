#include "hex_decode.h"
#include "uart.h"
#include <stdlib.h>


uint8_t h2i(char c) {
  if(c <= '9')
    return c - '0';
  return c - 'A' + 10;
}

inline uint8_t read_8() {
  uint8_t d = 0x00;
  d = h2i(uart_read());
  d <<= 4;
  d |= h2i(uart_read());
  return d;
}

inline uint16_t read_16() {
  uint16_t d = 0x00;
  d = h2i(uart_read());
  d <<= 4;
  d |= h2i(uart_read());
  d <<= 4;
  d |= h2i(uart_read());
  d <<= 4;
  d |= h2i(uart_read());
  return d;
}

uint8_t hex_read_mem(uint8_t *buffer, uint16_t start_addr, uint16_t end_addr) {
  uint16_t dsize = 0x00;
  uint16_t addr = 0x00;
  uint8_t cs = 0x00;
  uint8_t rb;

  while(1) {
    while(1) {
      if(uart_read() == ':')
        break;
    }

    dsize = read_8();
    addr = read_16();
    cs = dsize;
    cs += (uint8_t)addr;
    cs += (uint8_t)(addr>>8);

    if(addr > end_addr) 
      return 1;
    
    addr -= start_addr;

    if(uart_read() != '0')
      return 2;

    rb = uart_read();
    if(rb == '1')
      return 0;
    if(rb != '0')
      return 3;

    while(dsize > 0) {
      rb = read_8();
      buffer[addr++]  = rb;
      dsize--;
      cs += rb;
    }

    cs += read_8();
    if(cs != 0) 
      return 4;
  }  
}
