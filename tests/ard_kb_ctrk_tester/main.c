#include <string.h>
#include <stdio.h>
#include <avr/io.h>
#include <stdint.h>
#include <util/delay.h>

#define BAUD 9600
#define UBRR (F_CPU/16/BAUD-1)

#define DATA_DDR  DDRL
#define DATA_PORT PORTL
#define DATA_PIN  PINL

#define CTRL_DDR  DDRC 
#define CTRL_PORT PORTC 
#define CTRL_PIN  PINC 
#define CTRL_NCE  0 
#define CTRL_NOE  1 
#define CTRL_NWE  2 

#define ADDR_DDR  DDRA
#define ADDR_PORT PORTA
#define ADDR_PIN  PINA

void uart_init()
{
  UBRR0H = (uint8_t)(UBRR >> 8);
  UBRR0L = (uint8_t)UBRR;

  UCSR0B = (1<<RXEN0)|(1<<TXEN0);
  UCSR0C = (1<<USBS0)|(3<<UCSZ00);
}

void uart_send(char *data)
{
  /* Put data into buffer, sends the data */
  while(*data != 0x00){
    while (!(UCSR0A & (1<<UDRE0)));
    UDR0 = *data;
    data++;
  }
}

char state_buffer[64];

void setup() 
{
  DATA_DDR = 0x00; 
  DATA_PORT = 0x00; 
  CTRL_DDR = (1<<CTRL_NCE) | (1<<CTRL_NOE) | (1<<CTRL_NWE);
  CTRL_PORT = (1<<CTRL_NCE) | (1<<CTRL_NOE) | (1<<CTRL_NWE);
  ADDR_DDR = 0x03;
  ADDR_PORT = 0x00;
}

void test_tri_state_data() 
{
  uart_send("tri state data test\n");
  DATA_PORT = 0xff;
  _delay_ms(1);
  uint8_t dv = DATA_PIN; 
  sprintf(state_buffer, "data: %02X (FF)\n", dv);
  uart_send(state_buffer);

  CTRL_PORT &= ~(1<<CTRL_NCE);
  CTRL_PORT &= ~(1<<CTRL_NOE);
  _delay_ms(1);
  dv = DATA_PIN; 
  CTRL_PORT = (1<<CTRL_NCE) | (1<<CTRL_NOE) | (1<<CTRL_NWE);

  sprintf(state_buffer, "data: %02X (00)\n", dv);
  uart_send(state_buffer);
}

void test_constant_pull() 
{
  DATA_PORT = 0xff;
  CTRL_PORT &= ~(1<<CTRL_NCE);
  CTRL_PORT &= ~(1<<CTRL_NOE);
  
  while(1)
  {
    for(int i=0;i<4;i++) 
    {
      _delay_ms(10);
      ADDR_PORT = i; 
      uint8_t dv = DATA_PIN;
      if(dv != 0x00)
      {
        sprintf(state_buffer, "%02X%02X\n", i, dv);
        uart_send(state_buffer);
      }
    }
  }
}

int main() 
{
  setup();
  uart_init();
  uart_send("initialized ...\n");
  // _delay_ms(10);
  // test_tri_state_data();
  // _delay_ms(10);
  test_constant_pull();

  while(1){}
}
