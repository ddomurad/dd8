STD = --std=08	

all: build run
build:
	ghdl -a $(STD) ./ic_selector.vhdl 
	ghdl -a $(STD) ./ic_selector_tb.vhdl 
	ghdl -e $(STD) ic_selector
	ghdl -e $(STD) ic_selector_tb
run:
	ghdl -r $(STD) ic_selector_tb --vcd=ic_selector_tb.vcd
wave:
	gtkwave ic_selector_tb.vcd 
