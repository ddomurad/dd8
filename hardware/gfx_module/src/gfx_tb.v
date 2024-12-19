`timescale 1ns / 100ps 

`include "tb_defs.v"
`include "gfx_vga.v"
`include "gfx_addr_mux.v"

module GfxTb();
  localparam TCLK = 39.721946;


  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) cram32Inst0 (
    .i_ce_b(),
    .i_re_b(),
    .i_we_b(),
    .i_addr(),
    .io_data()
  );

  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst0 (
    .i_ce_b(),
    .i_re_b(),
    .i_we_b(),
    .i_addr(),
    .io_data()
  );


  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst1 (
    .i_ce_b(),
    .i_re_b(),
    .i_we_b(),
    .i_addr(),
    .io_data()
  );

  DualPortRam #(
    .AddrWidth(10),
    .Delay(35),
    .InitFile("./ram/color_palette_ram")
  ) paletteRamInst (
    .i_ce_l_b(),
    .i_rw_l_b(),
    .i_oe_l_b(),
    .i_addr_l(),
    .io_data_l(),
    .i_ce_r_b(),
    .i_rw_r_b(),
    .i_oe_r_b(),
    .i_addr_r()
  );


  GfxAddrMux gfxAddrMuxInst0(
   .i_addr_sel(),
   .i_dma_addr(),
   .i_vga_addr(),
   .o_addr()
  );

  GfxAddrMux gfxAddrMuxInst1(
   .i_addr_sel(),
   .i_dma_addr(),
   .i_vga_addr(),
   .o_addr(),
   .o_addr7_b()
  );



  DLatch latch0Inst (
    .i_clk(),
    .i_oe_b(),
    .i_in(),
    .o_out()
  );

  DLatch latch1Inst (
    .i_clk(),
    .i_oe_b(),
    .i_in(),
    .o_out()
  );

// BUS DRIVERS
// VGA 
// DMA 

endmodule
