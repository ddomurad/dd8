module vga_data_mux (
  input  wire [7:0] data,
  input  wire clk,  sw, en,
  output wire [3:0] o_data
);

reg [7:0] r_data;
assign o_data = en ? (sw ? r_data[7:4] : r_data[3:0]) : 0;

always @ (posedge clk) begin 
  if (sw)
    r_data <= data; //todo: vs '='
end
endmodule
