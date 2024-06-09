module dma(
  input wire [16:0] a,
  input wire [16:0] b,
  output wire [17:0] y
  output wire [17:0] z
);
assign y = a + b;
assign z = a - b;
endmodule
