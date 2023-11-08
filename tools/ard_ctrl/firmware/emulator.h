#ifndef __EMULATOR_H
#define __EMULATOR_H

#include <stdint.h>

void    emu_init();

uint8_t emu_read_monitor(uint16_t addr, uint8_t data);
void    emu_write_monitor(uint16_t addr, uint8_t data);

uint8_t emu_read(uint16_t addr);
void    emu_write(uint16_t addr, uint8_t data);

uint8_t *emu_rom_ptr();

void emu_dump_ram(uint16_t addr, uint16_t len);
// void emu_dump_rom(uint16_t addr, uint16_t len);

#endif // !__EMULATOR_H
