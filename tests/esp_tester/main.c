#include <avr/io.h>
#include <stdint.h>
#include <stdio.h>
#include <string.h>
#include <util/delay.h>

#define BAUD 9600
#define UBRR (F_CPU / 16 / BAUD - 1)

#define O_DDR DDRK
#define O_PORT PORTK

#define I_DDR DDRF
#define I_PORT PORTF
#define I_PIN PINF

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

char state_buffer[15];
void write_state(uint8_t addr_h, uint8_t addr_l, uint8_t data, uint8_t extra,
                 uint8_t clk) {
  return;
  // sprintf((char*)state_buffer, "%02X%02X %02X %02x %02x\n", addr_h, addr_l,
  // data, extra & 0x01, clk & 0x01); uart_send((uint8_t*)state_buffer);
}

#define TEST_CASES 32
uint8_t TEST_SIGNAL[TEST_CASES] = {
    // a15, a14, a13, clk, rw
    // clk low, write
    0b00000,
    0b00100,
    0b01000,
    0b01100,
    0b10000,
    0b10100,
    0b11000,
    0b11100,

    // clk high, write
    0b00010,
    0b00110,
    0b01010,
    0b01110,
    0b10010,
    0b10110,
    0b11010,
    0b11110,

    // clk low, read
    0b00001,
    0b00101,
    0b01001,
    0b01101,
    0b10001,
    0b10101,
    0b11001,
    0b11101,

    // clk high, read
    0b00011,
    0b00111,
    0b01011,
    0b01111,
    0b10011,
    0b10111,
    0b11011,
    0b11111,
};

// #define INPUT_MASK 0x3f

uint8_t TEST_EXP[TEST_CASES] = {
    // ce2_uart_b, ce1_uart, ce_rom_b, ce_ram_b, oe_b, we_b
    // low clk, write
    0b101011,
    0b101011,
    0b101011,
    0b101011,
    0b101111,
    0b101111,
    0b011111,
    0b100111,

    // high clk, write
    0b101010,
    0b101010,
    0b101010,
    0b101010,
    0b101110,
    0b101110,
    0b011110,
    0b100110,

    // low clk, read
    0b101011,
    0b101011,
    0b101011,
    0b101011,
    0b101111,
    0b101111,
    0b011111,
    0b100111,

    // low high, read
    0b101001,
    0b101001,
    0b101001,
    0b101001,
    0b101101,
    0b101101,
    0b011101,
    0b100101,
};

char msg_buffer[64];
void write_res(uint8_t sig, uint8_t exp, uint8_t act, uint8_t pass) {
  if (pass) {
    sprintf(msg_buffer, " ok %02X %02X %02X\n", sig, exp, act);
  } else {
    sprintf(msg_buffer, "ERR %02X %02X %02X\n", sig, exp, act);
  }
  uart_send(msg_buffer);
}

void test() {
  for (uint8_t i = 0; i < TEST_CASES; i++) {
    O_PORT = TEST_SIGNAL[i];
    _delay_ms(1);
    uint8_t actuall = I_PIN;
    uint8_t pass = actuall == TEST_EXP[i];

    write_res(TEST_SIGNAL[i], TEST_EXP[i], actuall, pass);
  }
}

int main() {
  I_DDR = 0x00;
  O_DDR = 0x00;
  I_PORT = 0x00;
  O_PORT = 0x00;

  O_DDR = 0xFF;

  uart_init();
  uart_send("ready...\n");

  while (1) {
    test();
    _delay_ms(500);
  }
}
