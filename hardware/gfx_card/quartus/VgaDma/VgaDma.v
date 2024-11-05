module VgaDma(
  input i_clk,
  input i_clk2,

  // src side interface
  input i_src_ce_b,
  input i_src_ce2_b,
  output o_src_re_b,
  inout io_src_we_b,
  // output o_src_ram_ce_b, 
  inout [12:0] io_src_addr, 
  output [1:0] o_src_ram_page,
  input [7:0] i_src_data,
  
  // dst side interface
  output o_dst_we_b,
  output [15:0] o_dst_addr,
  output [7:0] o_dst_data,
  
  input i_free_vbus_b,
  output o_active
);

 localparam CTRL_ADDR_SRC_ADDR_L = 3'h0,
            CTRL_ADDR_SRC_ADDR_H = 3'h1,
            CTRL_ADDR_DST_ADDR_L = 3'h2,
            CTRL_ADDR_DST_ADDR_H = 3'h3,
            CTRL_ADDR_WIDTH = 3'h4,
            CTRL_ADDR_HEIGHT = 3'h5,
            CTRL_ADDR_MASK = 3'h6,
            CTRL_ADDR_STATE = 3'h7;
				
 localparam CFG_CPY_NON_ZERO = 1'b0,
            CFG_CPY_ALL = 1'b1;
				
  reg reg_active;
  reg reg_ctrl_config;
  reg [4:0]  reg_ctrl_cpy_x_mask;
  reg [2:0]  reg_ctrl_cpy_y_mask;
  reg [7:0]  reg_ctrl_src_x_origin; 
  reg [4:0]  reg_ctrl_src_y_origin;
  reg [1:0]  reg_src_ram_page;      
  reg [12:0] reg_ctrl_src_addr;
  reg [7:0]  reg_ctrl_dst_x_origin;
  reg [7:0]  reg_ctrl_dst_y_origin;
  reg [7:0]  reg_ctrl_width;
  reg [7:0]  reg_ctrl_height;       
  reg [7:0]  reg_x_cnt;
  reg [7:0]  reg_y_cnt;
  

  wire [15:0] dst_addr_offset;
  wire dst_out_enabled;
  wire skip_cpy;
  wire src_ce;

  // stage wires and regs
  reg reg_active_stage;
  reg reg_active_prestage;
  reg [15:0] reg_dst_addr_staged;
  reg [7:0] reg_dst_data_staged;

  initial begin
    reg_ctrl_config = 1'b0;
    reg_ctrl_src_addr = 13'haa;
    reg_ctrl_dst_x_origin = 8'h00;
    reg_ctrl_dst_y_origin = 8'h00;
    reg_ctrl_width = 8'h00;
    reg_ctrl_height = 8'h00;
    reg_x_cnt = 8'h0;
    reg_y_cnt = 8'h0;
    reg_active = 1'b0;
    reg_ctrl_cpy_x_mask = 5'h1f;
    reg_ctrl_cpy_y_mask = 3'h7;

    reg_active_stage = 1'b0;
    reg_active_prestage = 1'b0;
    reg_dst_addr_staged = 16'h00;
    reg_dst_data_staged = 8'h00;
  end
  
  assign src_ce = ~i_src_ce_b && ~i_src_ce2_b; 
  assign dst_out_enabled = reg_active_stage && ~i_free_vbus_b;
  assign skip_cpy = reg_ctrl_config == CFG_CPY_NON_ZERO && reg_dst_data_staged  == 8'h00;
  
  assign #(8/3) io_src_addr = reg_active ? reg_ctrl_src_addr : 13'hz; // #15ns reg_active
  assign #(8/3) o_src_ram_page = reg_active ? reg_src_ram_page : 2'hz; // #15ns reg_active
  assign #(17/3) o_src_re_b = reg_active ? 1'h0 : 1'hz; 
  assign #(15/3) io_src_we_b = reg_active ? 1'h1 : 1'hz;
  
  assign dst_addr_offset = {reg_ctrl_dst_y_origin - reg_y_cnt, reg_ctrl_dst_x_origin - reg_x_cnt};
  assign #(0) o_dst_addr =  dst_out_enabled ? reg_dst_addr_staged : 16'hz; 
  assign #(0) o_dst_data = dst_out_enabled ? reg_dst_data_staged  : 8'hz; //#24ns reg_active
  assign #(18/3) o_dst_we_b = ~dst_out_enabled | skip_cpy | i_clk2; 
  assign #(8/3) o_active = reg_active_prestage;
  
  always @ (posedge i_clk)  begin
    reg_active_stage <= reg_active;
    reg_active <= reg_active_prestage;

    if (reg_active && ~i_free_vbus_b) begin 
      #(8/3) reg_dst_addr_staged <= dst_addr_offset; //todo: ths #ns is a gues
      #(8/3) reg_dst_data_staged <= i_src_data; //todo: ths #ns is a gues

      if (reg_y_cnt == 8'h00)  begin
        #(12/3) reg_y_cnt = reg_ctrl_height; //#31ns i_free_vbus_b reg_active
        #(28/3) reg_x_cnt <= reg_x_cnt - 1'h1; //#38ns i_free_vbus_b reg_active
        #(12/3) reg_ctrl_src_addr[12:8] <= reg_ctrl_src_y_origin;  //#38ns i_free_vbus_b reg_active
        #(12/3) reg_ctrl_src_addr[7:0] <= (reg_ctrl_src_addr[7:0] + 1'h1) & {1'b1, reg_ctrl_cpy_x_mask[4:1], 2'b11, reg_ctrl_cpy_x_mask[0]}; 
        if (reg_x_cnt == 8'h00) begin 
          reg_active <= 1'h0;
          reg_active_prestage <= 1'h0;
        end
      end else begin
        #(21/3) reg_y_cnt <= reg_y_cnt - 1'h1; //#31ns i_free_vbus_b reg_active
        #(12/3) reg_ctrl_src_addr[12:8] <= (reg_ctrl_src_addr[12:8] + 1'h1) & {reg_ctrl_cpy_y_mask[2], reg_ctrl_cpy_y_mask[1], 2'b11, reg_ctrl_cpy_y_mask[0]};
      end
		
    end else if (src_ce && ~io_src_we_b) begin
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
           reg_x_cnt <= reg_ctrl_width;
           reg_y_cnt <= reg_ctrl_height;
           reg_ctrl_src_addr[7:0] <= reg_ctrl_src_x_origin;
           reg_ctrl_src_addr[12:8] <= reg_ctrl_src_y_origin;
           reg_ctrl_config <= i_src_data[0];
           reg_active_prestage <= 1'h1;
         end
       endcase
     end 
  end

endmodule
