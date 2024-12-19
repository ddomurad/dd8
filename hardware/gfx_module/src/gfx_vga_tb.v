`timescale 1ns / 100ps 

`include "tb_defs.v"
`include "gfx_vga.v"
`include "gfx_addr_mux.v"

module GfxVgaTb();
  localparam TCLK = 39.721946;

  reg reg_clk;
  reg reg_ctrl_ce_b, reg_ctrl_ce2, reg_ctrl_w_b;
  reg [1:0] reg_ctrl_addr;
  reg [7:0] reg_ctrl_data;
  
  wire [15:0] vaddr;
  wire [15:0] vaddr_mux;
  wire  vaddr15_b;
  wire [7:0] vdata;
  wire [7:0] cdata;
  wire [7:0] pixel_data;
  wire [1:0] palette;
  wire       vga_latch;
  wire       vga_out_b;
  wire free_vbus_b;

  wire hsync, vsync;
  wire enabled_b;

  wire [7:0] latch_cdata;

  wire  vram1_cs;
  wire  vram2_cs;

  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst1 (
    .i_ce_b(vram1_cs),
    .i_re_b(1'b0),
    .i_we_b(1'b1),
    .i_addr(vaddr[14:0]),
    .io_data(vdata)
  );

  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst2 (
    .i_ce_b(vram2_cs),
    .i_re_b(1'b0),
    .i_we_b(1'b1),
    .i_addr(vaddr[14:0]),
    .io_data(vdata)
  );

  DualPortRam #(
    .AddrWidth(10),
    .Delay(35),
    .InitFile("./ram/color_palette_ram")
  ) paletteRamInst (
    .i_ce_l_b(enabled_b),
    .i_rw_l_b(1'b1),
    .i_oe_l_b(1'b0),
    .i_addr_l({palette, pixel_data}),
    // .i_addr_l({palette, vdata}),
    .io_data_l(cdata),

    .i_ce_r_b(1'b1),
    .i_rw_r_b(1'b1),
    .i_oe_r_b(1'b1),
    .i_addr_r(10'h000)
  );

  DLatch latch1Inst (
    .i_clk(vga_latch),
    .i_oe_b(vga_out_b),
    .i_in(cdata),
    .o_out(latch_cdata)
  );

  pullup(vram1_cs);
  pullup(vram2_cs);

  // pulldown(latch_cdata[0]);
  // pulldown(latch_cdata[1]);
  // pulldown(latch_cdata[2]);
  // pulldown(latch_cdata[3]);
  // pulldown(latch_cdata[4]);
  // pulldown(latch_cdata[5]);
  // pulldown(latch_cdata[6]);
  // pulldown(latch_cdata[7]);

  assign vram1_cs = vaddr[15];
  assign vram2_cs = vaddr15_b;


  GfxVga gfxVgaInst (
    .i_clk(reg_clk),

    .i_ctrl_ce_b(reg_ctrl_ce_b),
    .i_ctrl_ce2(reg_ctrl_ce2),
    .i_ctrl_w_b(reg_ctrl_w_b),
    .i_ctrl_addr(reg_ctrl_addr),
    .i_ctrl_data(reg_ctrl_data),

    .o_vaddr(vaddr),
    .i_vdata(vdata),
    .o_pixel_data(pixel_data),
    .o_palette(palette),
    .o_vga_latch(vga_latch),
    .o_vga_out_b(vga_out_b),
    .o_hsync(hsync),
    .o_vsync(vsync),
    .o_enabled_b(enabled_b),
    .o_free_vbus_b(free_vbus_b)
  );

  GfxAddrMux gfxAddrMuxInst0(
   .i_addr_sel(1'b0),
   .i_dma_addr({8'ha}),
   .i_vga_addr(vaddr[7:0]),

   .o_addr(vaddr_mux[7:0])
  );

  GfxAddrMux gfxAddrMuxInst1(
   .i_addr_sel(1'b0),
   .i_dma_addr({8'ha}),
   .i_vga_addr(vaddr[15:8]),

   .o_addr(vaddr_mux[15:8]),
   .o_addr7_b(vaddr15_b)
  );

  always begin 
    reg_clk = 1'h1;
    #(TCLK/2);
    reg_clk = 1'h0;
    #(TCLK/2);
  end

  initial begin 
    $dumpfile("./out/gfx_vga_tb.vcd");
    $dumpvars(0, GfxVgaTb);

    reg_ctrl_ce_b = 1'b0;
    reg_ctrl_ce2 = 1'b1;
    reg_ctrl_w_b = 1'b1;
    reg_ctrl_addr = 2'b00;
    reg_ctrl_data = 8'h01;
    #TCLK;
    reg_ctrl_w_b = 1'b0;
    #TCLK;
    reg_ctrl_w_b = 1'b1;


    // #TCLK;
    // reg_ctrl_addr = 2'b01;
    // reg_ctrl_data = 8'h01;
    // #TCLK;
    // reg_ctrl_w_b = 1'b0;
    // #TCLK;
    // reg_ctrl_w_b = 1'b1;

    // #(TCLK*200);
    // reg_ctrl_data = 8'h00;
    // #TCLK;
    // reg_ctrl_w_b = 1'b0;
    // #TCLK;
    // reg_ctrl_w_b = 1'b1;

    // #(TCLK*50);
    // reg_ctrl_data = 8'h01;
    // #TCLK;
    // reg_ctrl_w_b = 1'b0;
    // #TCLK;
    // reg_ctrl_w_b = 1'b1;


    #(TCLK*4000);
    $finish;
  end
endmodule
