/*
  *
  * This is a DMA module for copying graphics data from CPU RAM to the VRAM, and only in this one direction.
  * The DMA is setup by writing to the control registers. 
  * The io_src_addr, i_src_data ports are used for reading the source data from the CPU RAM, and to setup the control registers.
  *
  * Registers:
  * ADDR  |  NAME        | DATA SIZE  | DESCRIPTION
  * 0x00  |  SRC_ADDR_L  | 8 bit      | Source address low byte. X coordinate of upper left corner of the source location.
  * 0x01  |  SRC_ADDR_H  | 5 bit      | Source address high byte. Y coordinate of upper left corner of the source location.
  * 0x02  |  DST_ADDR_L  | 8 bit      | Destination address low byte. X coordinate of lower right corner of the destination location.
  * 0x03  |  DST_ADDR_H  | 8 bit      | Destination address high byte. Y coordinate of lower right corner of the destination location.
  * 0x04  |  WIDTH       | 8 bit      | Width of the image to copy.
  * 0x05  |  HEIGHT      | 5 bit      | Height of the image to copy.
  * 0x06  |  MASK        | 8 bit      | Copy mask, bit 0-4 for X, bit 5-7 for Y. Can be used to repeat the same image multiple times.
  * 0x07  |  STATE       | 1 bit      | Start the DMA operation. Only first bit is used. 0 for copy only non-zero pixels, 1 for copy all pixels.
  * 
  * TBD: Describe the MASK register in more detail.
  *
  * This module is intended to be programmed on the EPM7128S CPLD chip.
  *
  * NOTE about some design choices:
  * - The DMA is designed to be used with the GfxVga module.
  * - The DMA is designed to be simple and to be able to copy the
  *   data from the CPU RAM to the VRAM as fast as possible.
  * - Because of the high CPLD utilization (127 out of 128 macrocells), 
  *   some optimizations had to be made. For example:
  *   a) the synthesis tool better optimized a count-down counter + compare to zero + reset to register,
  *   than a count-up counter + compare to a register + reset to zero. 
  *   b) the source address space was limited to 8Kb, to reduce the number of registers used.
  *   Paging was added to access up to 32Kb of the CPU RAM. However, only 8Kb can be accessed at a time.
  *   c) image repetition was implemented using the MASK register, the simplicity of the implementation allowed to save on macrocells utilization.
  *
  *
  * NOTE about included delays:
  * - The delays are included for simulation purposes only. Synthesis tools will ignore them.
  * - The delays are divided by 3, to match the real measured delays on the CPLD chip.
  * - #(17/3) means 17ns delay predicted by the tool, divided by 3 to match the real delay.
  */

module GfxDma(
  input i_clk,  // main clock signal, 25.175 MHz clock signal
  input i_clk2, // secondary clock signal, 25.175 MHz clock signal, should precede i_clk by ~10ns
  // in parctice, the i_clk is delayed by ~10ns

  // src side interface, used to read the source data from the CPU RAM and to setup the control registers
  input i_src_ce_b,            // source RAM chip or DMA chip enable
  input i_src_ce2_b,           // aux DMA chip enable, can be connected with the GfxVga.i_ctrl_ce2 signal 
  output o_src_re_b,           // source RAM read enable
  inout io_src_we_b,           // source RAM or DMA chip write enable
  inout [12:0] io_src_addr,    // when inactive, the address is used to select the control register, when DMA is active, the address is used to select the source RAM address
  output [1:0] o_src_ram_page, // source RAM page select, 4x8Kb
  input [7:0] i_src_data,      // when inactive, the data is used to write to the control register, when DMA is active, the data is used to read the source RAM data
  
  // dst side interface, used to write the destination data to the VRAM
  output o_dst_we_b,        // destination VRAM write enable
  output [15:0] o_dst_addr, // destination VRAM address
  output [7:0] o_dst_data,  // destination VRAM data
  
  input i_free_vbus_b, // VRAM bus free and ready for DMA access, should be connected with GfxVga.o_free_vbus_b
  output o_active      // DMA active signal, can be used to detach part of the CPU RAM from CPU buss
);

// Constants
// control register addresses
 localparam CTRL_ADDR_SRC_ADDR_L = 3'h0,  // Source address low byte. X coordinate of upper left corner of the source location.
            CTRL_ADDR_SRC_ADDR_H = 3'h1,  // Source address high byte. Y coordinate of upper left corner of the source location.
            CTRL_ADDR_DST_ADDR_L = 3'h2,  // Destination address low byte. X coordinate of lower right corner of the destination location.
            CTRL_ADDR_DST_ADDR_H = 3'h3,  // Destination address high byte. Y coordinate of lower right corner of the destination location.
            CTRL_ADDR_WIDTH = 3'h4,       // Width of the image to copy.
            CTRL_ADDR_HEIGHT = 3'h5,      // Height of the image to copy.
            CTRL_ADDR_MASK = 3'h6,        // Copy mask, bit 0-4 for X, bit 5-7 for Y. Can be used to repeat the same image multiple times.
            CTRL_ADDR_STATE = 3'h7;       // Start the DMA operation. Only first bit is used. 0 for copy only non-zero pixels, 1 for copy all pixels.
				
 localparam CFG_CPY_NON_ZERO = 1'b0, // copy only non-zero pixels
            CFG_CPY_ALL = 1'b1;      // copy all pixels
				
  reg reg_active;                    // DMA state, 1 for active (copying), 0 for inactive (idle)
  reg reg_active_clk1;               // reg_active delayed by one clock cycle
  reg reg_active_clk2;               // reg_active delayed by two clock cycles
  reg reg_ctrl_config;               // DMA configuration register, 1 for copy all pixels, 0 for copy only non-zero pixels
  reg [4:0]  reg_ctrl_cpy_x_mask;    // copy mask
  reg [2:0]  reg_ctrl_cpy_y_mask;    // copy mask
  reg [7:0]  reg_ctrl_src_x_origin;  // source X origin
  reg [4:0]  reg_ctrl_src_y_origin;  // source Y origin
  reg [1:0]  reg_src_ram_page;       // source RAM page
  reg [12:0] reg_ctrl_src_addr;      // current source RAM address, incremented each clock cycle
  reg [7:0]  reg_ctrl_dst_x_origin;  // destination X origin
  reg [7:0]  reg_ctrl_dst_y_origin;  // destination Y origin
  reg [7:0]  reg_ctrl_width;         // image width
  reg [7:0]  reg_ctrl_height;        // image height
  reg [7:0]  reg_x_cnt;              // image width counter
  reg [7:0]  reg_y_cnt;              // image height counter
  reg [15:0] reg_dst_addr_hold;      // destination address hold
  reg [7:0]  reg_dst_data_hold;      // destination data hold
  

  wire [15:0] dst_addr_offset; // calculated destination address offset
  wire dst_out_enabled;        // destination output enabled, used to disable the DST output when the DMA is inactive
  wire skip_cpy;               // skip copying, used to skip copying when the destination data is 0

  // set initial values
  initial begin
    reg_active = 1'b0;
    reg_active_clk1 = 1'b0;
    reg_active_clk2 = 1'b0;

    reg_ctrl_config = 1'b0;
    reg_ctrl_src_addr = 13'h00;
    reg_ctrl_dst_x_origin = 8'h00;
    reg_ctrl_dst_y_origin = 8'h00;
    reg_ctrl_width = 8'h00;
    reg_ctrl_height = 8'h00;
    reg_ctrl_cpy_x_mask = 5'h1f;
    reg_ctrl_cpy_y_mask = 3'h7;

    reg_x_cnt = 8'h0;
    reg_y_cnt = 8'h0;

    reg_dst_addr_hold = 16'h00;
    reg_dst_data_hold = 8'h00;
  end
  
  assign dst_out_enabled = reg_active_clk2 && ~i_free_vbus_b; // determine if the destination output should be enabled
  assign skip_cpy = reg_ctrl_config == CFG_CPY_NON_ZERO && reg_dst_data_hold  == 8'h00; // determine if the copying should be skipped
  
  assign #(8/3) io_src_addr = reg_active_clk1 ? reg_ctrl_src_addr : 13'hz; 
  assign #(8/3) o_src_ram_page = reg_active_clk1 ? reg_src_ram_page : 2'hz;
  assign #(17/3) o_src_re_b = reg_active_clk1 ? 1'h0 : 1'hz; 
  assign #(15/3) io_src_we_b = reg_active_clk1 ? 1'h1 : 1'hz;
  
  assign dst_addr_offset = {reg_ctrl_dst_y_origin - reg_y_cnt, reg_ctrl_dst_x_origin - reg_x_cnt}; // calculate the destination address offset
  assign #(0) o_dst_addr =  dst_out_enabled ? reg_dst_addr_hold : 16'hz; 
  assign #(0) o_dst_data = dst_out_enabled ? reg_dst_data_hold  : 8'hz;
  assign #(18/3) o_dst_we_b = ~dst_out_enabled | skip_cpy | i_clk2;  // write to the VRAM in the low phase of the clock cycle
  assign #(8/3) o_active = reg_active;
  
  always @ (posedge i_clk)  begin
    reg_active_clk2 <= reg_active_clk1;
    reg_active_clk1 <= reg_active;

    if (reg_active_clk1 && ~i_free_vbus_b) begin 
      #(8/3) reg_dst_addr_hold <= dst_addr_offset;  //hold the destination address
      #(8/3) reg_dst_data_hold <= i_src_data;  //hold the destination data, will be written to the VRAM in the low phase of this clock cycle
      
      // the DMA is copying the data column-wise, from bottom to top, from right to left
      if (reg_y_cnt == 8'h00)  begin
        // reset the Y counter and move the X counter to the left
        #(12/3) reg_y_cnt = reg_ctrl_height;
        #(28/3) reg_x_cnt <= reg_x_cnt - 1'h1;
        #(12/3) reg_ctrl_src_addr[12:8] <= reg_ctrl_src_y_origin;
        #(12/3) reg_ctrl_src_addr[7:0] <= (reg_ctrl_src_addr[7:0] + 1'h1) & {1'b1, reg_ctrl_cpy_x_mask[4:1], 2'b11, reg_ctrl_cpy_x_mask[0]}; 
        if (reg_x_cnt == 8'h00) begin // note: for simulation has to be reg_x_cnt == 8'hff
          // last column copied, deactivate the DMA
          reg_active_clk1 <= 1'h0;
          reg_active <= 1'h0;
        end
      end else begin
        // advance copying the column
        #(21/3) reg_y_cnt <= reg_y_cnt - 1'h1; 
        #(12/3) reg_ctrl_src_addr[12:8] <= (reg_ctrl_src_addr[12:8] + 1'h1) & {reg_ctrl_cpy_y_mask[2], reg_ctrl_cpy_y_mask[1], 2'b11, reg_ctrl_cpy_y_mask[0]};
      end
		
    // store control registers
    end else if (~i_src_ce_b && ~i_src_ce2_b && ~io_src_we_b) begin
       case (io_src_addr[2:0])
         CTRL_ADDR_SRC_ADDR_L: reg_ctrl_src_x_origin <= i_src_data;
         CTRL_ADDR_SRC_ADDR_H:  begin
           reg_ctrl_src_y_origin <= i_src_data[4:0];
           reg_src_ram_page <= i_src_data[6:5];
         end
         CTRL_ADDR_DST_ADDR_L: reg_ctrl_dst_x_origin <= i_src_data;
         CTRL_ADDR_DST_ADDR_H: reg_ctrl_dst_y_origin <= i_src_data;
         CTRL_ADDR_WIDTH: reg_ctrl_width <= i_src_data;
         CTRL_ADDR_HEIGHT: reg_ctrl_height <= i_src_data;
         CTRL_ADDR_MASK: begin
           reg_ctrl_cpy_x_mask <= i_src_data[4:0];
           reg_ctrl_cpy_y_mask <= i_src_data[7:5];
         end
         CTRL_ADDR_STATE: begin
           // reset the counters 
           reg_x_cnt <= reg_ctrl_width;
           reg_y_cnt <= reg_ctrl_height;
           // reset the source address counter
           reg_ctrl_src_addr[7:0] <= reg_ctrl_src_x_origin;
           reg_ctrl_src_addr[12:8] <= reg_ctrl_src_y_origin;
           // store the copy configuration
           reg_ctrl_config <= i_src_data[0];
           // activate the DMA
           reg_active <= 1'h1;
         end
       endcase
     end 
  end

endmodule
