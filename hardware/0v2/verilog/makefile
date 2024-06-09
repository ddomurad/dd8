all: clean
	# verilator --cc ./dma.v
	# verilator --cc ./dram_rnd.v
	verilator --cc ./sprite_render.v
	# iverilog -Wall -o ./out/vga_sync ./vga_sync.v
	# iverilog -Wall -o ./out/vga_data_mux ./vga_data_mux.v
	# iverilog -Wall -g2001 -o ./out/dma ./dma.v
	# iverilog -Wall -g2001 -o ./out/dma_tb ./dma_tb.v
	iverilog -Wall -g2001 -o ./out/sprite_render ./sprite_render.v
	iverilog -Wall -g2001 -o ./out/sprite_render_tb ./sprite_render_tb.v
	iverilog -Wall -g2001 -o ./out/dram_rnd ./dram_rnd.v
test: all
	vvp ./out/sprite_render_tb
wave: test 
	gtkwave ./out/sprite_render_tb.vcd
clean:
	rm ./out/* -f

