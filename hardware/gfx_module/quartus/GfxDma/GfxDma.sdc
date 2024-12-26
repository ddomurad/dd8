## Generated SDC file "vga_dma.sdc"

## Copyright (C) 1991-2013 Altera Corporation
## Your use of Altera Corporation's design tools, logic functions 
## and other software and tools, and its AMPP partner logic 
## functions, and any output files from any of the foregoing 
## (including device programming or simulation files), and any 
## associated documentation or information are expressly subject 
## to the terms and conditions of the Altera Program License 
## Subscription Agreement, Altera MegaCore Function License 
## Agreement, or other applicable license agreement, including, 
## without limitation, that your use is for the sole purpose of 
## programming logic devices manufactured by Altera and sold by 
## Altera or its authorized distributors.  Please refer to the 
## applicable agreement for further details.


## VENDOR  "Altera"
## PROGRAM "Quartus II"
## VERSION "Version 13.0.1 Build 232 06/12/2013 Service Pack 1 SJ Web Edition"

## DATE    "Sun Oct 06 11:32:47 2024"

##
## DEVICE  "EPM7128SLC84-15"
##


#**************************************************************
# Time Information
#**************************************************************

set_time_format -unit ns -decimal_places 3



#**************************************************************
# Create Clock
#**************************************************************

create_clock -name {i_clk} -period 38.000 -waveform { 0.000 19.000 } [get_ports {i_clk}]
create_clock -name {i_clk2} -period 38.000 -waveform { 0.000 19.000 } [get_ports {i_clk2}]


  
#**************************************************************
# Create Generated Clock
#**************************************************************



#**************************************************************
# Set Clock Latency
#**************************************************************



#**************************************************************
# Set Clock Uncertainty
#**************************************************************



#**************************************************************
# Set Input Delay
#**************************************************************

set_input_delay  -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {i_clk3}]
set_input_delay  -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {i_src_ce_b}]
set_input_delay  -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {io_src_we_b}]
set_input_delay  -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {io_src_addr[*]}]
set_input_delay  -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {i_src_data[*]}]

set_input_delay  -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {i_free_vbus}]


  
#**************************************************************
# Set Output Delay
#**************************************************************
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_src_re_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {io_src_we_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {io_src_addr[*]}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_src_ram_page[*]}]

#set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_src_ram_ce_b}]

set_output_delay -add_delay  -clock [get_clocks {i_clk2}]  0.000 [get_ports {o_dst_we_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_dst_addr[*]}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_dst_data[*]}]

set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_active}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_addr_sel}]


#**************************************************************
# Set Clock Groups
#**************************************************************



#**************************************************************
# Set False Path
#**************************************************************



#**************************************************************
# Set Multicycle Path
#**************************************************************



#**************************************************************
# Set Maximum Delay
#**************************************************************



#**************************************************************
# Set Minimum Delay
#**************************************************************



#**************************************************************
# Set Input Transition
#**************************************************************

