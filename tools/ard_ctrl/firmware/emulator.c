#include "emulator.h"
#include "uart.h"


#include <stdint.h>
#include <string.h>
#include <stdio.h>

char state_buffer[15];

void emu_init() {
  memset((void*)state_buffer, 0, 15);
}

uint8_t emu_read(uint16_t addr) {
  uint8_t data = (addr | 0xff); 
  sprintf((char*)state_buffer, "r %04X %02X\n", addr,  data);
  uart_write(state_buffer, 15);
  return data;
}

void emu_write(uint16_t addr, uint8_t data){
  sprintf((char*)state_buffer, "W %04X %02X\n", addr,  data);
  uart_write(state_buffer, 15);
}
