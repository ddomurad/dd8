#ifndef __EMULATOR_H
#define __EMULATOR_H

#include <stdint.h>

#define EMU_STATE_STOPED 0
#define EMU_STATE_RUNNING 1

void emu_init();
void emu_start();
void emu_stop(); // will block until stooped
void emu_reset();
void emu_step();

void emu_set_brk(uint16_t addr);
void emu_clr_brk();

uint8_t emu_set_clok(uint8_t v);


#endif // !__EMULATOR_H
