#include <avr/io.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <util/delay.h>

#define DATA_DDR DDRA
#define DATA_PORT PORTA
#define DATA_PIN PINA

#define RS_DDR DDRC
#define RS_PORT PORTC

#define CTRL_DDR DDRL
#define CTRL_PORT PORTL
#define CTRL_RST_PIN PINL0
#define CTRL_RWB_PIN PINL1
#define CTRL_PHI2_PIN PINL2
#define CTRL_CS0_PIN PINL3
#define CTRL_CS1B_PIN PINL4

#define WATCH_DDR DDRK
#define WATCH_PORT PORTK
#define WATCH_IRQB_PIN PINK0
#define WATCH_PIN PINK

#define BAUD 9600
#define UBRR (F_CPU / 16 / BAUD - 1)

// 9600 baud,
// ctrl bits 1110

#define CLK_HIGH CTRL_PORT |= (1 << CTRL_PHI2_PIN)
#define CLK_LOW CTRL_PORT &= ~(1 << CTRL_PHI2_PIN)
#define CLK_DELAY _delay_us(1)

#define FULL_CLK                                                               \
  do {                                                                         \
    CLK_HIGH;                                                                  \
    CLK_DELAY;                                                                 \
    CLK_LOW;                                                                   \
    CLK_DELAY;                                                                 \
  } while (0)

#define CS_ON                                                                  \
  do {                                                                         \
    uint8_t tmp = CTRL_PORT;                                                   \
    tmp |= (1 << CTRL_CS0_PIN);                                                \
    tmp &= ~(1 << CTRL_CS1B_PIN);                                              \
    CTRL_PORT = tmp;                                                           \
  } while (0)
#define CS_OFF                                                                 \
  do {                                                                         \
    uint8_t tmp = CTRL_PORT;                                                   \
    tmp &= ~(1 << CTRL_CS0_PIN);                                               \
    tmp |= (1 << CTRL_CS1B_PIN);                                               \
    CTRL_PORT = tmp;                                                           \
  } while (0)

#define RW_READ CTRL_PORT |= (1 << CTRL_RWB_PIN)
#define RW_WRITE CTRL_PORT &= ~(1 << CTRL_RWB_PIN)

#define RESET_HIGH CTRL_PORT |= (1 << CTRL_RST_PIN)
#define RESET_LOW CTRL_PORT &= ~(1 << CTRL_RST_PIN)

void uart_init() {
  UBRR0H = (uint8_t)(UBRR >> 8);
  UBRR0L = (uint8_t)UBRR;

  UCSR0B = (1 << RXEN0) | (1 << TXEN0);
  UCSR0C = (1 << USBS0) | (3 << UCSZ00);
}

void uart_send(char *data) {
  /* Wait for empty transmit buffer */
  // todo: wait for the transmit buffer to be empty??

  /* Put data into buffer, sends the data */
  while (*data != 0x00) {
    while (!(UCSR0A & (1 << UDRE0)))
      ;
    UDR0 = *data;

    data++;
  }
}

char state_buffer[64];
void write_state(const char *state, uint8_t val) {
  sprintf((char *)state_buffer, "%s %02X\n", state, val);
  uart_send((char *)state_buffer);
}

void init() {
  DATA_DDR = 0x00;
  DATA_PORT = 0x00;
  RS_DDR = 0x00;
  RS_PORT = 0x00;
  CTRL_DDR = 0x00;
  CTRL_PORT = 0x00;
  WATCH_DDR = 0x00;
  WATCH_PORT = 0x00;

  CTRL_DDR = 0xff;
  RS_DDR = 0x03;
  CS_OFF;
  RW_READ;
}

void reset() {
  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CLK_DELAY;

  RESET_HIGH;

  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CLK_DELAY;

  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CLK_DELAY;

  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CLK_DELAY;

  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
}

uint8_t read_status_reg() {
  // DATA_DDR = 0x00;
  // DATA_PORT = 0x00;
  RS_PORT = 0x01;
  // RW_READ;
  // CLK_DELAY;
  CS_ON;
  CLK_HIGH;
  CLK_DELAY;
  uint8_t status_reg = DATA_PIN;
  CLK_LOW;
  CS_OFF;
  CLK_DELAY;

  return status_reg;
}

void write_ctrl_reg() {
  RS_PORT = 0x03;
  RW_WRITE;

  DATA_PORT = 0b00011110; // 1 stop bit, 8 bits, 9600 baud
  DATA_DDR = 0xff;

  CS_ON;
  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CS_OFF;

  DATA_DDR = 0x00;
  DATA_PORT = 0x00;
  RW_READ;

  CLK_DELAY;
}

void write_reg(uint8_t addr, uint8_t data) {
  RS_PORT = addr;
  RW_WRITE;

  DATA_PORT = data; // 1 stop bit, 8 bits, 9600 baud
  DATA_DDR = 0xff;

  CS_ON;
  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CS_OFF;

  DATA_DDR = 0x00;
  DATA_PORT = 0x00;
  RW_READ;

  CLK_DELAY;
}

void write_cmd_reg() {
  RS_PORT = 0x02;
  RW_WRITE;

  DATA_PORT = 0b00001001; // DTR = 0;
  DATA_DDR = 0xff;

  CS_ON;
  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CS_OFF;

  DATA_DDR = 0x00;
  DATA_PORT = 0x00;
  RW_READ;

  CLK_DELAY;
}

void write_transfer(char c) {
  RS_PORT = 0x00;
  RW_WRITE;

  DATA_PORT = c;
  DATA_DDR = 0xff;

  CS_ON;
  CLK_HIGH;
  CLK_DELAY;
  CLK_LOW;
  CS_OFF;

  DATA_DDR = 0x00;
  DATA_PORT = 0x00;
  RW_READ;

  CLK_DELAY;
}

uint8_t read_register(uint8_t addr) {
  RS_PORT = addr;
  CS_ON;
  CLK_HIGH;
  CLK_DELAY;
  uint8_t status_reg = DATA_PIN;
  CLK_LOW;
  CS_OFF;
  CLK_DELAY;

  return status_reg;
}

int main() {
  init();
  uart_init();
  reset();

  write_state("init", 0x00);
  write_reg(0x01, 0x00);
  uint8_t r = read_status_reg();
  write_state("st", r);

  write_state("wch", WATCH_PIN & 0x01);

  write_ctrl_reg();
  FULL_CLK;
  write_cmd_reg();
  FULL_CLK;

  FULL_CLK;
  FULL_CLK;
  write_transfer('4');
  FULL_CLK;

  uint8_t i = 0;
  for (i = 0; i < 10; i++) {
    write_state("wch", WATCH_PIN & 0x01);

    // r = read_status_reg();
    // write_state("st", r);

    FULL_CLK;
    FULL_CLK;
    FULL_CLK;
    FULL_CLK;
    FULL_CLK;
    FULL_CLK;
  }

  uint8_t *buffer = (uint8_t*)"daniel";

  while (1) {
    FULL_CLK;
    if((WATCH_PIN & 0x01) == 0x00) {
      r = read_status_reg();
      write_state("st",r);

      write_transfer(*buffer);
      buffer++;
      if(*buffer == 0x00) {
        break;
      }
    }
  }

  while (1) {
  }
}
