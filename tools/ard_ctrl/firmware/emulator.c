#include "emulator.h"
#include "uart.h"

#include <stdint.h>
#include <string.h>
#include <stdio.h>
#include <avr/pgmspace.h>

char send_buffer[15];
uint8_t _read(uint16_t addr);
void _write(uint16_t addr, uint8_t data);

uint8_t RAM_DATA[256] = {
};

const uint8_t ROM_DATA[256] = {
  0xa9, 0x10, 0x18, 0x69, 0x01, 0x85, 0x00, 0xea,
  0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea,
  0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea,
  0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea,
  0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea,
  0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea,
  0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea,
  0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea, 0xea,
};

void emu_init() {
  memset((void*)send_buffer, 0, 15);
}

uint8_t emu_read(uint16_t addr) {
  uint8_t data = _read(addr);
  sprintf((char*)send_buffer, "r %04X %02X\n", addr,  data);
  uart_write(send_buffer, 15);
  return data;
}

void emu_write(uint16_t addr, uint8_t data){
  sprintf((char*)send_buffer, "W %04X %02X\n", addr,  data);
  uart_write(send_buffer, 15);
  _write(addr,data);
}


uint8_t _read(uint16_t addr) {
  if(addr == 0xfffc) 
    return 0x00;

  if(addr == 0xfffd) 
    return 0x80;

  uint8_t addr_h = addr >> 8;
  uint8_t addr_l = addr & 0xff;

  if(addr_h == 0x80)
    return ROM_DATA[addr_l];

  if(addr_h == 0x00)
    return RAM_DATA[addr_l];

  return 0xea;
}

void _write(uint16_t addr, uint8_t data) {
  uint8_t addr_h = addr >> 8;
  uint8_t addr_l = addr & 0xff;

  if(addr_h == 0x00)
    RAM_DATA[addr_l] = data;
}
