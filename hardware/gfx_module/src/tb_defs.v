`timescale 1ns / 100ps 

module Ram #(
  parameter AddrWidth = 16,
  parameter InitFile = "",
  parameter Delay = 0
)(
  input wire i_ce_b,
  input wire i_re_b,
  input wire i_we_b,
  input wire [AddrWidth-1:0] i_addr,
  inout wire [7:0] io_data
);
  reg [7:0] reg_mem [0:(1<<AddrWidth)-1];

  assign #Delay io_data = ~i_re_b && ~i_ce_b ? reg_mem[i_addr] : 8'hzz;

  integer i;
  initial begin 
    if (InitFile != "")  begin
      $readmemh(InitFile, reg_mem);
    end else begin 
        for (i = 0; i < ((1<<AddrWidth)-1); i = i + 1) begin
          reg_mem[i] = i ;
        end
    end
  end

  always @ (i_ce_b or  i_we_b or i_re_b or i_addr or io_data) begin 
    if (~i_ce_b && ~i_we_b && i_re_b) begin
      #Delay reg_mem[i_addr] <= io_data;
    end
  end
endmodule

// note: this is a simplified dual port ram, not all functions have been implemented
module DualPortRam #(
  parameter AddrWidth = 16,
  parameter InitFile = "",
  parameter Delay = 0
) (
  input wire i_ce_l_b,
  input wire i_rw_l_b,
  input wire i_oe_l_b,
  input wire [9:0] i_addr_l, 
  inout wire [7:0] io_data_l, 

  input wire i_ce_r_b,
  input wire i_rw_r_b,
  input wire i_oe_r_b,
  input wire [9:0] i_addr_r, 
  inout wire [7:0] io_data_r
);

  reg [7:0] reg_mem [0:(1<<AddrWidth)-1];
  wire [7:0] int_data_l;
  wire [7:0] int_data_r;

  assign #Delay int_data_l = reg_mem[i_addr_l];
  assign #Delay int_data_r = reg_mem[i_addr_r];

  assign io_data_l = ~i_oe_l_b & ~i_ce_l_b & i_rw_l_b ? int_data_l : 8'hz;
  assign io_data_r = ~i_oe_r_b & ~i_ce_r_b & i_rw_r_b ? int_data_r : 8'hz;

  initial begin 
    if (InitFile != "")  begin
      $readmemh(InitFile, reg_mem);
    end 
  end

  always @ (i_ce_l_b or i_rw_l_b or i_addr_l or io_data_l) begin 
    if (~i_ce_l_b && ~i_rw_l_b) begin
      #Delay reg_mem[i_addr_l] <= io_data_l;
    end
  end
  always @ (i_ce_r_b or i_rw_r_b or i_addr_r or io_data_r) begin 
    if (~i_ce_r_b && ~i_rw_r_b) begin
      #Delay reg_mem[i_addr_r] <= io_data_r;
    end
  end
endmodule

module DLatch #(
  parameter Delay = 5
) (
  input wire i_oe_b,
  input wire i_clk,
  input wire [7:0] i_in,

  output wire [7:0] o_out
);

  reg [7:0] reg_q;
  assign #Delay o_out = i_oe_b ? 8'hzz : reg_q;

  always @ (posedge i_clk) begin 
    reg_q <= i_in;
  end

endmodule

module BusDriver #(
  parameter BusWidth = 8,
  parameter Delay = 5
)(
  input i_oe_b,
  input [BusWidth-1:0] i_a,
  output [BusWidth-1:0] o_y
); 
  assign #(Delay) o_y = i_oe_b ? {BusWidth{1'hz}} : i_a;
endmodule

module BidBusDriver #(
  parameter BusWidth = 8,
  parameter Delay = 5
)(
  input i_oe_b,
  input i_dir,
  inout [BusWidth-1:0] io_a,
  inout [BusWidth-1:0] io_b
);

  //i_dir 0: a<-b, 1: a->b
  assign #(Delay) io_a = (i_oe_b || i_dir) ? {BusWidth{1'hz}} : io_b;
  assign #(Delay) io_b = (i_oe_b || !i_dir) ? {BusWidth{1'hz}} : io_a;

endmodule
