#include "def.h"
#include "uart.h"

#include <stdint.h>
#include <stdio.h>
#include <string.h>

#include <avr/interrupt.h>
#include <avr/io.h>
#include <avr/pgmspace.h>
#include <util/delay.h>
#include <util/delay_basic.h>

#include "./gfx_data.h"

void write_byte(uint16_t addr, uint8_t data) {
  ADDR_H_PORT = (uint8_t)(addr >> 8);
  ADDR_L_PORT = (uint8_t)(addr & 0xff);
  DATA_PORT = data;

  CTRL_PORT &= ~(1 << CTRL_WE_B_PIN); // go low
  CTRL_PORT |= (1 << CTRL_WE_B_PIN);  // go high
}

void write_word(uint16_t addr, uint16_t data) {
  write_byte(addr, data & 0xff);
  write_byte(addr + 1, (data >> 8) & 0xff);
}

int main() {
  CTRL_DDR = (1 << CTRL_CRAM_CS_B_PIN) | (1 << CTRL_VRAM_CS_B_PIN) |
             (1 << CTRL_WE_B_PIN);
  CTRL_PORT = (1 << CTRL_CRAM_CS_B_PIN) | (1 << CTRL_VRAM_CS_B_PIN) |
              (1 << CTRL_WE_B_PIN);

  ADDR_H_DDR = 0xff;
  ADDR_L_DDR = 0xff;
  DATA_DDR = 0xff;

  ADDR_H_PORT = 0x00;
  ADDR_L_PORT = 0x00;
  DATA_PORT = 0x00;

  uart_init();
  uart_writeln("READY");
  // uart_writeln("wait for start");
  // uart_read();

  uart_writeln("setup palette ram");
  CTRL_PORT &= ~(1 << CTRL_CRAM_CS_B_PIN);
  uint16_t addr = 0x0000;
  for (addr = 0x0000; addr < 0x400; addr++) {
    if ((addr >> 8) == 0x01) {
      write_byte(addr, ~addr);
    } else {
      write_byte(addr, addr & 0xff);
    }
  }
  CTRL_PORT |= (1 << CTRL_CRAM_CS_B_PIN);

  uart_writeln("setup vga ram");
  CTRL_PORT &= ~(1 << CTRL_VRAM_CS_B_PIN);
  for (addr = 0x0000; addr < 0x8000; addr++) {
    if ((addr & 0xff00) == 0x00) {
      write_byte(addr, 0b00000111);
      continue;
    }

    if ((addr & 0xff00) == 0xff) {
      write_byte(addr, 0b00111111);
      continue;
    }

    if ((addr & 0xff) == 0) {
      write_byte(addr, 0b11000111);
      continue;
    }

    if ((addr & 0xff) == 1) {
      write_byte(addr, 0b11111111);
      continue;
    }

    if ((addr & 0xff) == 0xff) {
      write_byte(addr, 0b00111000);
      continue;
    }

    if ((addr & 0xff) == 0xfe) {
      write_byte(addr, 0xff);
      continue;
    }

    write_byte(addr, 0x00);

    // write_byte(addr, addr + (addr  >> 8));
    // if(((addr & 0xff) == 64) || ((addr & 0xff) == 128 + 64)){
    //   write_byte(addr, 0b11001100);
    //   continue;
    // }
  }
  CTRL_PORT |= (1 << CTRL_VRAM_CS_B_PIN);

  uart_writeln("done");
  for (;;) {
    asm volatile("nop"); // we dont need this :)
  }
}
