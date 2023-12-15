#include "./emulator.h"
#include "./def.h"
#include "rom.h"
#include "uart.h"

#include <stdint.h>
#include <stdio.h>

#include <avr/interrupt.h>
#include <avr/io.h>
#include <util/delay.h>

#define ENABLE_TMR_INT()                                                       \
  do {                                                                         \
    TIMSK1 = (1 << TOIE1);                                                     \
  } while (0)
#define DISABLE_TMR_INT()                                                      \
  do {                                                                         \
    TIMSK1 = 0x00;                                                             \
  } while (0)

uint8_t clk_state = 0; // 0 - low transition, 1 - low sustain,
                       // 2 - high transition, 3 - high sustain

volatile uint8_t state = EMU_STATE_STOPED;
volatile uint8_t req_state = EMU_STATE_STOPED;
volatile uint8_t monitor_disabled = 0;
volatile uint8_t single_step = 0;
volatile uint16_t brk_addr = 0;
volatile uint8_t brk_set = 0;

#define SET_CLK_LOW() (CLK_PORT &= ~(1 << CLK_PIN))
#define SET_CLK_HIGH() (CLK_PORT |= (1 << CLK_PIN))

#define IS_CE_SET(x) (x & (1 << CTRL_CE_ROM_PIN))
#define IS_OE_SET(x) (x & (1 << CTRL_OE_PIN))

char monitor_buffer[16];
void write_monitor(uint8_t ctrl, uint16_t addr, uint8_t data) {
  if (ctrl & (1 << CTRL_CE_ROM_PIN)) {
    uart_write("   ", 3);
  } else {
    uart_write("ro ", 3);
  }
  if (ctrl & (1 << CTRL_CE_RAM_PIN)) {
    uart_write("   ", 3);
  } else {
    uart_write("ra ", 3);
  }
  if (ctrl & (1 << CTRL_CE_KB_PIN)) {
    uart_write("   ", 3);
  } else {
    uart_write("kb ", 3);
  }
  if (ctrl & (1 << CTRL_CE_LCD_PIN)) {
    uart_write("ld ", 3);
  } else {
    uart_write("   ", 3);
  }

  if (ctrl & (1 << CTRL_OE_PIN)) {
    uart_write("   ", 3);
  } else {
    uart_write("oe ", 3);
  }
  if (ctrl & (1 << CTRL_WE_PIN)) {
    uart_write("   ", 3);
  } else {
    uart_write("we ", 3);
  }

  sprintf(monitor_buffer, "%04X %02X", addr, data);
  uart_writeln(monitor_buffer, 16);
}

ISR(TIMER1_OVF_vect) {
  uint16_t addr = 0x0000;
  uint8_t ctrl_port = 0x00;

  if (clk_state == 0) {
    SET_CLK_LOW();
    DATA_DDR = 0x00;  // data port to input mode
    DATA_PORT = 0x00; // make sure not pull up on data port
    if (req_state == EMU_STATE_STOPED) {
      DISABLE_TMR_INT();
      state = EMU_STATE_STOPED;
      return;
    } else {
      if (state == EMU_STATE_RUNNING && single_step) {
        DISABLE_TMR_INT();
        state = EMU_STATE_STOPED;
        single_step = 0;
        return;
      }
      state = EMU_STATE_RUNNING;
    }
    clk_state = 1;
    return;
  } else if (clk_state == 1) {
    clk_state = 2;
    return;
  } else if (clk_state == 2) {
    SET_CLK_HIGH();
    addr = ((ADDR_H_PIN << 8) | ADDR_L_PIN);
    ctrl_port = CTRL_PINS;
    if (!IS_CE_SET(ctrl_port) && !IS_OE_SET(ctrl_port)) {
      DATA_DDR = 0xff; // data port to output mode
      DATA_PORT = rom_read(addr);
    }
    clk_state = 3;
    return;
  } else if (clk_state == 3) {
    addr = ((ADDR_H_PIN << 8) | ADDR_L_PIN);
    if (!monitor_disabled || single_step) {
      write_monitor(CTRL_PINS, addr, DATA_PIN);
    }
    if (brk_set && addr == brk_addr) {
      if(monitor_disabled) {
        write_monitor(CTRL_PINS, addr, DATA_PIN);
      }
      req_state = EMU_STATE_STOPED;
    }
    clk_state = 0;
    return;
  }

  clk_state = 0;
}

void reset_procedure() {
  RSTB_PORT &= ~(1 << RSTB_PIN); // pull RESET LOW
  for (int8_t i = 0; i < 4; i++) {
    SET_CLK_HIGH();
    _delay_ms(1);
    SET_CLK_LOW();
    _delay_ms(1);
  }
  RSTB_PORT |= (1 << RSTB_PIN); // pull RESET HIGH
  _delay_ms(1);
}

void emu_init() { // todo: do we need this  ?
  // IO setup
  // setup clk pin as output
  CLK_DDR |= 1 << CLK_PIN;
  CLK_PORT &= ~(1 << CLK_PIN); // make sure it is set to low at start

  // setup reset pin as output
  RSTB_DDR |= 1 << RSTB_PIN;
  RSTB_PORT &= ~(1 << RSTB_PIN); // make sure it is set to low at start

  // setup addr ports as input
  ADDR_H_DDR = 0x00;
  ADDR_L_DDR = 0x00;
  ADDR_H_PORT = 0x00;
  ADDR_L_PORT = 0x00;

  // setup data port as input
  DATA_DDR = 0x00;
  DATA_PORT = 0x00;

  // setup CTRL port as input
  CTRL_DDR = 0x00;
  CTRL_PORT = 0xff; // internal pull up for all CTRL lines

  // TIMER SETUP
  DISABLE_TMR_INT();
  // FAST PWM, TOP OCR1A
  TCCR1A = (1 << WGM10) | (1 << WGM11);
  // FAST PWM, TOP OCR1A, 1024 clock prescaler
  TCCR1B = (0 << WGM12) | (0 << WGM13) | (1 << CS10) | (1 << CS12);
  // reset timer counter
  TCNT1L = 0x00;
  TCNT1H = 0x00;

  reset_procedure();
}

void emu_start() {
  req_state = EMU_STATE_RUNNING;
  ENABLE_TMR_INT();
  while (state != EMU_STATE_RUNNING)
    ;
}

void emu_stop() {
  req_state = EMU_STATE_STOPED;
  while (state != EMU_STATE_STOPED)
    ;
}

void emu_reset() {
  emu_stop();
  reset_procedure();
  emu_start();
}

uint8_t emu_set_clok(uint8_t v) {
  cli();
  if (v == 0) {
    TCCR1A = (1 << WGM10) | (1 << WGM11);
    TCCR1B = (0 << WGM12) | (0 << WGM13) | (1 << CS10) | (1 << CS12);
  } else if (v == 1) {
    TCCR1A = (1 << WGM10) | (0 << WGM11);
    TCCR1B = (0 << WGM12) | (0 << WGM13) | (1 << CS10) | (1 << CS12);
  } else if (v == 2) {
    TCCR1A = (0 << WGM10) | (0 << WGM11);
    TCCR1B = (0 << WGM12) | (0 << WGM13) | (1 << CS10);
  } else if (v == 3) {
    TCCR1A = (0 << WGM10) | (1 << WGM11);
    TCCR1B = (0 << WGM12) | (0 << WGM13) | (1 << CS10);
  } else if (v == 4) {
    TCCR1A = (1 << WGM10) | (1 << WGM11);
    TCCR1B = (1 << WGM12) | (1 << WGM13) | (1 << CS10);
    OCR1A = 0x10;
  }
  monitor_disabled = v >= 3;

  sei();
  return v < 0 || v > 4; // error if v not int range [0, 4]
}

void emu_step() {
  emu_stop();
  single_step = 1;
  req_state = EMU_STATE_RUNNING;
  ENABLE_TMR_INT();
  while (single_step)
    ;
}

void emu_set_brk(uint16_t addr) {
  brk_addr = addr;
  brk_set = 1;
}
void emu_clr_brk() { brk_set = 0; }
