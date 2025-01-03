/*
  *
  * This is a DMA module for copying graphics data from CPU RAM to the VRAM, and only in this one direction.
  * The DMA is setup by writing to the control registers. 
  * The io_src_addr and i_src_data ports are used for reading the source data from the CPU RAM, and to setup the control registers.
  *
  * Registers:
  * ADDR  |  NAME        | DATA SIZE  | DESCRIPTION
  * 0x00  |  SRC_ADDR_L  | 8 bits     | Source address low byte. X coordinate of the upper left corner of the source location.
  * 0x01  |  SRC_ADDR_H  | 7 bits     | Source address high byte. b[4:0] Y coordinate of the upper left corner of the source location. b[6:5] 8K source RAM bank.
  * 0x02  |  DST_ADDR_L  | 8 bits     | Destination address low byte. X coordinate of the lower right corner of the destination location.
  * 0x03  |  DST_ADDR_H  | 8 bits     | Destination address high byte. Y coordinate of the lower right corner of the destination location.
  * 0x04  |  WIDTH       | 8 bits     | Width of the image to copy.
  * 0x05  |  HEIGHT      | 8 bits     | Height of the image to copy.
  * 0x06  |  MASK        | 8 bits     | Copy mask, bit 0-4 for X, bit 5-7 for Y. Can be used to repeat the same image multiple times.
  * 0x07  |  STATE       | 3 bits     | Start the DMA operation. b[0]: 0 for copy only non-zero pixels, 1 for copy all pixels. b[2:1] sets reg_ctrl_data_mask. 
  * 
  * 
  * TBD: Describe the MASK register in more detail.
  * TBD: Describe the reg_ctrl_data_mask in more detail.
  *
  * This module is intended to be programmed on the EPM7128S-15ns CPLD chip.
  *
  *
  * NOTE about some design choices:
  * - The DMA is designed to be used with the GfxVga module.
  * - The DMA is designed to be simple and to be able to copy the
  *   data from the CPU RAM to the VRAM as fast as possible.
  * - Because of the high CPLD utilization (>95% macro-cell utilization), 
  *   some optimizations had to be made. For example:
  *   a) the synthesis tool better optimized a count-down counter + compare to zero + reset to register,
  *   than a count-up counter + compare to a register + reset to zero. 
  *   b) the source address space was limited to 8Kb, to reduce the number of registers needed.
  *   Paging was added to access up to 32Kb of the CPU RAM. However, only 8Kb can be accessed at a time.
  *   c) image repetition was implemented using the MASK register, the simplicity of the implementation allowed to save on macro-cell utilization.
  *
  *
  * NOTE about included delays:
  * - The delays are included for simulation purposes only. Synthesis tools will ignore them.
  * - The delays are divided by 3, to match the real measured delays on the CPLD chip.
  * - #(17/3) means 17ns delay predicted by the tool, divided by 3 to match the real delay.
  */

module GfxDma(
  input i_clk,  // main clock signal, 25.175 MHz clock signal
  input i_clk2, // secondary clock signal, 25.175 MHz clock signal, should precede i_clk by ~5-10ns
  input i_clk3, // alternate clk path, to allow for shorter routing, and shorter delays
  // in practice, the i_clk, and i_clk3 are delayed by ~5-10ns

  // src side interface, used to read the source data from the CPU RAM and to setup the control registers
  input i_src_ce_b,            // source RAM chip or DMA chip enable
  inout io_src_we_b,           // source RAM or DMA chip write enable
  inout [12:0] io_src_addr,    // when inactive, the address is used to select the control register, when DMA is active, the address is used to select the source RAM address
  output [1:0] o_src_ram_page, // source RAM page select, 4x8Kb
  inout [7:0] i_src_data,      // when inactive, the data is used to write to the control register, when DMA is active, the data is used to read the source RAM data
  
  // dst side interface, used to write the destination data to the VRAM
  output o_dst_we_b,        // destination VRAM write enable
  output [15:0] o_dst_addr, // destination VRAM address
  output [7:0] o_dst_data,  // destination VRAM data
  
  input i_free_vbus, // VRAM bus free and ready for DMA access, should be connected with GfxVga.o_free_vbus_b
  output o_addr_sel,  // VRAM address select signal, DMA line vs VGA line, high for DMA
  output o_active   // DMA active signal, can be used to detach part of the CPU RAM from CPU bus
);

// control register addresses
 localparam CTRL_ADDR_SRC_ADDR_L = 3'h0,  // Source address low byte. X coordinate of upper left corner of the source location.
            CTRL_ADDR_SRC_ADDR_H = 3'h1,  // Source address high byte. Y coordinate of upper left corner of the source location.
            CTRL_ADDR_DST_ADDR_L = 3'h2,  // Destination address low byte. X coordinate of lower right corner of the destination location.
            CTRL_ADDR_DST_ADDR_H = 3'h3,  // Destination address high byte. Y coordinate of lower right corner of the destination location.
            CTRL_ADDR_WIDTH = 3'h4,       // Width of the image to copy.
            CTRL_ADDR_HEIGHT = 3'h5,      // Height of the image to copy.
            CTRL_ADDR_MASK = 3'h6,        // Copy mask, bit 0-4 for X, bit 5-7 for Y. Can be used to repeat the same image multiple times.
            CTRL_ADDR_STATE = 3'h7;       // Start the DMA operation. Only first bit is used. 0 for copy only non-zero pixels, 1 for copy all pixels.

// DMA configuration
 localparam CFG_CPY_NON_ZERO = 1'b0, // copy only non-zero pixels
            CFG_CPY_ALL = 1'b1;      // copy all pixels
				
  reg reg_active;                    // DMA state, 1 for active (copying), 0 for inactive (idle)
  reg reg_active_clk1;               // reg_active delayed by one clock cycle
  reg reg_active_clk2;               // reg_active delayed by two clock cycles
  reg reg_free_vbus;                 // VRAM bus free and ready for DMA access
  reg reg_ctrl_config;               // DMA configuration register, 1 for copy all pixels, 0 for copy only non-zero pixels
  reg [1:0]  reg_ctrl_data_mask;     // can be used to set two most significant bits of the destination data
  reg [4:0]  reg_ctrl_cpy_x_mask;    // copy mask
  reg [2:0]  reg_ctrl_cpy_y_mask;    // copy mask
  reg [7:0]  reg_ctrl_src_x_origin;  // source X origin
  reg [4:0]  reg_ctrl_src_y_origin;  // source Y origin
  reg [1:0]  reg_ctrl_src_ram_page;  // source RAM page
  reg [7:0]  reg_ctrl_dst_x_origin;  // destination X origin
  reg [7:0]  reg_ctrl_dst_y_origin;  // destination Y origin
  reg [7:0]  reg_ctrl_width;         // image width
  reg [7:0]  reg_ctrl_height;        // image height
  reg [7:0]  reg_x_cnt;              // image width counter
  reg [7:0]  reg_y_cnt;              // image height counter
  reg [12:0] reg_src_addr;           // current source RAM address, incremented each clock cycle
  reg [15:0] reg_dst_addr_hold;      // destination address hold
  reg [7:0]  reg_dst_data_hold;      // destination data hold

  wire [15:0] dst_addr_offset;  // calculated destination address offset
  wire dst_out_enabled;         // destination output enabled, used to disable the DST output when the DMA is inactive
  wire skip_cpy;                // skip copying, used to skip copying when the destination data is 0

  // set initial values
  initial begin
    reg_active = 1'b0;
    reg_active_clk1 = 1'b0;
    reg_active_clk2 = 1'b0;
    reg_free_vbus = 1'b0;

    reg_ctrl_config = 1'b0;
    reg_ctrl_data_mask = 2'h0;
    reg_ctrl_cpy_x_mask = 5'h1f;
    reg_ctrl_cpy_y_mask = 3'h7;
    reg_ctrl_src_x_origin = 8'h00;
    reg_ctrl_src_y_origin = 5'h00;
    reg_ctrl_src_ram_page = 2'h0;
    reg_ctrl_dst_x_origin = 8'h00;
    reg_ctrl_dst_y_origin = 8'h00;
    reg_ctrl_width = 8'h00;
    reg_ctrl_height = 8'h00;

    reg_x_cnt = 8'h0;
    reg_y_cnt = 8'h0;

    reg_src_addr = 13'h00;
    reg_dst_addr_hold = 16'h00;
    reg_dst_data_hold = 8'haa;
  end
  
  assign dst_out_enabled = reg_active_clk2 && reg_free_vbus; // determine if the destination output should be enabled
  assign skip_cpy = reg_ctrl_config == CFG_CPY_NON_ZERO && reg_dst_data_hold  == 8'h00; // determine if the copying should be skipped
  
  assign #(8/3) io_src_addr = reg_active_clk1 ? reg_src_addr : 13'hz; 
  assign #(8/3) o_src_ram_page = reg_active_clk1 ? reg_ctrl_src_ram_page : 2'hz;
  assign #(15/3) io_src_we_b = reg_active_clk1 ? 1'h1 : 1'hz; 
  
  assign dst_addr_offset = {reg_ctrl_dst_y_origin - reg_y_cnt, reg_ctrl_dst_x_origin - reg_x_cnt}; // calculate the destination address offset
  assign #(0) o_dst_addr =  reg_dst_addr_hold;
  assign #(0) o_dst_data = dst_out_enabled && !i_clk3 ? reg_dst_data_hold  : 8'hz;
  assign #(18/3) o_dst_we_b = ~dst_out_enabled || skip_cpy || i_clk2 ;  // write to the VRAM in the low phase of the clock cycle
  assign #(8/3) o_active = reg_active;
  assign #(15/3) o_addr_sel = reg_free_vbus;
  
  always @ (posedge i_clk)  begin
    reg_free_vbus <= i_free_vbus;
    reg_active_clk2 <= reg_active_clk1;
    reg_active_clk1 <= reg_active;

    if (reg_active_clk1 && i_free_vbus) begin 
      #(8/3) reg_dst_addr_hold <= dst_addr_offset;  //hold the destination address
      #(8/3) reg_dst_data_hold <= i_src_data | {reg_ctrl_data_mask, 6'h00};  //hold the destination data, will be written to the VRAM in the low phase of this clock cycle
      
      // the DMA is copying the data column-wise, from top to bottom, from left to right
      // if (reg_y_cnt == 8'hff)  begin
      if (reg_y_cnt == 8'h00)  begin //note: might require 8'hff for synthesis
        // reset the Y counter and move the X counter to the left
        #(12/3) reg_y_cnt = reg_ctrl_height;
        #(28/3) reg_x_cnt <= reg_x_cnt - 1'h1;
        #(12/3) reg_src_addr[12:8] <= reg_ctrl_src_y_origin;
        #(12/3) reg_src_addr[7:0] <= (reg_src_addr[7:0] + 1'h1) & {1'b1, reg_ctrl_cpy_x_mask[4:1], 2'b11, reg_ctrl_cpy_x_mask[0]}; 
        if (reg_x_cnt == 8'hff) begin // note: for simulation has to be reg_x_cnt == 8'hff, 8'h00 for synthesis
          // last column copied, deactivate the DMA
          reg_active_clk1 <= 1'h0;
          reg_active <= 1'h0;
        end
      end else begin
        // advance copying the column
        #(21/3) reg_y_cnt <= reg_y_cnt - 1'h1; 
        #(12/3) reg_src_addr[12:8] <= (reg_src_addr[12:8] + 1'h1) & {reg_ctrl_cpy_y_mask[2], reg_ctrl_cpy_y_mask[1], 2'b11, reg_ctrl_cpy_y_mask[0]};
      end
		
    // store control registers
    end else if (~i_src_ce_b && ~io_src_we_b) begin
       case (io_src_addr[2:0])
         CTRL_ADDR_SRC_ADDR_L: reg_ctrl_src_x_origin <= i_src_data;
         CTRL_ADDR_SRC_ADDR_H:  begin
           reg_ctrl_src_y_origin <= i_src_data[4:0];
           reg_ctrl_src_ram_page <= i_src_data[6:5];
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
           reg_src_addr[7:0] <= reg_ctrl_src_x_origin;
           reg_src_addr[12:8] <= reg_ctrl_src_y_origin;
           // store the copy configuration
           reg_ctrl_config <= i_src_data[0];
           reg_ctrl_data_mask <= i_src_data[2:1];
           // activate the DMA
           reg_active <= 1'h1;
         end
       endcase
     end 
  end

endmodule
