module dma(
  input wire clk,
  input wire [4:0] reg_sel, // todo:
  input wire [7:0] reg_data,
  inout wire reg_we,
  output wire busy,

  inout wire [16:0] op_addr,
  inout wire [7:0] op_data,
  inout wire op_re,
  inout wire op_we,
  input wire vga_blank
);

reg [7:0] src_offset_0;
reg [7:0] src_offset_1;
reg src_offset_2;
reg [7:0] dst_offset_0;
reg [7:0] dst_offset_1;
reg dst_offset_2;
reg [7:0] row_width_0;
reg [7:0] cols_number_0;
reg cols_number_1;
reg dir;
reg status;
reg sub_state;
reg next_sub_state;

wire [16:0] op_addr_int; 
wire [7:0] op_data_int;

assign busy = reg_we & status;
assign op_addr = status & vga_blank ? op_addr_int : 17'bz;
assign op_data = sub_state & vga_blank ? op_data_int : 17'bz;

always @ (posedge clk) begin
  if (reg_we & ~status & vga_blank) begin 
    case (reg_sel)
      4'h0 : src_offset_0 = reg_data;
      4'h1 : src_offset_1 = reg_data;
      4'h2 : src_offset_2 = reg_data[0];
      4'h3 : dst_offset_0 = reg_data;
      4'h4 : dst_offset_1 = reg_data;
      4'h5 : dst_offset_2 = reg_data[0];
      4'h6 : row_width_0 = reg_data;
      4'h7 : cols_number_0 = reg_data;
      4'h8 : cols_number_1 = reg_data[0];
      4'h9 : begin 
        dir = reg_data[0]; 
        status = reg_data[1]; 
        // cpy_coutner = 16'h0;
        // sub_state = 0;
      end 
    endcase 
  end
end

reg [7:0] row_cnt;
reg [8:0] col_cnt;
reg [7:0] tmp_reg;
wire [7:0] row_addr;
wire [8:0] col_addr;
wire [8:0] cols_number = {cols_number_1, cols_number_0};

assign op_re = status & vga_blank & ~sub_state;
assign op_we = status & vga_blank & sub_state;
assign row_addr = src_offset_0 + row_cnt;
assign col_addr = {src_offset_2, src_offset_1} + col_cnt;
assign op_addr_int =  {col_addr, row_addr};
assign op_data = sub_state ? tmp_reg : 8'bz;

always @ (posedge clk) begin
  if (status & vga_blank) begin 
    if ( sub_state ) begin  // write data
    end else begin // read data
      tmp_reg = op_data;
    end
  end
end

always @ (negedge clk) begin
  if (status & vga_blank) begin 
    sub_state = next_sub_state;
    if ( sub_state ) begin  // write data
      row_cnt = row_cnt + 1;
      if (row_cnt >= row_width_0) begin 
        row_cnt = 0;
        col_cnt = col_cnt + 1;
      end
      if (col_cnt >= cols_number) begin 
        // status = 0;
      end
    end else begin // read data
    end

    next_sub_state = ~next_sub_state & status;
  end
end

endmodule

