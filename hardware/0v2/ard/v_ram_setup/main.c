#include "def.h"
#include "uart.h"
#include "img.h"

#include <stdint.h>
#include <stdio.h>
#include <string.h>

#include <avr/interrupt.h>
#include <avr/io.h>
#include <util/delay.h>
#include <util/delay_basic.h>

#define HALT                                                                   \
  do {                                                                         \
  } while (1)

int main() {
  ADDR_DDR_0 = 0x00;
  ADDR_DDR_1 = 0x00;
  ADDR_DDR_2 = 0x00;

  ADDR_PORT_0 = 0x00;
  ADDR_PORT_1 = 0x00;
  ADDR_PORT_2 = 0x00;

  DATA_DDR = 0x00;
  DATA_PORT = 0x00;

  CTRL_DDR = (0x01 << CTRL_EN_PIN) | (0x01 << CTRL_WE_PIN) | (0x01 << CTRL_CE2_PIN);
  CTRL_PORT = (0x01 << CTRL_EN_PIN) | (0x01 << CTRL_WE_PIN);

  uart_init();
  uart_writeln("ready...", 8);

  while (1) {
    uint8_t data_out = CTRL_PIN & (1 << CTRL_DO_PIN);
    if (data_out == 0x00) {
      CTRL_PORT &= ~(0x01 << CTRL_EN_PIN);
      _delay_ms(10);
      data_out = CTRL_PIN & (1 << CTRL_DO_PIN);
      if (data_out == 0x00) {
        uart_writeln("video_disabled", 15);
        break;
      } else {
        uart_writeln("video_NOT_disabled", 20);
        HALT;
      }
    }
  }

  ADDR_DDR_0 = 0xff;
  ADDR_DDR_1 = 0xff;
  ADDR_DDR_2 = 0x01;

  uint16_t x = 0;
  uint16_t y = 0;

  CTRL_PORT |= (1 << CTRL_CE2_PIN);

  for (y=0; y < 480; y++) {
    for (x=0; x < 256; x++) {
      ADDR_PORT_0 = x;
      ADDR_PORT_1 = y;
      ADDR_PORT_2 = 0x00;

      uint16_t addr = (y<<8) | x;
      addr %= 256*16;
      // if (addr > 256*16) {
      //   // uart_writeln("addr too big", 12);
      //   addr = 0x00;
      // }
      uint8_t pix2 = pgm_read_byte(&(img[addr]));

      DATA_DDR = 0xff;
      DATA_PORT = pix2;
      CTRL_PORT &= ~(0x01 << CTRL_WE_PIN);
      _delay_us(10); //probably too long ... 
      CTRL_PORT |= (0x01 << CTRL_WE_PIN);
      DATA_PORT = 0x00;
      DATA_DDR = 0x00;
    }
  }

  uart_writeln("ram data setup", 15);

  ADDR_PORT_0 = 0x00;
  ADDR_PORT_1 = 0x00;
  ADDR_PORT_2 = 0x00;

  ADDR_DDR_0 = 0x00;
  ADDR_DDR_1 = 0x00;
  ADDR_DDR_2 = 0x00;

  DATA_PORT = 0x00;
  DATA_DDR = 0x00;

  CTRL_PORT |= (1 << CTRL_EN_PIN);

  uart_writeln("done...", 8);
  HALT;

  return 0;
}
