#include "def.h"
#include "emulator.h"
#include "uart.h"

#include <avr/io.h>
#include <stdint.h>
#include <util/delay.h>

#define ADDR_L_DDR    DDRA
#define ADDR_H_DDR    DDRC
#define DATA_DDR      DDRL
#define CTRL_DDR      DDRK

#define ADDR_L_PORT   PORTA
#define ADDR_H_PORT   PORTC
#define DATA_PORT     PORTL
#define CTRL_PORT     PORTK

#define ADDR_L_PIN    PINA
#define ADDR_H_PIN    PINC
#define DATA_PIN      PINL
#define CTR_PINS      PINK

#define CLK_DDR       DDRG
#define CLK_PORT      PORTG
#define CLK_PIN       PING0 

#define RSTB_DDR       DDRG
#define RSTB_PORT      PORTG
#define RSTB_PIN       PING1 

#define CTRL_RW_PIN    PINK0

#define CLK_LOW (CLK_PORT &= ~(1<<CLK_PORT))
#define CLK_HIGH (CLK_PORT |= (1<<CLK_PORT))

#define CLK_HALF_DELAY CLK_HALF_DELAY_FUNC

#define ADDR_FULL ((ADDR_H_PIN << 8) | ADDR_L_PIN)
#define IS_WRITE_OP(x) (!(x & (1 << CTRL_RW_PIN)))

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
    CLK_HALF_DELAY;
    CLK_HALF_DELAY;
    CLK_LOW;
    CLK_HALF_DELAY;
    CLK_HALF_DELAY;
  }
}

inline static void io_update() {
  uint8_t ctrl_pin_val;

  CLK_LOW;
  CLK_HALF_DELAY;
  ctrl_pin_val = CTR_PINS;

  if(IS_WRITE_OP(ctrl_pin_val)) {
    DATA_DDR = 0x00; //DATA bus to input mode
    DATA_PORT = 0x00; //make sure no pull up on data port
  }
  else {
    DATA_DDR = 0xff; //DATA bus to output mode
    DATA_PORT = emu_read(ADDR_FULL);
  }
  CLK_HALF_DELAY;

  CLK_HIGH;
  CLK_HALF_DELAY;
  if(IS_WRITE_OP(ctrl_pin_val)) {
    emu_write(ADDR_FULL, DATA_PIN);
  }
  CLK_HALF_DELAY;
}

int main() {
  io_setup();
  uart_init();
  emu_init();

  while(1){
    io_update();
  }

  return 0;
}
