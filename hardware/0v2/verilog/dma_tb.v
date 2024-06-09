`timescale 1ns / 1ns
`include "dma.v"

module dma_tb();

reg clk;
reg [3:0] reg_addr;
reg [7:0] reg_data;
reg reg_cs;
reg reg_we;
wire reg_wait;


dma dut(
  .reg_clk(clk),
  .reg_addr(reg_addr),
  .reg_data(reg_data),
  .reg_cs(reg_cs),
  .reg_we(reg_we),
  .reg_wait(reg_wait),

  .v_clk(clk)
);


task write_reg;
  input [3:0] addr;
  input [7:0] data;
begin 
  reg_cs = 1;
  reg_we = 1;
  reg_addr = addr;
  reg_data = data;
  #10 clk = 1;
  #10 clk = 0;
end
endtask

task setup_dma;
  input [16:0] src_off;
  input [16:0] dst_off;
  input [7:0] width;
  input [8:0] height;
  input [7:0] fill;
  input [7:0] ctrl;
begin 
  write_reg(4'h0, src_off[7:0]); // REG_SRCOFF_0
  write_reg(4'h1, src_off[15:8]); // REG_SRCOFF_1
  write_reg(4'h2, src_off[16]); // REG_SRCOFF_2
  write_reg(4'h3, dst_off[7:0]); // REG_DSTOFF_0
  write_reg(4'h4, dst_off[15:8]); // REG_DSTOFF_1
  write_reg(4'h5, dst_off[16]); // REG_DSTOFF_2
  write_reg(4'h6, width); // REG_WIDTH
  write_reg(4'h7, height[7:0]); // REG_HEIGHT_0
  write_reg(4'h8, height[8]); // REG_HEIGHT_1
  write_reg(4'h9, fill);  // REG_FILL
  write_reg(4'ha, ctrl); // REG_CTRL
  reg_cs = 0; 
  reg_we = 0;

end
endtask

initial begin 
  $dumpfile("./out/dma_tb.vcd");
  $dumpvars(0, dma_tb);

  #10 clk = 0;
  setup_dma(17'h01010, 17'h03030, 8'h10, 9'h010, 8'h00, 8'h02);

  #10 clk = 1;
  #10 clk = 0;
  #10 clk = 1;
  #10 clk = 0;
  #10 clk = 1;
  #10 clk = 0;
end

endmodule   
