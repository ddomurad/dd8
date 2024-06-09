// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vdma.h for the primary calling header

#ifndef VERILATED_VDMA___024ROOT_H_
#define VERILATED_VDMA___024ROOT_H_  // guard

#include "verilated.h"

class Vdma__Syms;

class Vdma___024root final : public VerilatedModule {
  public:

    // DESIGN SPECIFIC STATE
    VL_IN8(reg_clk,0,0);
    VL_IN8(v_clk,0,0);
    VL_IN8(reg_addr,3,0);
    VL_IN8(reg_data,7,0);
    VL_IN8(reg_cs,0,0);
    VL_IN8(reg_we,0,0);
    VL_OUT8(reg_wait,0,0);
    VL_INOUT8(v_data,7,0);
    VL_OUT8(v_re,0,0);
    VL_OUT8(v_we,0,0);
    VL_IN8(v_blank,0,0);
    VL_IN8(row_cnt,7,0);
    VL_IN8(cnt_done,0,0);
    CData/*0:0*/ dma__DOT___run_reg;
    CData/*0:0*/ dma__DOT___sub_state;
    CData/*7:0*/ dma__DOT___src_data;
    CData/*0:0*/ dma__DOT____VdfgTmp_h2394a40c__0;
    CData/*0:0*/ __Vtrigprevexpr_hb954e6fa__0;
    CData/*0:0*/ __Vtrigrprev__TOP__reg_clk;
    CData/*0:0*/ __Vtrigrprev__TOP__v_clk;
    CData/*0:0*/ __VactContinue;
    VL_IN16(col_cnt,8,0);
    VL_INOUT(v_addr,16,0);
    IData/*16:0*/ dma__DOT___src_offset;
    IData/*16:0*/ dma__DOT___dst_offset;
    IData/*16:0*/ dma__DOT___v_src_addr;
    IData/*31:0*/ __VstlIterCount;
    IData/*31:0*/ __VicoIterCount;
    IData/*31:0*/ __VactIterCount;
    VlTriggerVec<1> __VstlTriggered;
    VlTriggerVec<1> __VicoTriggered;
    VlTriggerVec<3> __VactTriggered;
    VlTriggerVec<3> __VnbaTriggered;

    // INTERNAL VARIABLES
    Vdma__Syms* const vlSymsp;

    // CONSTRUCTORS
    Vdma___024root(Vdma__Syms* symsp, const char* v__name);
    ~Vdma___024root();
    VL_UNCOPYABLE(Vdma___024root);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);


#endif  // guard
