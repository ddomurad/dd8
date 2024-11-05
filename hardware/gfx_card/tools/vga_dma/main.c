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

uint8_t uart_read_hex_byte() {
  char z[3];

  z[0] = uart_read();
  z[1] = uart_read();
  z[2] = 0x00;

  uint16_t x = 0;
  sscanf(z, "%2X", &x);
  return x;
}

uint16_t uart_read_hex_word() {
  uint8_t b1 = uart_read_hex_byte();
  uint8_t b2 = uart_read_hex_byte();
  return (b1 << 8) | b2;
}

int main() {
  CTRL_DDR = (1 << CTRL_DMA_CS_B_PIN) | (1 << CTRL_WE_B_PIN);
  CTRL_PORT = (1 << CTRL_DMA_CS_B_PIN) | (1 << CTRL_WE_B_PIN);

  ADDR_H_DDR = 0xff;
  ADDR_L_DDR = 0xff;
  DATA_DDR = 0xff;

  ADDR_H_PORT = 0x00;
  ADDR_L_PORT = 0x00;
  DATA_PORT = 0x00;

  uint16_t src_addr = 0;
  uint16_t dst_addr = 0;
  uint8_t cpy_width = 0;
  uint8_t cpy_height = 0;
  uint8_t cpy_mask = 0;
  uint8_t cpy_command = 0;

  uart_init();
  uint8_t sep;
  char prt[] = "0000 0000 00 00 00 00\0\0";
  uint8_t i=0;
  for (;;) {
    for(i=0;i<100;i++) {
      uart_read_nb();
    }

    uart_write("ready: ");
    src_addr = uart_read_hex_word();
    sep = uart_read();
    if (sep != ' ') {
      uart_writeln("input error! A");
      sprintf(prt, "ch: %c", sep);
      uart_writeln(prt);
      continue;
    }
    dst_addr = uart_read_hex_word();
    sep = uart_read();
    if (sep != ' ') {
      uart_writeln("input error! B");
      continue;
    }
    cpy_width = uart_read_hex_byte();
    sep = uart_read();
    if (sep != ' ') {
      uart_writeln("input error! C");
      continue;
    }
    cpy_height = uart_read_hex_byte();
    sep = uart_read();
    if (sep != ' ') {
      uart_writeln("input error! D");
      continue;
    }
    cpy_mask = uart_read_hex_byte();
    sep = uart_read();
    if (sep != ' ') {
      uart_writeln("input error! E");
      continue;
    }
    cpy_command = uart_read_hex_byte();
    sep = uart_read();
    if (sep != '\r') {
      uart_writeln("input error! F");
      sprintf(prt, "ch: %c", sep);
      uart_writeln(prt);
      continue;
    }

    // setup DMA
    sprintf(prt, "%4X %4X %2X %2X %2X %2X", src_addr, dst_addr, cpy_width, cpy_height, cpy_mask, cpy_command);
    uart_write("exec:  ");
    uart_writeln(prt);
    uart_writeln("");

    CTRL_PORT &= ~(1<<CTRL_DMA_CS_B_PIN);
    write_word(0x0000, src_addr);
    write_word(0x0002, dst_addr);
    write_byte(0x0004, cpy_width);
    write_byte(0x0005, cpy_height);
    write_byte(0x0006, cpy_mask);
    write_byte(0x0007, cpy_command);
    CTRL_PORT |= (1<<CTRL_DMA_CS_B_PIN);
  }
}
