#ifndef __UART_H
#define __UART_H

#include <stdint.h>


void uart_init();
void uart_write(char *str, uint8_t max_len);
uint8_t uart_read();
uint8_t uart_read_nb();

// void uart_int_mode();
// char* uart_get_line();
#endif // !__UART_H

