## Generated SDC file "VgaGen.sdc"

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

## DATE    "Sun Sep 29 18:34:56 2024"

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

create_clock -name {i_clk} -period 39.720 -waveform { 0.000 19.860 } [get_ports {i_clk}]
create_clock -name {i_ctrl_w_b} -period 79.440 -waveform { 0.000 39.720 } [get_ports {i_ctrl_w_b}]

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
set_input_delay -clock [get_clocks {i_clk}] 0.000 [get_ports {i_vdata[*]}]

set_input_delay -clock [get_clocks {i_ctrl_w_b}] 0.000 [get_ports {i_ctrl_ce_b}]
set_input_delay -clock [get_clocks {i_ctrl_w_b}] 0.000 [get_ports {i_ctrl_ce2}]
set_input_delay -clock [get_clocks {i_ctrl_w_b}] 0.000 [get_ports {i_ctrl_w_b}]
set_input_delay -clock [get_clocks {i_ctrl_w_b}] 0.000 [get_ports {i_ctrl_addr[*]}]
set_input_delay -clock [get_clocks {i_ctrl_w_b}] 0.000 [get_ports {i_ctrl_data[*]}]


#**************************************************************
# Set Output Delay
#**************************************************************
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_active_b}]

set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_rgb_data[*]}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_palette[*]}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_vga_latch}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_vga_out_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_free_vbus_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_enabled_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_frame_start_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_frame_progress_b}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_frame_end_b}]

set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_hsync}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_vsync}]

set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_vaddr[*]}]
set_output_delay -add_delay  -clock [get_clocks {i_clk}]  0.000 [get_ports {o_vaddr15_b}]

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

