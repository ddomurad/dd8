#ifndef __DEF_H
#define __DEF_H

#define ADDR_L_DDR      DDRA
#define ADDR_H_DDR      DDRC
#define ADDR_L_PORT     PORTA
#define ADDR_H_PORT     PORTC
#define ADDR_L_PIN      PINA
#define ADDR_H_PIN      PINC

#define DATA_DDR        DDRL
#define DATA_PORT       PORTL
#define DATA_PIN        PINL

#define CLK_DDR         DDRG
#define CLK_PORT        PORTG
#define CLK_PIN         PING0 

#define RSTB_DDR        DDRG
#define RSTB_PORT       PORTG
#define RSTB_PIN        PING1 

#define CTRL_DDR        DDRK
#define CTRL_PORT       PORTK
#define CTRL_PINS       PINK
#define CTRL_OE_PIN     PINK0
#define CTRL_WE_PIN     PINK1
#define CTRL_CE_ROM_PIN PINK2
#define CTRL_C1 PINK3
#define CTRL_C2 PINK4
#define CTRL_C3 PINK5

#endif //__DEF_H
