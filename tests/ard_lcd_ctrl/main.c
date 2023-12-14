#include <string.h>
#include <stdio.h>
#include <avr/io.h>
#include <stdint.h>
#include <util/delay.h>

#define CTRL_DDR  DDRC
#define CTRL_PORT PORTC
#define CTRL_RS   0
#define CTRL_E    1
#define CTRL_RW   2

#define DATA_DDR  DDRA
#define DATA_PORT PORTA

#define TEST_CHAR 0x33 //should be a '0'


// LiquidCrystal lcd(37, 36, 26, 27, 28, 29); // Creates an LCD object. Parameters: (rs, enable, d4, d5, d6, 7)

int main() 
{
  CTRL_DDR  = 0x07; 
  CTRL_PORT = 0x00; 
  DATA_DDR  = 0x00;
  DATA_PORT = 0x00;
  _delay_ms(100);

  DATA_DDR = 0xff;
  DATA_PORT = 0b00001100; //display on
  CTRL_PORT = 0x02; //RS=0, E=1, RW=0
  _delay_ms(10);

  CTRL_PORT = 0x00; //RS=0, E=0, RW=0
  _delay_ms(10);

  DATA_PORT = 0x80; //set DDRAM to 0x00
  CTRL_PORT = 0x02; //RS=0, E=1, RW=0
  _delay_ms(10);

  CTRL_PORT = 0x00; //RS=0, E=0, RW=0
  _delay_ms(10);

  DATA_PORT = TEST_CHAR;
  CTRL_PORT = 0x03; //RS=1, E=1, RW=0
  _delay_ms(10);

  CTRL_PORT = 0x00; //RS=0, E=0, RW=0
  _delay_ms(10);

  while(1){}
}
