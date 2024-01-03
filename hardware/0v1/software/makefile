all: clean build

build:
	dd8asm -out ./out/rom.hex -list ./list ./main.asm

clean:
	rm -f ./list/*
	rm -f ./out/*

sim: clean build
	srec_cat ./out/rom.hex   -Intel -o ./out/rom.srec -Motorola
	sim65

monitor:
	socat /dev/ttyUSB0,rawer,b19200,crnl -
