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

  output wire [9:0] h_cnt,
  output wire [9:0] v_cnt
);

  wire cnt_h_rst;
  wire cnt_v_rst;
  wire clk_int;

  assign cnt_h_rst = h_cnt == 800;
  assign h_sync = (h_cnt > 655+1) && (h_cnt < 752+1);

  assign cnt_v_rst = v_cnt == 525;
  assign v_sync = (v_cnt > 489) && (v_cnt < 492);
  assign d_out = (h_cnt < 640) && (v_cnt < 480);
  assign d_out_b = !d_out;
  assign clk_int = en ? clk : 0; // clk_int = clk & en;

  vga_cntr cnt0 (.clk (clk_int), .clr(cnt_h_rst), .cnt(h_cnt));
  vga_cntr cnt1 (.clk (cnt_h_rst), .clr(cnt_v_rst), .cnt(v_cnt));

endmodule
