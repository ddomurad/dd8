module dma(
  // CPU side interface 
  input wire reg_clk,
  input wire [3:0] reg_addr,
  input wire [7:0] reg_data, 
  input wire reg_cs, 
  input wire reg_we,
  output wire reg_wait,

  // video side interface
  input wire v_clk,
  inout wire [16:0] v_addr,
  inout wire [7:0] v_data,
  output wire v_re,
  output wire v_we,
  input wire v_blank,

  input wire [7:0] row_cnt,
  input wire [8:0] col_cnt,
  input wire cnt_done
);

localparam REG_SRCOFF_0 = 4'h0,
           REG_SRCOFF_1 = 4'h1,
           REG_SRCOFF_2 = 4'h2,
           REG_DSTOFF_0 = 4'h3,
           REG_DSTOFF_1 = 4'h4,
           REG_DSTOFF_2 = 4'h5,
           REG_FILL = 4'h9,
           REG_CTRL = 4'ha
           ;

// setup registers
reg [16:0] _src_offset;
reg [16:0] _dst_offset;
reg [7:0] _fill_color;
reg _fill_reg;
reg _run_reg;

// temp regis
reg _sub_state;
reg [7:0] _src_data;

// internal wires
wire [16:0] _v_addr;
// wire [7:0] _v_data;
wire [7:0] _src_row_addr;
wire [8:0] _src_col_addr;
wire [16:0] _v_src_addr;
wire [16:0] _v_dst_addr;

assign reg_wait = reg_cs & _run_reg;
assign v_addr = _run_reg & v_blank ? _v_addr : 17'bz;
assign v_data = _run_reg & _sub_state & v_blank ? _src_data : 8'bz;

assign v_re = _run_reg & v_blank & ~_sub_state; //todo: invert ?
assign v_we = _run_reg & v_blank & _sub_state;

assign _src_row_addr = _src_offset[7:0] + row_cnt;
assign _src_col_addr = _src_offset[16:8] + col_cnt;
assign _v_src_addr = { _src_col_addr, _src_row_addr };
assign _v_dst_addr = _v_src_addr + _dst_offset;
assign _v_addr = _sub_state ? _v_src_addr : _v_dst_addr;

always @ (posedge (reg_clk | v_clk)) begin 
  if (reg_cs && reg_we && reg_addr == REG_CTRL) begin
    _run_reg = reg_data[1];
  end else begin 
    if(cnt_done) begin // todo: maybe replace with someting like _done_wire ?
      _run_reg = 0;
    end
  end
end

always @ (posedge reg_clk) begin 
  if (reg_cs & reg_we) begin
    case (reg_addr)
      REG_SRCOFF_0: _src_offset[7:0] <= reg_data;
      REG_SRCOFF_1: _src_offset[15:8] <= reg_data;
      REG_SRCOFF_2: _src_offset[16] <= reg_data[0];

      REG_DSTOFF_0: _dst_offset[7:0] <= reg_data;
      REG_DSTOFF_1: _dst_offset[15:8] <= reg_data;
      REG_DSTOFF_2: _dst_offset[16] <= reg_data[0];

      REG_FILL: _fill_color <= reg_data;
      REG_CTRL: begin 
        _fill_reg <= reg_data[0];
      end
      default: ;
    endcase
  end
end

always @ (posedge v_clk) begin 
  if (_run_reg & v_blank) begin
    if(_sub_state) begin 
      _src_data = v_data;
    end
    _sub_state = ~_sub_state;
  end else if (~_run_reg) begin
    _sub_state = 0;
  end
end

endmodule
