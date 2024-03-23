all: clean build 

clean:
	rm ./out/* -rf 

build:
	dd8asm -out ./out/bios.hex -list ./list/ ./main.asm
	srec_cat ./out/bios.hex -Intel -o ./out/bios.srec -Motorola

monitor:
	socat /dev/ttyUSB0,rawer,b19200,crnl -
