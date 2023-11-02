#ifndef __HEX_DECODE_H
#define __HEX_DECODE_H

#include <stdint.h>

uint8_t hex_read_mem(uint8_t *buffer, uint16_t start_addr, uint16_t end_addr);

#endif // __HEX_DECODE_H
