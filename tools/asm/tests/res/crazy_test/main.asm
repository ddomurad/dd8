.inc "defs.asm"
.inc "test_lib.asm"

.org 0x1000
ldai 0x72
sta REG1
ldai 0x1c 
sta REG2
jsr simple_add
