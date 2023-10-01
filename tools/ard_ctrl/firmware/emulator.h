#ifndef __EMULATOR_H
#define __EMULATOR_H

#include <stdint.h>

void    emu_init();
uint8_t emu_read(uint16_t addr);
void    emu_write(uint16_t addr, uint8_t data);

#endif // !__EMULATOR_H
