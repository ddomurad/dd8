module vga_dump(
  input wire d_out,
  input wire clk,

  input wire [7:0] pix_data,
  output wire [8:0] vga_data_b // inverted vga color bits
);

reg [8:0] _vga_data;

assign vga_data_b = d_out ? _vga_data : 9'b111111111;

always @ (negedge clk) begin 
  if (d_out)
    _vga_data <= ~{2'b00, pix_data[6:0]};
  else
    _vga_data <= 9'b111111111;
end
endmodule
