#include <string.h>
#include <stdio.h>
#include <avr/io.h>
#include <stdint.h>
#include <util/delay.h>

#define CTRL_DDR  DDRL
#define CTRL_PORT PORTL
#define CTRL_RS   0
#define CTRL_RW   1
#define CTRL_E    2

#define DATA_DDR  DDRA
#define DATA_PORT PORTA

#define TEST_CHAR 0x03 //should be a '0'


int main() 
{
  CTRL_DDR  = 0x07; 
  CTRL_PORT = 0x00; 
  DATA_DDR  = 0x00;
  DATA_PORT = 0x00;

  DATA_DDR = 0xff;
  DATA_PORT = 0b00001110;
  CTRL_PORT = 0x04; //RS=0, RW=0, E=1
  _delay_ms(2);
  DATA_PORT = 0x00;

  DATA_PORT = 0x01;
  CTRL_PORT = 0x04; //RS=0, RW=0, E=1
  _delay_ms(2);
  DATA_PORT = 0x00;
}
