#include <string.h>
#include <stdio.h>
#include <avr/io.h>
#include <stdint.h>
#include <util/delay.h>

#include "rom.h"

#define BAUD 9600
#define UBRR (F_CPU/16/BAUD-1)

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

#define CLK_DELAY _delay_ms(5)
#define IS_WRITE_OP(x) (x & 0x01)

void uart_init()
{
  UBRR0H = (uint8_t)(UBRR >> 8);
  UBRR0L = (uint8_t)UBRR;

  UCSR0B = (1<<RXEN0)|(1<<TXEN0);
  UCSR0C = (1<<USBS0)|(3<<UCSZ00);
}

void uart_send(uint8_t *data)
{
  /* Wait for empty transmit buffer */
  //todo: wait for the transmit buffer to be empty??

  /* Put data into buffer, sends the data */
  while(*data != 0x00){
    while (!(UCSR0A & (1<<UDRE0)));
    UDR0 = *data;

    data++;
  }
}

char state_buffer[15];
void write_state(uint8_t addr_h, uint8_t addr_l, uint8_t data, uint8_t extra, uint8_t clk) 
{
  return;
  // sprintf((char*)state_buffer, "%02X%02X %02X %02x %02x\n", addr_h, addr_l, data, extra & 0x01, clk & 0x01);
  // uart_send((uint8_t*)state_buffer);
}

const char* message = "hello world !";
int main() 
{
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

  memset((void*)state_buffer, 0, 15);
  uart_init();


  CLK_PORT ^= 1<<CLK_PIN;
  _delay_ms(1);
  CLK_PORT ^= 1<<CLK_PIN;
  _delay_ms(1);

  CLK_PORT ^= 1<<CLK_PIN;
  _delay_ms(1);
  CLK_PORT ^= 1<<CLK_PIN;
  _delay_ms(1);

  CLK_PORT ^= 1<<CLK_PIN;
  _delay_ms(1);
  CLK_PORT ^= 1<<CLK_PIN;
  _delay_ms(1);
  uart_send((uint8_t*)"------------\n");
  RSTB_PORT |= 1<<RSTB_PIN;
  _delay_ms(1);

  uint8_t extra_pin = 0;
  while (0x01)
  {
    // CLK -> low
    CLK_PORT &= ~(1<<CLK_PIN);
    CLK_DELAY;

    extra_pin = EXTRA_PIN;
    
    if(IS_WRITE_OP(extra_pin)) {
      DATA_DDR = 0x00;
    } else {
      DATA_DDR = 0xff;
      DATA_PORT = ROM_DATA[AD]
    }

    write_state(
      ADDR_H_PIN,
      ADDR_L_PIN,
      DATA_PIN,
      EXTRA_PIN,
      0
    );

    //CLK -> high
    CLK_PORT |= 1<<CLK_PIN;
    CLK_DELAY;

    write_state(
      ADDR_H_PIN,
      ADDR_L_PIN,
      DATA_PIN,
      EXTRA_PIN,
      1
    );
  }

  return 0;
}
