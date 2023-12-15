#ifndef __ROM_H
#define __ROM_H

#include <stdint.h>

#define ROM_START_ADDR 0xF000 
#define ROM_SIZE       0x0FFF

void rom_init();
uint8_t *rom_buffer();
uint8_t rom_read(uint16_t addr);

#endif // !DEBUG
