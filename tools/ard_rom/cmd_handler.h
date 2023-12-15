#ifndef __CMD_HANDLER
#define __CMD_HANDLER

#include <stdint.h>

void cmd_init();
uint8_t cmd_read_and_execute();

#endif // !__CMD_HANDLER
