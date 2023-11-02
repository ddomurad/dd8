#include "emulator.h"
#include "uart.h"

#include <stdint.h>
#include <string.h>
#include <stdio.h>
#include <avr/pgmspace.h>

char send_buffer[15];
uint8_t _read(uint16_t addr);
void _write(uint16_t addr, uint8_t data);

uint8_t RAM_DATA[0x400];
uint8_t ROM_DATA[0x400];

void emu_init() {
  memset((void*)send_buffer, 0, 15);
  memset((void*)RAM_DATA, 0, 0x400);
  memset((void*)ROM_DATA, 0, 0x400);

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
    return 0x10;

  if(addr < 0x400)
    return RAM_DATA[addr];

  if(addr >= 0x1000 && addr < 0x1400)
    return ROM_DATA[addr - 0x1000];
  return 0xea;
}

void _write(uint16_t addr, uint8_t data) {
  uint8_t addr_h = addr >> 8;
  uint8_t addr_l = addr & 0xff;

  if(addr_h == 0x00)
    RAM_DATA[addr_l] = data;
}


uint8_t *emu_rom_ptr() {
  return ROM_DATA;
}

void emu_dump_rom(uint16_t addr, uint16_t len) {
  uint16_t endAddr = addr + len;
  if(endAddr > 0x1400) 
    endAddr = 0x1400; 

  sprintf(send_buffer, "ROM DUMP:\n");
  uart_read(send_buffer, 15);

  for(uint16_t i = addr; i<endAddr; i++) {
    sprintf((char*)send_buffer, "%04X %02X\n", i, ROM_DATA[i-0x1000]);
    uart_write(send_buffer, 15);
  }
}

void emu_dump_ram(uint16_t addr, uint16_t len) {
  uint16_t endAddr = addr + len;
  if(endAddr > 0x400) 
    endAddr = 0x400; 

  sprintf(send_buffer, "RAM DUMP:\n");
  uart_read(send_buffer, 15);

  for(uint16_t i=addr; i<endAddr; i++) {
    sprintf((char*)send_buffer, "%04X %02X\n", i, RAM_DATA[i]);
    uart_write(send_buffer, 15);
  }
}

