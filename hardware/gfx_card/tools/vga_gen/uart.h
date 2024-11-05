#ifndef __UART_H
#define __UART_H

#include <stdint.h>

void uart_init();

void uart_write(const char *str);
void uart_writeln(const char *str);

uint8_t uart_read();
uint8_t uart_read_nb();

#endif // !__UART_H
