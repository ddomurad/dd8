module dma(
  input wire clk,

  output wire [14:0] s_addr,
  inout wire [7:0] s_data,
  output wire s_re,
  output wire s_we,

  input wire [8:0] line_n,
  input wire off_src_line,

  output wire [9:0] r_addr,
  output wire r_we,

  input wire i_clk,
  input wire [1:0] i_addr,
  input wire [7:0] i_data,
  input wire i_we,
  input wire i_cs,
  output wire i_wait
);

localparam  ST_READ_Y = 3'h0,
            ST_READ_X = 3'h1,
            ST_READ_C = 3'h2,
            ST_READ_I = 3'h3,
            ST_WRITE = 3'h4;

reg [14:0] _i_addr;

reg [2:0] _state;
reg [5:0] _spr_cnt;
reg [31:0] _spr_reg; 

reg [2:0] _cpy_cnt;
reg _last_rst;

wire [7:0] _rnd_line;
wire [7:0] _y_diff;
wire [8:0] _spr_w_addr;
wire _write_state; 
wire _valid_pixel;
wire _line_reset;

assign _line_reset = line_n[0];
assign _rnd_line = line_n[8:1];
assign _y_diff = _rnd_line - _spr_reg[7:0];
assign _write_state = _state == ST_WRITE ? 1'h1 : 1'h0 ;
assign _valid_pixel = s_data == 8'h00 ? 1'h0 : 1'h1; // can be handled externally

assign s_addr = off_src_line ? _i_addr : (_write_state ? { _spr_reg[31:23], _y_diff[2:0], _cpy_cnt[2:0]} : {7'h0, _spr_cnt, _state[1:0]});
assign _spr_w_addr = _spr_reg[16:8] + {6'h0, _cpy_cnt};
assign r_addr = {~line_n[1], _spr_w_addr}; // line_n[1] can be wired externally
assign r_we = _write_state & clk & _valid_pixel ? 1'h1 : 1'h0; //todo: inverted signal ? 
assign i_wait = i_cs && i_addr == 2'h3 && !off_src_line;

assign s_data = off_src_line ? i_data : 8'bz;
assign s_re = !off_src_line; // can be probably removed and replaces by external wire
assign s_we = off_src_line && i_cs && i_we && i_addr == 2'h3;

// note: grey zone....
initial begin 
  _state = 3'b111;
  _spr_cnt = 6'h3f;
  _spr_reg = 32'h0;
  _cpy_cnt = 3'h0;
  _last_rst = 1'h0;
end

always @ (posedge clk) begin 
  case (_state) 
    ST_READ_Y: _spr_reg[7:0] = s_data;
    ST_READ_X: _spr_reg[15:8] = s_data;
    ST_READ_C: _spr_reg[23:16] = s_data;
    ST_READ_I: _spr_reg[31:24] = s_data;
    ST_WRITE: begin 
    end
    default: begin end
  endcase
end

always @ (negedge clk) begin 
  if(_last_rst != _line_reset) begin 
    if(~_line_reset) begin 
      _state = 0;
      _cpy_cnt = 0; 
      _spr_cnt = 0;
    end
    _last_rst = _line_reset;
  end else if(_y_diff > 8'h7) begin 
      _cpy_cnt = 0; 
      _state = 0;
      _spr_cnt = _spr_cnt + 1'h1;
  end else if(_write_state) begin 
    if(_cpy_cnt == 3'h7) begin 
      _cpy_cnt = 0;
      _state = 0;
      _spr_cnt = _spr_cnt + 1'h1;
    end else begin 
      _cpy_cnt = _cpy_cnt + 1'h1;
    end
  end else begin
    _state = _state + 1'h1;
  end
end

always @ (posedge i_clk) begin 
  if(i_cs && i_we && off_src_line) begin
    case (i_addr)
      2'h0: _i_addr[7:0] = i_data;
      2'h1: _i_addr[14:8] = i_data[6:0];
      default: begin end
    endcase 
  end
end
endmodule
