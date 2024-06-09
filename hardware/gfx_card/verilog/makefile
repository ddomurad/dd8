all: clean
	verilator --cc ./vga_sync.v
	verilator --cc ./vga_dump.v
	iverilog -Wall -o ./out/vga_sync ./vga_sync.v
	iverilog -Wall -o ./out/vga_dump ./vga_dump.v
clean:
	rm ./out/* -f

