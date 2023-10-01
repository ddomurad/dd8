#ifndef __DEF_H
#define __DEF_H

#define ADDR_L_DDR    DDRA
#define ADDR_H_DDR    DDRC
#define ADDR_L_PORT   PORTA
#define ADDR_H_PORT   PORTC
#define ADDR_L_PIN    PINA
#define ADDR_H_PIN    PINC

#define DATA_DDR      DDRL
#define DATA_PORT     PORTL
#define DATA_PIN      PINL

#define CLK_DDR       DDRG
#define CLK_PORT      PORTG
#define CLK_PIN       PING0 

#define RSTB_DDR       DDRG
#define RSTB_PORT      PORTG
#define RSTB_PIN       PING1 

#define CTRL_DDR      DDRK
#define CTRL_PORT     PORTK
#define CTRL_PINS      PINK
#define CTRL_RW_PIN    PINK0

#define CLK_HALF_DELAY_FUNC _delay_ms(100)

#endif //__DEF_H
