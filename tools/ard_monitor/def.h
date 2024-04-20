#ifndef __DEF_H
#define __DEF_H

#define ADDR_L_DDR  DDRC
#define ADDR_L_PORT PORTC
#define ADDR_L_PINS PINC

#define ADDR_H_DDR  DDRA
#define ADDR_H_PORT PORTA
#define ADDR_H_PINS PINA

#define DATA_DDR  DDRL
#define DATA_PORT PORTL
#define DATA_PINS PINL

#define CTRL_DDR  DDRK
#define CTRL_PORT PORTK
#define CTRL_PINS PINK

#define CLK_DDR  DDRB
#define CLK_PORT PORTB
#define CLK_PINS PINB
#define CLK_PIN  PINB0



/*
 * Memory Lock (MLB)  1
 * Read/Write RWB     2
 * Reset RESB         3
 * Synchorize with opcode SYNC 4
 * Vector Pull VSS    5
 * */
#endif //__DEF_H
