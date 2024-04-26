STD = --std=08	

all: build run
build:
	ghdl -a $(STD) ./board_ctrl.vhdl
	# ghdl -a $(STD) ./board_ctrl_tb.vhdl
	ghdl -e $(STD) board_ctrl
	# ghdl -e $(STD) board_ctrl_tb
	ghdl --vpi-compile gcc -c entry.c -o vpi.o
	ghdl --vpi-link gcc vpi.o -o vpi.vpi
run:
	# ghdl -r $(STD) board_ctrl_tb --vcd=board_ctrl_tb.vcd
	ghdl  -r $(STD) board_ctrl --vpi=./vpi.vpi
	# ghdl -r tb --vpi=./vpi.vpi
wave:
	gtkwave board_ctrl_tb.vcd 
