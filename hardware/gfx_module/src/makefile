all: clean build test
build:
	iverilog -Wall -o ./out/gfx_vga ./gfx_vga.v
	# iverilog -Wall -o ./out/gfx_vga_tb ./gfx_vga_tb.v
	iverilog -Wall -o ./out/gfx_dma ./gfx_dma.v
	# iverilog -Wall -o ./out/gfx_dma_tb ./gfx_dma_tb.v
	iverilog -Wall -o ./out/gfx_tb ./gfx_tb.v
test: all
	# vvp ./out/gfx_vga_tb
	# vvp ./out/gfx_dma_tb
	vvp ./out/gfx_tb
wave_vga: test
	gtkwave ./out/gfx_vga_tb.vcd
wave_dma: test
	gtkwave ./out/gfx_dma_tb.vcd
wave_gfx: test
	gtkwave ./out/gfx_tb.vcd
clean:
	rm ./out/* -f
