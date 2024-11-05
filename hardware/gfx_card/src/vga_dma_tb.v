`timescale 1ns / 100ps 

`include "tb_defs.v"
`include "vga_dma.v"

module VgaDmaTb();
  localparam VTCLK = 39.721946;
  localparam UTCLK = VTCLK*2;

  localparam LOW = 1'b0;
  localparam HIGH = 1'b1;
  localparam HZ = 1'bz;

  reg reg_clk2;
  reg [12:0] reg_addr;
  reg [7:0] reg_data;
  reg reg_vga_active_b;
  reg reg_free_vbus_b;
  reg reg_ram_re_b;
  reg reg_ram_we_b;
  reg reg_vga_ce_b;
  reg reg_ram_ce_b;


  wire clk;
  wire [12:0] cram_addr;
  wire [7:0] cram_data;
  wire vga_ce_b;
  wire cram_ce_b;
  wire cram_oe_b;
  wire cram_we_b;

  wire [15:0] vram_addr;
  wire vram_addr15_b;
  wire [7:0] vram_data;
  wire vram_we_b;

  wire dma_active;


  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15),
    .InitFile("./ram/dma_ram")
  ) cram32Inst (
    .i_ce_b(cram_ce_b),
    .i_re_b(cram_oe_b),
    .i_we_b(cram_we_b),
    .i_addr({2'b0, cram_addr}),
    .io_data(cram_data)
  );


  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst1 (
    .i_ce_b(vram_addr[15]),
    .i_re_b(reg_vga_active_b),
    .i_we_b(vram_we_b),
    .i_addr(vram_addr[14:0]),
    .io_data(vram_data)
  );

  Ram #( // according to AS7C256B datasheet
    .AddrWidth(15),
    .Delay(15)
  ) vram32Inst2 (
    .i_ce_b(vram_addr15_b),
    .i_re_b(reg_vga_active_b),
    .i_we_b(vram_we_b),
    .i_addr(vram_addr[14:0]),
    .io_data(vram_data)
  );

  BussDriver #(
    .BusWidth(13)
  ) bdAddrInst (
    .i_oe_b(dma_active),
    .i_a(reg_addr),
    .o_y(cram_addr)
  );

  BussDriver bdDataInst (
    .i_oe_b(dma_active),
    .i_a(reg_data),
    .o_y(cram_data)
  );

  // BidBussDriver bdDataInst (
  //   .i_oe_b(dma_active),
  //   .i_re_b(reg_ram_re_b),
  //   .io_a(reg_data),
  //   .io_y(cram_data)
  // );

  wire [7:0] ctrlBusInput;
  wire [7:0] ctrlBusOutput;
  BussDriver bdCtrlInst (
    .i_oe_b(dma_active),
    .i_a(ctrlBusInput),
    .o_y(ctrlBusOutput) 
  );
  assign ctrlBusInput = {4'h0, reg_ram_ce_b, reg_ram_we_b, reg_ram_re_b, reg_vga_ce_b};
  assign vga_ce_b = ctrlBusOutput[0];
  assign cram_oe_b = ctrlBusOutput[1];
  assign cram_we_b = ctrlBusOutput[2];
  assign cram_ce_b = ctrlBusOutput[3];

  pulldown(cram_ce_b);
  pullup(cram_oe_b);
  pullup(cram_we_b);
  pullup(vga_ce_b);

  VgaDma vgaDmaInst (
    .i_clk(clk),
    .i_clk2(reg_clk2),

    .i_src_ce_b(vga_ce_b),
    .i_src_ce2_b(cram_addr[3]),
    .o_src_re_b(cram_oe_b),
    .io_src_we_b(cram_we_b),
    // .o_src_ram_ce_b(),
    .io_src_addr(cram_addr),
    .i_src_data(cram_data),

    .o_dst_we_b(vram_we_b),
    .o_dst_addr(vram_addr),
    .o_dst_data(vram_data),

    .i_free_vbus_b(reg_free_vbus_b),
    .o_active(dma_active) 
  );

  always begin 
    reg_clk2 = 1'h1;
    #(VTCLK/2);
    reg_clk2 = 1'h0;
    #(VTCLK/2);
  end

  assign #(10) clk = reg_clk2;

  initial begin 
    reg_vga_ce_b = HIGH;
    reg_ram_ce_b = HIGH;
    reg_ram_we_b = HIGH;
    reg_ram_re_b = HIGH;
    reg_addr = 14'h00;
    reg_data = 8'h00;
    reg_free_vbus_b = LOW;
    reg_vga_active_b = HIGH;
  end

  initial begin 
    $dumpfile("./out/vga_dma_tb.vcd");
    $dumpvars(0, VgaDmaTb);

    #UTCLK;

    //note: src addr low 
    reg_vga_ce_b = LOW; 
    reg_ram_we_b = HIGH; reg_addr = 14'h00; reg_data = 8'h00;
    #UTCLK; reg_ram_we_b = LOW;
    // note: src addr high
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'h01; reg_data = 8'h00;
    #UTCLK; reg_ram_we_b = LOW;
    // note: dst addr low
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'h02; reg_data = 8'h10;
    #UTCLK; reg_ram_we_b = LOW;
    // note: dst addr high
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'h03; reg_data = 8'h10;
    #UTCLK; reg_ram_we_b = LOW;
    // note: width
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'h04; reg_data = 8'h05;
    #UTCLK; reg_ram_we_b = LOW;
    // note: height
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'h05; reg_data = 8'h05;
    #UTCLK; reg_ram_we_b = LOW;
    // note: mask
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'h06; reg_data = 8'hff;
    #UTCLK; reg_ram_we_b = LOW;
    // note: start
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'h07; reg_data = 8'h00;
    #UTCLK; reg_ram_we_b = LOW;
    #UTCLK; reg_ram_we_b = HIGH;  reg_addr = 14'hzz; reg_data = 8'hzz;

    #UTCLK; reg_vga_ce_b = HIGH;


    #(UTCLK*5.64);
    reg_free_vbus_b = HIGH;

    #(UTCLK*2);
    reg_free_vbus_b = LOW;

    #(UTCLK*2000);
    $finish;
  end

endmodule
