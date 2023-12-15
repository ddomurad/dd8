#include "cmd_handler.h"
#include "def.h"
#include "emulator.h"
#include "uart.h"

#include <stdint.h>
#include <stdio.h>
#include <string.h>

#include <avr/io.h>
#include <avr/interrupt.h>
#include <util/delay.h>
#include <util/delay_basic.h>

int main() {
  emu_init();
  uart_init();
  cmd_init();

  sei();

  uart_writeln("!ready", 6);

  while (1) {
    cmd_read_and_execute();
  }
}
