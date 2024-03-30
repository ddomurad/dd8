all: clean cpy build 

clean:
	rm ./out/* -rf 
	rm ./bios.asm -f

cpy:
	cp ../../bios_v0/exp/bios.asm . -f

build:
	dd8asm -out ./out/bios.hex -list ./list/ ./main.asm
	srec_cat ./out/bios.hex -Intel -o ./out/bios.srec -Motorola

monitor:
	socat /dev/ttyUSB0,rawer,b19200,crnl -
