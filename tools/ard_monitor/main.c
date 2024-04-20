#include "def.h"
#include "uart.h"

#include <stdint.h>
#include <stdio.h>
#include <string.h>

#include <avr/interrupt.h>
#include <avr/io.h>
#include <util/delay.h>
#include <util/delay_basic.h>

#define CLK_SUB_PHASE 1
#define CLK_STEPS 3

void init_ports() {
  ADDR_L_DDR = 0x00;
  ADDR_H_DDR = 0x00;
  DATA_DDR = 0x00;
  CTRL_DDR = 0x00;

  CLK_DDR |= (1 << CLK_PIN);
  CLK_PORT &= ~(1 << CLK_PIN);
}

uint8_t clk = 0;
uint8_t loop_i = 0;

char send_buffer[64];

void write_status(uint8_t addrh, uint8_t addrl, uint8_t data, uint8_t ctrl,
                  uint8_t clk) {
  sprintf(send_buffer, "%02X%02X %02X %02X %02X", addrh, addrl, data, ctrl,
          clk);
  uart_writeln(send_buffer, 64);
}

int main() {
  init_ports();
  uart_init();
  uart_writeln("ready...", 8);

  while (1) {
    for (loop_i = 0; loop_i < CLK_STEPS; loop_i++) {
      write_status(ADDR_H_PINS, ADDR_L_PINS, DATA_PINS, CTRL_PINS & 0x1f, clk);
    }

    if (clk) {
      CLK_PORT &= ~(1 << CLK_PIN);
      clk = 0x00;
    } else {
      CLK_PORT |= (1 << CLK_PIN);
      clk = 0x01;
    }
  }
}
