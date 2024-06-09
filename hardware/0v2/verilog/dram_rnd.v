module dma(
  input  wire clk,   
  output wire h_sync, v_sync,
  output wire d_out,  d_out_b,

  output reg [9:0] h_cnt,
  output reg [9:0] v_cnt
);

  assign h_sync = (h_cnt < 656) || (h_cnt > 751);
  assign v_sync = (v_cnt < 490) || (v_cnt > 491);
  assign d_out = (h_cnt < 640) && (v_cnt < 480);
  assign d_out_b = !d_out;


  always @ (posedge clk) begin 
    if (h_cnt == 800) begin
      h_cnt <= 10'h0;
      v_cnt <= v_cnt +1;
    end else begin
      h_cnt <= h_cnt + 1;
    end

    if (v_cnt == 525) begin
      v_cnt <= 10'h0;
    end
  end

endmodule
