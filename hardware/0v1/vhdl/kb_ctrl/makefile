STD = --std=08	

all: build run

build:
	ghdl -a $(STD) ./kb_addr_decoder.vhdl
	ghdl -a $(STD) ./kb_addr_decoder_tb.vhdl
	ghdl -a $(STD) ./kb_ctrl.vhdl
	ghdl -a $(STD) ./kb_ctrl_tb.vhdl
	ghdl -e $(STD) kb_addr_decoder
	ghdl -e $(STD) kb_addr_decoder_tb
	ghdl -e $(STD) kb_ctrl 
	ghdl -e $(STD) kb_ctrl_tb 
run:
	ghdl -r $(STD) kb_ctrl_tb --vcd=kb_ctrl.vcd
	ghdl -r $(STD) kb_addr_decoder_tb --vcd=kb_addr_decoder_tb.vcd
wave:
	gtkwave kb_ctrl.vcd 
