module GfxAddrMux(
  input i_addr_sel,
  input [7:0] i_dma_addr,
  input [7:0] i_vga_addr,

  output [7:0] o_addr,
  output o_addr7_b
);

assign o_addr = i_addr_sel ? i_dma_addr : i_vga_addr;
assign o_addr7_b = i_addr_sel ? ~i_dma_addr[7]  : ~i_vga_addr[7];

endmodule
