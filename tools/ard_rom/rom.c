#include "rom.h"
#include <stdint.h>
#include <string.h>

uint8_t ROM_DATA[ROM_SIZE];

void rom_init() { memset(ROM_DATA, 0x00, ROM_SIZE); }

uint8_t *rom_buffer() { return ROM_DATA; }

uint8_t rom_read(uint16_t addr) {
  addr -= ROM_START_ADDR; 
  if(addr >= ROM_SIZE) {
    return 0;
  }

  return ROM_DATA[addr];
}
