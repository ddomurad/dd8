// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vsprite_test.h for the primary calling header

#ifndef VERILATED_VSPRITE_TEST___024ROOT_H_
#define VERILATED_VSPRITE_TEST___024ROOT_H_  // guard

#include "verilated.h"

class Vsprite_test__Syms;

class Vsprite_test___024root final : public VerilatedModule {
  public:

    // DESIGN SPECIFIC STATE
    VL_IN8(clk,0,0);
    VL_IN8(s_data,7,0);
    VL_OUT8(s_re,0,0);
    VL_OUT8(s_rs,0,0);
    VL_OUT8(r_we,0,0);
    CData/*2:0*/ dma__DOT___state;
    CData/*5:0*/ dma__DOT___spr_cnt;
    CData/*2:0*/ dma__DOT___cpy_cnt;
    CData/*0:0*/ dma__DOT___last_rst;
    CData/*7:0*/ dma__DOT___y_diff;
    CData/*0:0*/ dma__DOT___write_state;
    CData/*0:0*/ __Vtrigrprev__TOP__clk;
    CData/*0:0*/ __VactContinue;
    VL_OUT16(s_addr,14,0);
    VL_IN16(line_n,8,0);
    VL_OUT16(r_addr,9,0);
    IData/*31:0*/ dma__DOT___spr_reg;
    IData/*31:0*/ __VstlIterCount;
    IData/*31:0*/ __VicoIterCount;
    IData/*31:0*/ __VactIterCount;
    VlTriggerVec<1> __VstlTriggered;
    VlTriggerVec<1> __VicoTriggered;
    VlTriggerVec<2> __VactTriggered;
    VlTriggerVec<2> __VnbaTriggered;

    // INTERNAL VARIABLES
    Vsprite_test__Syms* const vlSymsp;

    // CONSTRUCTORS
    Vsprite_test___024root(Vsprite_test__Syms* symsp, const char* v__name);
    ~Vsprite_test___024root();
    VL_UNCOPYABLE(Vsprite_test___024root);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);


#endif  // guard
