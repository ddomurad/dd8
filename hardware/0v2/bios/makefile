all: clean build 

clean:
	rm ./out/* -rf 

build:
	dd8asm -out ./out/bios.hex -list ./list/ ./main.asm

monitor:
	socat /dev/ttyUSB0,rawer,b19200,crnl -
