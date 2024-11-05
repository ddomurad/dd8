all: clean build test
build:
	iverilog -Wall -o ./out/vga_gen ./vga_gen.v
	iverilog -Wall -o ./out/vga_gen_tb ./vga_gen_tb.v
	iverilog -Wall -o ./out/vga_dma ./vga_dma.v
	iverilog -Wall -o ./out/vga_dma_tb ./vga_dma_tb.v
test: all
	# vvp ./out/vga_gen_tb
	vvp ./out/vga_dma_tb
wave: test 
	# gtkwave ./out/vga_gen_tb.vcd
	gtkwave ./out/vga_dma_tb.vcd
clean:
	rm ./out/* -f
