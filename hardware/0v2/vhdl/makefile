STD = --std=08	

all: build run
build:
	ghdl -a $(STD) ./board_ctrl.vhdl
	ghdl -a $(STD) ./board_ctrl_tb.vhdl
	ghdl -e $(STD) board_ctrl
	ghdl -e $(STD) board_ctrl_tb
run:
	ghdl -r $(STD) board_ctrl_tb --vcd=board_ctrl_tb.vcd
wave:
	gtkwave board_ctrl_tb.vcd 
