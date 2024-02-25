#include "def.h"
#include "uart.h"

#include <stdint.h>
#include <stdio.h>
#include <string.h>

#include <avr/interrupt.h>
#include <avr/io.h>
#include <util/delay.h>
#include <util/delay_basic.h>

void init_ports() {
  ADDR_L_DDR = 0x00;
  ADDR_H_DDR = 0x00;
  DATA_DDR = 0x00;
  CTRL_DDR = 0x00;

  CLK_DDR &= ~(1 << CLK_PIN);
  CLK_PORT &= ~(1 << CLK_PIN);
}

uint8_t last_clk = 0xff;
uint8_t clk_sub_phase = 0x01;

char send_buffer[64];
void write_status(uint8_t addrh, uint8_t addrl, uint8_t data, uint8_t ctrl, uint8_t clk) {
  sprintf(send_buffer, "%02X%02X %02X %02X %02X", addrh, addrl, data, ctrl, clk);
  uart_writeln(send_buffer, 64);
}

int main() {
  uart_init();
  uart_writeln("ready...", 8);

  while (1) {
    uint8_t clk = CLK_PINS; 
    clk = clk & (1<<CLK_PIN);
    if (last_clk != clk) {
      if (clk || clk_sub_phase) {
        write_status(ADDR_H_PINS, ADDR_L_PINS, DATA_PINS, CTRL_PINS & 0x03, clk);
      }
    }

    last_clk = clk;
  }
}
