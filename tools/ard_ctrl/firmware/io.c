#include "io.h"

#include <avr/io.h>
#include <util/delay.h>

#define ADDR_L_DDR    DDRA
#define ADDR_H_DDR    DDRC
#define DATA_DDR      DDRL
#define EXTRA_DDR     DDRK

#define ADDR_L_PORT   PORTA
#define ADDR_H_PORT   PORTC
#define DATA_PORT     PORTL
#define EXTRA_PORT    PORTK

#define ADDR_L_PIN    PINA
#define ADDR_H_PIN    PINC
#define DATA_PIN      PINL
#define EXTRA_PIN     PINK

#define CLK_DDR       DDRG
#define CLK_PORT      PORTG
#define CLK_PIN       PING0 

#define RSTB_DDR       DDRG
#define RSTB_PORT      PORTG
#define RSTB_PIN       PING1 

#define CLK_LOW (CLK_PORT &= ~(1<<CLK_PORT))
#define CLK_HIGH (CLK_PORT |= (1<<CLK_PORT))


void io_setup() {
  CLK_DDR |= 1<<CLK_PIN;
  CLK_PORT &= ~(1<<CLK_PIN);

  RSTB_DDR |= 1<<RSTB_PIN;
  RSTB_PORT &= ~(1<<RSTB_PIN);

  ADDR_H_DDR = 0x00; 
  ADDR_L_DDR = 0x00; 
  DATA_DDR   = 0x00; 
  EXTRA_DDR  = 0x00; 

  ADDR_H_PORT = 0x00; 
  ADDR_L_PORT = 0x00; 
  DATA_PORT   = 0x00; 
  EXTRA_PORT  = 0x00; 


  CLK_HIGH;
  _delay_ms(1);
  CLK_LOW;
  _delay_ms(1);

  CLK_HIGH;
  _delay_ms(1);
  CLK_LOW;
  _delay_ms(1);

  CLK_HIGH;
  _delay_ms(1);
  CLK_LOW;
  _delay_ms(1);

  CLK_HIGH;
  _delay_ms(1);
}


void io_update() {
  CLK_HIGH;
  _delay_ms(1);

  CLK_LOW;
  _delay_ms(1);
}
