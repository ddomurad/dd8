#include "uart.h"

#include <avr/io.h>

#define BAUD 9600
#define UBRR (F_CPU/16/BAUD-1)


void uart_init() {
  UBRR0H = (uint8_t)(UBRR >> 8);
  UBRR0L = (uint8_t)UBRR;

  UCSR0B = (1<<RXEN0)|(1<<TXEN0);
  UCSR0C = (1<<USBS0)|(3<<UCSZ00);
}

void uart_write(char *str, uint8_t max_len) {
  /* Wait for empty transmit buffer */
  //todo: wait for the transmit buffer to be empty??

  /* Put data into buffer, sends the data */
  while(*str != 0x00 && max_len > 0) {
    while (!(UCSR0A & (1<<UDRE0)));
    UDR0 = *str;
    str++;
    max_len--;
  }
}

uint8_t uart_read(char *out_Str, uint8_t max_len) {
  //todo: implement me pls
  return 0x00;
}

