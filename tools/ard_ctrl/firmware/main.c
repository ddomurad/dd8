#include "def.h"
#include "io.h"
#include "serial.h"
#include "emulator.h"

int main() {
  io_setup();

  while(1){
    io_update();
  }

  return 0;
}
