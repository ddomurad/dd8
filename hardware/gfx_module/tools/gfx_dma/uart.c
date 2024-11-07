#include "uart.h"

#include <avr/io.h>
#include <avr/pgmspace.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>

#define BAUD 19200
#define UBRR (F_CPU / 16 / BAUD - 1)

void uart_init() {
  UBRR0H = (uint8_t)(UBRR >> 8);
  UBRR0L = (uint8_t)UBRR;

  UCSR0B = (1 << RXEN0) | (1 << TXEN0);
  UCSR0C = (0 << USBS0) | (3 << UCSZ00);
}

void uart_write(const char *str) {
  /* Put data into buffer, sends the data */
  while (*str != 0x00) {
    while (!(UCSR0A & (1 << UDRE0)))
      ;
    UDR0 = *str;
    str++;
  }
}

void uart_writeln(const char *str) {
  uart_write(str);
  uart_write("\n");
}

uint8_t uart_read() {
  while ((UCSR0A & (1 << RXC0)) == 0)
    ;
  return UDR0;
}

uint8_t uart_read_nb() {
  if ((UCSR0A & (1 << RXC0)) == 0) {
    return 0x00;
  }
  return UDR0;
}
