  /*
  *
  * This is a VGA signal generator module. It generates a VGA image from 64Kb of video RAM, and 1K of color map RAM.
  * The module is controlled by a set of registers, which can be accessed through a control interface.
  * The generated VGA signal is of 640x480 resolution, with 60Hz refresh rate.
  * There are two resolution modes available: color mode and double resolution mode.
  * Color mode: 256x240 resolution, 8-bit color depth. 
  * Double resolution mode: 512x480 resolution, 2-bit color depth. 
  *
  * The output image can be shifted in the horizontal and vertical by two 8-bit registers.
  * When in high resolution mode, the image is shifted by 4 pixels (4 pixels per byte).
  *
  * The o_pixel_data can be directly converted to VGA signals using a resistor ladder, or a DAC, 
  * or can be used to drive a color mapper RAM, however this will require additional logic. 
  *
  * It is highly recommended to strobe all VGA signals on the rising edge of the o_vga_latch signal.
  *
  * This module is intended to be programmed on the EPM7128S-15ns CPLD chip.
  *
  * NOTE about included delays:
  * - The delays are included for simulation purposes only. Synthesis tools will ignore them.
  * - The delays are divided by 3, to match the real measured delays on the CPLD chip.
  * - #(17/3) means 17ns delay predicted by the tool, divided by 3 to match the real delay.
  *
  */


module GfxVga (
  input wire         i_clk, // 25.175 MHz clock signal 

  // chip control interface
  input wire         i_ctrl_ce_b,  // chip enabled
  input wire         i_ctrl_w_b,   // control write signal
  input wire  [1:0]  i_ctrl_addr,  // control register address
  input wire  [7:0]  i_ctrl_data,  // control data

  // vram interface
  output wire [15:0] o_vaddr,     // 64Kb video RAM pixel address
  input wire  [7:0]  i_vdata,     // video RAM pixel data

  // other control and status signals
  output wire        o_enabled_b,        // goes low when the VGA is enabled 
  output wire        o_frame_start_b,    // goes low for the whole duration of first visible scan-line (0). Can be used to trigger an interrupt
  output wire        o_frame_progress_b, // goes low for the whole duration of every 32nd scan-line. Can be used to trigger an interrupt 
  output wire        o_frame_end_b,      // goes low for the whole duration of the first invisible scan-line (480). Can be used to trigger an interrupt
  output wire        o_free_vbus,        // VRAM bus free and ready for DMA access
  output wire        o_vga_out_b,        // enable VGA color output, active during displaying the actual image
  output wire        o_vga_latch,        // VGA signal latch signal (latch strobe)
  output wire [7:0]  o_pixel_data,       // pixel output to palette ram
  output wire [1:0]  o_palette,          // palette selector
  output wire        o_hsync,            // VGA horizontal sync signal
  output wire        o_vsync             // VGA vertical sync signal

);

  // Constants
  localparam H_CENTER_SHIFT    = 10'd64;  // shift the 512 width image to the center, todo: probably needs adjustment
  localparam H_CNT_RST         = 10'd799; // horizontal pixel counter reset threshold 
  localparam V_CNT_RST         = 10'd525; // vertical line counter reset threshold
  localparam HSYNC_START       = 10'd656; // horizontal VGA sync signal start
  localparam HSYNC_END         = 10'd752; // horizontal VGA sync signal end
  localparam VSYNC_START       = 10'd490; // vertical VGA sync signal start
  localparam VSYNC_END         = 10'd492; // vertical VGA sync signal end
  localparam ACTIVE_H_END      = 10'd515; // chip active signal end; for each scan line; relative to the shifted h_cnt
  localparam ACTIVE_V_END      = 10'd480; // chip active signal end; for each frame; relative to the shifted h_cnt
  localparam FRAME_START_LINE  = 10'd000; // frame start interrupt signal 
  localparam FRAME_END_LINE    = 10'd480; // frame end interrupt signal
  localparam VGA_OUT_START     = 10'd68;  // VGA output enabled start; 
  localparam VGA_OUT_END       = 10'd579; // VGA output enabled end; 

  // Control address parameters
  localparam CTRL_ADDR_STATUS  = 2'b00; // control address for status
  localparam CTRL_ADDR_X_SHIFT = 2'b01; // control address for X shift
  localparam CTRL_ADDR_Y_SHIFT = 2'b10; // control address for Y shift
  localparam CTRL_ADDR_PALETTE = 2'b11; 

  // Control flags
  localparam CTRL_ENABLE     = 2'b00; // control enable flag
  localparam CTRL_DOUBLE_RES = 2'b01; // control double resolution flag

  // Internal registers
  reg [1:0] reg_ctrl_status;  // chip status: {enabled, double resolution mode}
  reg [7:0] reg_ctrl_x_shift; // image shift in the horizontal axis
  reg [7:0] reg_ctrl_y_shift; // image shift in the vertical axis
  reg [1:0] reg_ctrl_palette; // selected color palette

  reg [9:0] reg_h_cnt; // horizontal counter
  reg [9:0] reg_v_cnt; // vertical counter

  reg [7:0] reg_vdata; // pixel data store; will be passed to the color palette ram 

  // Internal control signals
  wire ctrl_enable;     // chip enabled
  wire ctrl_double_res; // double resolution mode enalbed; used for text mode

  wire h_cnt_rst; // reset horizontal counter signal
  wire v_cnt_rst; // reset vertical counter signal

  wire active;
  wire [9:0] h_cnt_shifted;  // shifted horizontal count; used for center the 512 pix image in a 640 resolution
  wire [15:0] vaddr;         // computed VRAM pixel address
  wire [15:0] vaddr_shifted; // vaddr shifted by reg_ctrl_x_shift and reg_ctrl_y_shift
  wire [7:0] dbr_pixel;        // VRAM pixel mapped to a double resolution 
  wire [7:0] dbr_pixel0, dbr_pixel1, dbr_pixel2, dbr_pixel3; 
  
  // Assign internal control signals from status register
  assign ctrl_enable     = reg_ctrl_status[CTRL_ENABLE];
  assign ctrl_double_res = reg_ctrl_status[CTRL_DOUBLE_RES];

  assign h_cnt_rst = (reg_h_cnt == H_CNT_RST);
  assign v_cnt_rst = (reg_v_cnt == V_CNT_RST);
  assign h_cnt_shifted = reg_h_cnt - H_CENTER_SHIFT;

  assign #(17/3) o_hsync = (reg_h_cnt < HSYNC_START) || (reg_h_cnt >= HSYNC_END);
  assign #(17/3) o_vsync = (reg_v_cnt < VSYNC_START) || (reg_v_cnt >= VSYNC_END);

  assign #(17/3) o_enabled_b = !ctrl_enable; 
  assign active = (h_cnt_shifted < ACTIVE_H_END) && (reg_v_cnt < ACTIVE_V_END);
  
  assign #(19/3) o_free_vbus = ~active || reg_h_cnt[0];

  // Video address calculation with shifting
  assign vaddr = ctrl_double_res ? {reg_v_cnt[8:0], h_cnt_shifted[8:2]} : {reg_v_cnt[8:1], h_cnt_shifted[8:1]};
  assign vaddr_shifted = {(vaddr[15:8] + reg_ctrl_y_shift), (vaddr[7:0] + reg_ctrl_x_shift)};
 
  assign #(36/3) o_vaddr = (active && ctrl_enable) ? vaddr_shifted: 16'hzzzz;

  // Double resolution mapping
  assign dbr_pixel0 = {6'h00, reg_vdata[1:0]};
  assign dbr_pixel1 = {6'h00, reg_vdata[3:2]};
  assign dbr_pixel2 = {6'h00, reg_vdata[5:4]};
  assign dbr_pixel3 = {6'h00, reg_vdata[7:6]};
  assign dbr_pixel = reg_h_cnt[1] ? (reg_h_cnt[0] ? dbr_pixel0 : dbr_pixel3) : (reg_h_cnt[0] ? dbr_pixel2 : dbr_pixel1);

  assign #(17/3) o_pixel_data =  ctrl_double_res ? dbr_pixel : reg_vdata;

  assign #(8/3) o_palette = reg_ctrl_palette; 
  assign #(17/3) o_vga_latch =  ctrl_double_res ? i_clk : ~reg_h_cnt[0];
  assign #(17/3) o_vga_out_b = (reg_h_cnt < VGA_OUT_START || reg_h_cnt > VGA_OUT_END) || ~ctrl_enable;
  
  assign #(17/3) o_frame_start_b = reg_v_cnt != FRAME_START_LINE || ~ctrl_enable;
  assign #(17/3) o_frame_progress_b = reg_v_cnt[4:0] != 5'h0 || ~ctrl_enable;
  assign #(17/3) o_frame_end_b = reg_v_cnt != FRAME_END_LINE || ~ctrl_enable;

  initial begin
    reg_ctrl_status  = 2'h0;
    reg_ctrl_x_shift = 8'h00;
    reg_ctrl_y_shift = 8'h00;
    reg_ctrl_palette = 2'b00;
    reg_h_cnt        = 10'h000;
    reg_v_cnt        = 10'd000;
    reg_vdata        = 8'h00;
  end


  always @ (negedge i_ctrl_w_b) begin 
    if (~i_ctrl_ce_b) begin 
      case (i_ctrl_addr)
        CTRL_ADDR_STATUS: reg_ctrl_status  <= i_ctrl_data[1:0];
        CTRL_ADDR_X_SHIFT: reg_ctrl_x_shift  <= i_ctrl_data;
        CTRL_ADDR_Y_SHIFT: reg_ctrl_y_shift  <= i_ctrl_data;
        CTRL_ADDR_PALETTE: reg_ctrl_palette  <= i_ctrl_data[1:0];
      endcase
    end
  end

  always @ (posedge i_clk) begin 
    // Horizontal counter logic
    if (h_cnt_rst) begin 
      #(3) reg_h_cnt <= 10'h0;
      #(13/3) reg_v_cnt <= reg_v_cnt + 10'h1;
    end else if (ctrl_enable) begin 
      #(12/3) reg_h_cnt <= reg_h_cnt + 10'h1;
    end

    // Vertical counter logic
    if (v_cnt_rst) begin 
      #(3) reg_v_cnt <= 10'h0;
    end 
  end

  // i_vdata_delayed has been introduced for simulation purposes only
  // just to add a simulation delay
  wire [7:0] i_vdata_delayed;
  assign #(9/3) i_vdata_delayed = i_vdata; 

  always @ (posedge i_clk) begin 
    if (ctrl_enable && active) begin  //note: we don't take this delays into account during simulation
      if (reg_h_cnt[0]) begin
        #(3) reg_vdata <= i_vdata_delayed; 
      end
    end else begin 
      #(3) reg_vdata <= 8'h00;
    end
  end
endmodule
