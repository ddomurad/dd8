module vga_cntr(input clk, clr, output reg [9:0] cnt);
always @ (posedge clk) begin 
  if (clr)
    cnt <= 0;
  else
    cnt <= cnt + 1;
end
endmodule

module vga_sync(
  input  wire clk,    en,
  output wire h_sync, v_sync,
  output wire d_out,  d_out_b,

  inout wire [9:0] h_cnt,
  inout wire [9:0] v_cnt
);

  wire cnt_h_rst;
  wire cnt_v_rst;
  wire clk_int;

  wire [9:0] h_cnt_int;
  wire [9:0] v_cnt_int;

  assign cnt_h_rst = h_cnt_int == 800;
  assign h_sync = (h_cnt_int < 592) || (h_cnt_int > 687);
  // assign h_sync = (h_cnt < 656) || (h_cnt > 751);
  assign cnt_v_rst = v_cnt_int == 525;
  assign v_sync = (v_cnt_int < 490) || (v_cnt_int > 491);
  assign d_out = (h_cnt_int < 512) && (v_cnt_int < 480);
  assign d_out_b = !d_out;
  assign clk_int = en ? clk : 0; // clk_int = clk & en;

  vga_cntr cnt0 (.clk (clk_int), .clr(cnt_h_rst), .cnt(h_cnt_int));
  vga_cntr cnt1 (.clk (cnt_h_rst), .clr(cnt_v_rst), .cnt(v_cnt_int));

  assign h_cnt = d_out ? h_cnt_int : 10'bzzzzzzzzzz;
  assign v_cnt = d_out ? v_cnt_int : 10'bzzzzzzzzzz;
endmodule
