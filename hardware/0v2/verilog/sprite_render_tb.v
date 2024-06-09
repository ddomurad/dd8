`timescale 1ns / 1ns
`include "sprite_render.v"

module sprite_ram(
  input wire s_re,
  input wire [14:0] s_addr,
  inout wire [7:0] s_data
);

reg [7:0] mem [0:16383];
assign s_data = s_re ? mem[s_addr] : 8'hz;

initial begin 
  $readmemh("sprite_ram.mem", mem);
end
endmodule

module sprite_render_tb();
reg clk;
reg [8:0] line_n;

reg i_clk;
reg [7:0] i_data;
reg [1:0] i_addr;
reg i_we;
reg i_cs;

wire s_re;
wire [14:0] s_addr;
wire [9:0] r_addr;
wire r_we;
wire [7:0] s_data;

sprite_ram spr_ram(
  .s_addr(s_addr),
  .s_data(s_data),
  .s_re(s_re)
);

dma dut(
  .clk(clk),
  .s_addr(s_addr),
  .s_data(s_data),
  .s_re(s_re),
  .line_n(line_n),
  .r_addr(r_addr),
  .r_we(r_we),
  .i_clk(i_clk),
  .i_data(i_data),
  .i_addr(i_addr),
  .i_we(i_we),
  .i_cs(i_cs)
);


parameter T_CLK2 = 10;

initial begin 
  $dumpfile("./out/sprite_render_tb.vcd");
  $dumpvars(0, sprite_render_tb);

  i_clk = 0;
  i_data = 8'h00;
  i_addr = 2'h0;
  i_we = 0;
  i_cs = 0;

  line_n = 9'h14;
  clk = 0;
  for (integer i = 0; i < 15; i = i + 1) begin 
    #T_CLK2 clk = 0; #T_CLK2 clk = 1;
  end

  line_n = 9'h15;
  for (integer i = 0; i < 10; i = i + 1) begin 
    #T_CLK2 clk = 0; #T_CLK2 clk = 1;
  end

  line_n = 9'h16;
  for (integer i = 0; i < 15; i = i + 1) begin 
    #T_CLK2 clk = 0; #T_CLK2 clk = 1;
  end

  line_n = 9'h17;
  for (integer i = 0; i < 10; i = i + 1) begin 
    #T_CLK2 clk = 0; #T_CLK2 clk = 1;
  end


  line_n = 9'h14;
  clk = 0;
  for (integer i = 0; i < 5; i = i + 1) begin 
    #T_CLK2 clk = 0; #T_CLK2 clk = 1;
  end

end

endmodule
