`timescale 1ns / 100ps 

`include "tb_defs.v"
`include "gfx_vga.v"
`include "gfx_addr_mux.v"

module GfxTb();
  localparam T_VCLK = 39.721946;
  localparam T_UCLK = T_VCLK*2;

  localparam LOW = 1'b0;
  localparam HIGH = 1'b1;
  localparam HZ = 1'bz;

  reg reg_clk2;
  wire clk;
  assign #(10) clk = reg_clk2;
  

  // cbus - CPU bus, outside bus drivers
  // rbus - internal ram bus, between bus drivers and DMA 
  // vbus - VGA bus, between DMA and VGA modules

  // CPU BUS 
  reg reg_cbus_vga_ce_b;
  reg reg_cbus_dram_ce_b;
  reg reg_cbus_dma_ce_b;
  reg reg_cbus_plt_ce_b;
  reg reg_cbus_re_b;
  reg reg_cbus_we_b;
  reg [12:0] reg_cbus_addr;
  reg [1:0] reg_cbus_bank_sel;
  reg [7:0] reg_cbus_data;

  // INTERNAL DATA RAM
  wire dram_ce_b;
  wire dram_re_b;
  wire dram_we_b;
  wire [14:0] dram_addr;
  wire [7:0] dram_data;

  // DMA 
  wire dma_active;
  wire [12:0] dma_vram_addr; //todo: 12 ?
  wire dma_vaddr_sel;
  
  // VGA
  wire [15:0] vga_vram_addr;
  wire [7:0] vga_pixel_data;
  wire [1:0] vga_palette_data;
  wire vga_latch;
  wire vga_out_b;
  wire vga_hsync;
  wire vga_vsync;
  wire vga_enabled_b;
  wire vga_free_vbus_b;

  // VRAM
  wire vram0_cs_b;
  wire vram1_cs_b; // assigned by the addr mux 
  assign vram0_cs_b = vram_addr[15];
  wire [15:0] vram_addr;
  wire [7:0] vram_data;

  // PALETTE RAM
  wire [7:0] plt_data;

  // VGA SIGNAL 
  wire vga_signal_hsync;
  wire vga_signal_vsync;
  wire [7:0] vga_signal_rgb;


  wire [3:0] bd_int_ctrl;
  assign dram_ce_b = bd_int_ctrl[0];
  //todo: DMA CS 
  assign dram_re_b = bd_int_ctrl[2];
  assign dram_we_b = bd_int_ctrl[3];
  BussDriver #(
    .BusWidth(4)
  )bdCtrlInst (
    .i_oe_b(dma_active),
    .i_a({reg_cbus_we_b, reg_cbus_re_b, reg_cbus_dma_ce_b, reg_cbus_dram_ce_b}),
    .o_y(bd_int_ctrl)
  );

  BussDriver #(
    .BusWidth(14)
  ) bdAddrInst (
    .i_oe_b(dma_active),
    .i_a({reg_cbus_bank_sel, reg_cbus_addr}),
    .o_y(dram_addr)
  );

  BidBussDriver#(
    .BusWidth(8)
  ) bdDataInst (
    .i_oe_b(dma_active),
    .i_dir(reg_cbus_re_b), //todo: verify
    .io_a(reg_cbus_data),
    .io_y(dram_data)
  );

  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) dram32Inst0 (
    .i_ce_b(dram_ce_b),
    .i_re_b(dram_re_b),
    .i_we_b(dram_we_b),
    .i_addr(dram_addr),
    .io_data(dram_data)
  );

  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst0 (
    .i_ce_b(vram0_cs_b),
    .i_re_b(), //todo:
    .i_we_b(), //todo:
    .i_addr(vram_addr[14:0]),
    .io_data(vram_data)
  );


  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst1 (
    .i_ce_b(vram1_cs_b),
    .i_re_b(),//todo:
    .i_we_b(),//todo:
    .i_addr(vram_addr[14:0]),
    .io_data(vram_data)
  );

  DualPortRam #(
    .AddrWidth(10),
    .Delay(35),
    .InitFile("./ram/color_palette_ram")
  ) paletteRamInst (
    .i_ce_l_b(reg_cbus_plt_ce_b),
    .i_oe_l_b(reg_cbus_re_b),
    .i_rw_l_b(reg_cbus_we_b),
    .i_addr_l(reg_cbus_addr[9:0]),
    .io_data_l(reg_cbus_data),
    .i_ce_r_b(vga_enabled_b), //todo: confirm
    .i_oe_r_b(vga_enabled_b),
    .i_rw_r_b(1'b1),
    .i_addr_r({vga_palette_data, vga_pixel_data})
    .io_data_r(plt_data),
  );

  GfxAddrMux gfxAddrMuxInst0(
   .i_addr_sel(dma_vaddr_sel),
   .i_dma_addr(dma_vram_addr[7:0]),
   .i_vga_addr(vga_vram_addr[7:0]),
   .o_addr(vram_addr[7:0])
  );

  GfxAddrMux gfxAddrMuxInst1(
   .i_addr_sel(dma_vaddr_sel),
   .i_dma_addr({000, dma_vram_addr[12:8]}), //todo:
   .i_vga_addr(vga_vram_addr[15:8]),
   .o_addr(vram_addr[15:8]),
   .o_addr7_b(vram1_cs_b)
  );

  GfxVga gfxVgaInst (
    .i_clk(clk),

    .i_ctrl_ce_b(reg_cbus_vga_ce_b), //note: connected directrly to cpu bus
    .i_ctrl_ce2(1'b1), //todo: propabbly we dont need this
    .i_ctrl_w_b(reg_cbus_we_b),
    .i_ctrl_addr(reg_cbus_addr), //todo: size
    .i_ctrl_data(reg_cbus_data),

    .o_vaddr(vga_vram_addr),
    .i_vdata(vram_data),
    .o_pixel_data(vga_pixel_data),
    .o_palette(vga_palette_data),
    .o_vga_latch(vga_latch),
    .o_vga_out_b(vga_out_b),
    .o_hsync(vga_hsync),
    .o_vsync(vga_vsync),
    .o_enabled_b(vga_enabled_b),
    .o_free_vbus_b(vga_free_vbus_b)
  );

  wire [7;0] latch_vga_signals;
  assign vga_signal_hsync = latch_vga_signals[0];
  assign vga_signal_vsync = latch_vga_signals[1];
  DLatch latch0Inst (
    .i_clk(vga_latch),
    .i_oe_b(vga_enabled_b), //todo: confirm
    .i_in({6'h00, vga_vsync, vga_hsync}),
    .o_out(latch_vga_signals)
  );

  DLatch latch1Inst (
    .i_clk(vga_latch),
    .i_oe_b(vga_out_b), //todo: confirm
    .i_in(plt_data),
    .o_out(vga_signal_rgb)
  );

  GfxDma gfxDmaInst (
    .i_clk(clk),
    .i_clk2(reg_clk2),

    .i_src_ce_b(),
    .i_src_ce2_b(),
    .o_src_re_b(),
    .io_src_we_b(),
    // .o_src_ram_ce_b(),
    .io_src_addr(),
    .i_src_data(),

    .o_dst_we_b(),
    .o_dst_addr(),
    .o_dst_data(),

    .i_free_vbus_b(),
    .o_active(dma_active),
    .o_addr_sel(dma_vaddr_sel)
  );

  always begin 
    reg_clk2 = 1'h1;
    #(VTCLK/2);
    reg_clk2 = 1'h0;
    #(VTCLK/2);
  end

  initial begin 

  end

  initial begin 
    $dumpfile("./out/gfx_tb.vcd");
    $dumpvars(0, GfxTb);
  end

endmodule
