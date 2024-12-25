/*
* Simple 8 bit multiplexer for selecting between two addresses.
*/

module GfxAddrMux(
  input i_addr_sel,       // 0 = VGA, 1 = DMA
  input [7:0] i_dma_addr, // DMA address
  input [7:0] i_vga_addr, // VGA address

  output [7:0] o_addr,    // Output address
  output o_addr7_b        // Output address bit 7 inverted
);

assign #(10) o_addr = i_addr_sel ? i_dma_addr : i_vga_addr;
assign #(10) o_addr7_b = i_addr_sel ? ~i_dma_addr[7]  : ~i_vga_addr[7];

endmodule
