SRC_DIR := ./
OUT_DIR := ./out
SRC_FILES := $(wildcard $(SRC_DIR)/*.c)
OBJ_FILES := $(patsubst $(SRC_DIR)/%.c,$(OUT_DIR)/%.o,$(SRC_FILES))
OUT_PRG	:= $(OUT_DIR)/ard_monitor.elf
OUT_HEX	:= $(OUT_DIR)/ard_monitor.hex

MMCU := atmega1280

DEFS := -DF_CPU=16000000
CFLAGS := -Wall -O3 -mmcu=$(MMCU)

PROGRAMMER := arduino
PROGRAMMER_PORT := /dev/ttyUSB0
PROGRAMMER_RATE := 57600

GCC=avr-gcc

all: clean $(OBJ_FILES)
	$(GCC) -mmcu=$(MMCU) $(OBJ_FILES) -o $(OUT_PRG)
	objcopy -O ihex $(OUT_PRG) $(OUT_HEX)
	avr-size -C --mcu=$(MMCU) $(OUT_PRG)

clean:
	rm ./out -r -f 
	mkdir ./out

$(OUT_DIR)/%.o: $(SRC_DIR)/%.c 
	$(GCC) $(CFLAGS) -c $(DEFS) $< -o $@ 
	
flash: all
	avrdude -p $(MMCU) -c $(PROGRAMMER) -P $(PROGRAMMER_PORT) -b $(PROGRAMMER_RATE) -D -U flash:w:$(OUT_HEX)

monitor:
	socat /dev/ttyUSB0,rawer,b19200,crnl -

build_gfx_data:
	(cd ../../scripts/ && python3 ./img_conv.py)
	xxd -i ../../scripts/gfx_dada.bin | sed 's/______scripts_gfx_dada_bin/gfx_data/g' | sed 's/= {/PROGMEM = {/g' | sed 's/unsigned/const unsigned/g' > gfx_data.h
