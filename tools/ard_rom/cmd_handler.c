#include "./cmd_handler.h"
#include "emulator.h"
#include "hex_reader.h"
#include "rom.h"
#include "uart.h"
#include <stdint.h>
#include <stdio.h>
#include <string.h>

#define CMD_BUFFER_SIZE 256
#define OUT_BUFFER_SIZE 64

#define HEX_OK 0x00
#define HEX_ERR_UNEXPECTED_TERMINATION 0x01
#define HEX_ERR_ADDR_OVERFLOW 0x02
#define HEX_ERR_UNEXPECTED_RECORD_TYPE 0x03
#define HEX_ERR_BAD_CHECK_SUM 0x04

#define CHECK_HEX_INDEX()                                                      \
  do {                                                                         \
    if (hex_index >= hex_size)                                                 \
      return HEX_ERR_UNEXPECTED_TERMINATION;                                   \
  } while (0)

char cmd_buffer[CMD_BUFFER_SIZE + 1];
uint16_t cmd_buffer_pointer = 0;
uint8_t cmd_buffer_overflow = 0;

char out_buffer[OUT_BUFFER_SIZE];

void write_error(const char *str) {
  uart_write("!err ", 5);
  uart_writeln(str, -1);
}

void dump_rom_buffer() {
  uint8_t *rom = rom_buffer();

  uart_write(":", 1);
  for (uint16_t i = 0; i < ROM_SIZE; i++) {
    sprintf(out_buffer, "%02X ", rom[i]);
    uart_write(out_buffer, OUT_BUFFER_SIZE);
  }

  uart_write("\n", 1);
}

uint8_t h2i(char c) {
  if (c <= '9')
    return c - '0';
  return c - 'A' + 10;
}

uint8_t read_8(char *str, uint16_t *pindex, uint16_t str_len) {
  uint16_t index = *pindex;
  if (index + 1 >= str_len) {
    *pindex = str_len;
    return 0;
  }

  uint8_t d = 0x00;
  d = h2i(str[index]);
  d <<= 4;
  d |= h2i(str[index + 1]);

  *pindex = index + 2;
  return d;
}

uint16_t read_16(char *str, uint16_t *pindex, uint16_t str_len) {
  uint16_t index = *pindex;
  if (index + 3 >= str_len) {
    *pindex = str_len;
    return 0;
  }

  uint16_t d = 0x00;
  d = h2i(str[index + 0]);
  d <<= 4;
  d |= h2i(str[index + 1]);
  d <<= 4;
  d |= h2i(str[index + 2]);
  d <<= 4;
  d |= h2i(str[index + 3]);

  *pindex = index + 4;
  return d;
}

uint8_t hex_write_to_rom_memory(uint8_t *rom, uint16_t addr_start,
                                uint16_t addr_end, char *hex,
                                uint16_t hex_size) {
  uint16_t hex_index = 0;
  uint16_t dsize = 0x00;
  uint16_t addr = 0x00;
  uint8_t cs = 0x00;
  uint8_t rb;

  if (hex_size < 0x0b) {
    return HEX_ERR_UNEXPECTED_TERMINATION;
  }

  if (hex[7] != '0') {
    return HEX_ERR_UNEXPECTED_RECORD_TYPE;
  }
  if (hex[8] == '1') {
    // note: this is an optimistic assumption
    return 0; // probably a :00000001FF
  }
  if (hex[8] != '0') { // note: we either expect a 00 or 01
    return HEX_ERR_UNEXPECTED_RECORD_TYPE;
  }

  hex_index++; // skip the first ':' character

  dsize = read_8(hex, &hex_index, hex_size);
  addr = read_16(hex, &hex_index, hex_size);
  cs = dsize;
  cs += (uint8_t)addr;
  cs += (uint8_t)(addr >> 8);

  if (addr < addr_start) {
    return HEX_ERR_ADDR_OVERFLOW;
  }

  if (addr >= addr_end) {
    return HEX_ERR_ADDR_OVERFLOW;
  }

  addr -= addr_start;

  hex_index += 2;
  CHECK_HEX_INDEX();

  while (dsize > 0) {
    rb = read_8(hex, &hex_index, hex_size);
    CHECK_HEX_INDEX();

    rom[addr++] = rb;
    dsize--;
    cs += rb;
  }

  CHECK_HEX_INDEX();
  cs += read_8(hex, &hex_index, hex_size);
  if (cs != 0) {
    return HEX_ERR_BAD_CHECK_SUM;
  }

  return HEX_OK;
}

uint8_t execute_cmd_buffer() {
  uint8_t r = 0;
  if (!cmd_buffer_pointer) {
    goto unknown_cmd;
  }

  if (*cmd_buffer == ':') {
    emu_stop();
    r = hex_write_to_rom_memory(rom_buffer(), ROM_START_ADDR,
                                ROM_START_ADDR + ROM_SIZE, cmd_buffer,
                                cmd_buffer_pointer);
    switch (r) {
    case HEX_OK:
      uart_writeln("!ok", 3);
      return 0;
    case HEX_ERR_UNEXPECTED_TERMINATION:
      write_error("HEX_ERR_UNEXPECTED_TERMINATION");
      return 1;
    case HEX_ERR_ADDR_OVERFLOW:
      write_error("HEX_ERR_ADDR_OVERFLOW");
      return 1;
    case HEX_ERR_UNEXPECTED_RECORD_TYPE:
      write_error("HEX_ERR_UNEXPECTED_RECORD_TYPE");
      return 1;
    case HEX_ERR_BAD_CHECK_SUM:
      write_error("HEX_ERR_BAD_CHECK_SUM");
      return 1;
    }
  }
  if (*cmd_buffer == 'd') {
    dump_rom_buffer();
    uart_writeln("!ok", 3);
    return 0;
  }
  if (*cmd_buffer == 's') {
    emu_stop();
    uart_writeln("!ok", 3);
    return 0;
  }
  if (*cmd_buffer == 'r') {
    emu_start();
    uart_writeln("!ok", 3);
    return 0;
  }
  if (*cmd_buffer == 'x') {
    emu_reset();
    uart_writeln("!ok", 3);
    return 0;
  }
  if (*cmd_buffer == 'c') {
    uint8_t r = emu_set_clock(cmd_buffer[1] - '0');
    uart_writeln("!ok", 3);
    if (r) {
      write_error("UNSUPPORTED_CLOCK_SETTING");
      return 1;
    }
    return 0;
  }
  if (*cmd_buffer == 'n') {
    emu_step();
    uart_writeln("!ok", 3);
    return 0;
  }
  if (*cmd_buffer == 'b') {
    if (cmd_buffer_pointer != 7) {
      write_error("UNSUPPORTED_BRK_FORMAT");
      return 1;
    }
    uint16_t index = 2;
    uint16_t addr = read_16(cmd_buffer, &index, cmd_buffer_pointer);
    emu_set_brk(addr);
    uart_writeln("!ok", 3);
    return 0;
  }
  if (*cmd_buffer == 'B') {
    emu_clr_brk();
    uart_writeln("!ok", 3);
    return 0;
  }
unknown_cmd:
  write_error("UNKNOWN_CMD");
  return 1;
}

void cmd_init() { cmd_buffer_pointer = 0; } // note: do we need this ?

uint8_t _cmd_read_and_execute() {
  uint8_t d = uart_read_nb();
  if (d == 0x00) {
    return 0;
  }

  if (d == '\n') {
    uint8_t r = 0;
    if (!cmd_buffer_overflow) {
      r = execute_cmd_buffer();
    }

    cmd_buffer[0] = 0;
    cmd_buffer_pointer = 0;
    cmd_buffer_overflow = 0;
    return r;
  }

  if (cmd_buffer_overflow) {
    return 1;
  }

  cmd_buffer[cmd_buffer_pointer] = d;
  cmd_buffer[cmd_buffer_pointer + 1] = 0; // note: do we need this ?
  cmd_buffer_pointer++;

  if (cmd_buffer_pointer >= CMD_BUFFER_SIZE) {
    write_error("CMD_BUFFER_OVERFLOW");
    cmd_buffer_overflow = 1;
    return 1;
  }

  return 0;
}

uint8_t cmd_read_and_execute() {
  uint8_t r = _cmd_read_and_execute();
  if (r) {
    emu_stop();
  }

  return r;
}
