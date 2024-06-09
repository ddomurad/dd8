#include "def.h"
#include "uart.h"

#include <stdint.h>
#include <stdio.h>
#include <string.h>

#include <avr/interrupt.h>
#include <avr/io.h>
#include <util/delay.h>
#include <util/delay_basic.h>

void wait_for_h_sync_end() {
  while ((CTRL_PIN & (1 << CTRL_H_SYNC_PIN)) != 0x00) {
    asm("");
  };
  while ((CTRL_PIN & (1 << CTRL_H_SYNC_PIN)) == 0x00) {
    asm("");
  };
}

void wait_for_v_sync() {
  uint16_t cnt = 0;
  while ((CTRL_PIN & (1 << CTRL_V_SYNC_PIN)) != 0x00) {
    asm("");
  };
  while ((CTRL_PIN & (1 << CTRL_V_SYNC_PIN)) == 0x00) {
    asm("");
  };

  for (cnt = 0; cnt < 31; cnt++) {
    wait_for_h_sync_end();
  }
}

uint8_t wait_for_line_change(uint8_t l) {
  while ((CTRL_PIN & (1 << CTRL_LINE_PIN)) == l) {
    asm("");
  }

  return (CTRL_PIN & (1 << CTRL_LINE_PIN));
}

int main() {
  CTRL_DDR = (1 << CTRL_EN_PIN) | (1 << CTRL_RW_PIN) | (1 << CTRL_CE_PIN) |
             (1 << CTRL_CEV_PIN);
  CTRL_PORT = (0 << CTRL_EN_PIN) | (1 << CTRL_RW_PIN) | (1 << CTRL_CE_PIN) |
              (1 << CTRL_CEV_PIN);

  ADDR_DDR_0 = 0xff;
  ADDR_PORT_0 = 0x00;
  ADDR_DDR_1 = 0xff;
  ADDR_PORT_1 = 0x00;
  DATA_DDR = 0xff;
  DATA_PORT = 0x00;

  uint16_t i = 0;

  for (i = 0; i < 1024; i++) {
    ADDR_PORT_0 = i & 0xff;
    ADDR_PORT_1 = (i >> 8) & 0xff;

    // if(i%512 == 0) {
    //   DATA_PORT = 0x01;
    // } else if(i % 512 == 318) {
    //   DATA_PORT = 0x01;
    // } else if(i % 512 == 319) {
    //   DATA_PORT = 0xff;
    // } else {
    //   DATA_PORT = 0x00;
    // }
    //
    DATA_PORT = i;

    CTRL_PORT &= ~(1 << CTRL_CE_PIN);
    CTRL_PORT &= ~(1 << CTRL_RW_PIN);
    _delay_us(1);
    CTRL_PORT |= (1 << CTRL_RW_PIN);
    CTRL_PORT |= (1 << CTRL_CE_PIN);
  }

  CTRL_PORT |= (1 << CTRL_EN_PIN);
  CTRL_PORT &= ~(1 << CTRL_CEV_PIN);

  ADDR_PORT_1 = 0x00;

  uint16_t x = 0;
  uint16_t y = 0;

  // for (;;) {
  //   asm("nop");
  // }

  for (;;) {
    wait_for_v_sync();
    uint8_t l = (CTRL_PIN & (1 << CTRL_LINE_PIN));
    for (y = 0; y < 240; y++) {
      l = wait_for_line_change(l);
      if (l) {
        ADDR_PORT_1 = 0x00;
      } else {
        ADDR_PORT_1 = 0x02;
      }

      CTRL_PORT &= ~(1 << CTRL_CE_PIN);
      for (x = 0; x < 50; x++) {
        ADDR_PORT_0 = x;
        DATA_PORT = x + y;

        CTRL_PORT &= ~(1 << CTRL_RW_PIN);
        CTRL_PORT |= (1 << CTRL_RW_PIN);
      }
      CTRL_PORT |= (1 << CTRL_CE_PIN);
    }
  }
}
