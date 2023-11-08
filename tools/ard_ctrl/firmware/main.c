#include "def.h"
#include "emulator.h"
#include "uart.h"
#include "hex_decode.h"

#include <avr/io.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <util/delay.h>
#include <util/delay_basic.h>

#define CLK_LOW (CLK_PORT &= ~(1<<CLK_PIN))
#define CLK_HIGH (CLK_PORT |= (1<<CLK_PIN))

#define CLK_HALF_DELAY CLK_HALF_DELAY_FUNC

#define ADDR_FULL ((ADDR_H_PIN << 8) | ADDR_L_PIN)
#define IS_WRITE_OP(x) (!(x & (1 << CTRL_RW_PIN)))
#define IS_CHIP_SELECTED(a) ((a >> 8) & 0x80)

inline static void io_setup() {
  // setup clk pin as output
  CLK_DDR |= 1<<CLK_PIN;
  CLK_PORT &= ~(1<<CLK_PIN); //make sure it is set to low at start

  // setup reset pin as output 
  RSTB_DDR |= 1<<RSTB_PIN;
  RSTB_PORT &= ~(1<<RSTB_PIN); //make sure it is set to low at start

  // setup addr ports as input 
  ADDR_H_DDR = 0x00; 
  ADDR_L_DDR = 0x00; 
  ADDR_H_PORT = 0x00; 
  ADDR_L_PORT = 0x00; 

  // setup data port as input 
  DATA_DDR   = 0x00; 
  DATA_PORT   = 0x00; 

  // setup ctrl (rw, cs) port as input
  CTRL_DDR   = 0x00; 
  CTRL_PORT   = 0x00; 

  // 
  for(int8_t i=0; i<4; i++) {
    CLK_HIGH;
    _delay_ms(1);
    CLK_LOW;
    _delay_ms(1);
  }

  // set reset pin high
  RSTB_PORT |= 1<<RSTB_PIN;
  _delay_ms(1);
}

inline static void io_update() {
  uint8_t ctrl_pin_val;
  uint8_t chip_select;
  uint16_t addr;

  CLK_LOW;
  CLK_HALF_DELAY;

  addr = ADDR_FULL;
  ctrl_pin_val = CTRL_PINS;
  chip_select = IS_CHIP_SELECTED(addr);
  
  if(IS_WRITE_OP(ctrl_pin_val) || chip_select == 0) {
    DATA_DDR = 0x00; //DATA bus to input mode
    DATA_PORT = 0x00; //make sure no pull up on data port
  }
  else {
    DATA_DDR = 0xff; //DATA bus to output mode
    DATA_PORT = emu_read(addr);
  }
  CLK_HALF_DELAY;

  CLK_HIGH;
  CLK_HALF_DELAY;
  if(chip_select == 0) {
    if(IS_WRITE_OP(ctrl_pin_val)) {
      emu_write_monitor(addr, DATA_PIN);
    } else {
      emu_read_monitor(addr, DATA_PIN);
    }
  } else {
    if(IS_WRITE_OP(ctrl_pin_val)) {
      emu_write(addr, DATA_PIN);
    }
  }
  CLK_HALF_DELAY;
}

char print_buff[64];

int main() {
  io_setup();
  uart_init();
  emu_init();

  uart_write("ready for data...\n", 20);
  uint8_t rc = hex_read_mem(emu_rom_ptr(), 0x8000, 0x8400);
  if(rc != 0) {
    sprintf(print_buff, "hex read failed: %d\n", rc);
    uart_write(print_buff, 64);
    while(1) {
    }
  }

  while(1){
    io_update();
    if(uart_read_nb() == 'q') {
      break;
    }
  }

  emu_dump_ram(0x00, 0x10);

  return 0;
}
