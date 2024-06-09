// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vdram_rnd.h for the primary calling header

#ifndef VERILATED_VDRAM_RND___024ROOT_H_
#define VERILATED_VDRAM_RND___024ROOT_H_  // guard

#include "verilated.h"

class Vdram_rnd__Syms;

class Vdram_rnd___024root final : public VerilatedModule {
  public:

    // DESIGN SPECIFIC STATE
    VL_IN8(clk,0,0);
    VL_OUT8(h_sync,0,0);
    VL_OUT8(v_sync,0,0);
    VL_OUT8(d_out,0,0);
    VL_OUT8(d_out_b,0,0);
    CData/*0:0*/ __Vtrigrprev__TOP__clk;
    CData/*0:0*/ __VactContinue;
    VL_OUT16(h_cnt,9,0);
    VL_OUT16(v_cnt,9,0);
    IData/*31:0*/ __VstlIterCount;
    IData/*31:0*/ __VactIterCount;
    VlTriggerVec<1> __VstlTriggered;
    VlTriggerVec<1> __VactTriggered;
    VlTriggerVec<1> __VnbaTriggered;

    // INTERNAL VARIABLES
    Vdram_rnd__Syms* const vlSymsp;

    // CONSTRUCTORS
    Vdram_rnd___024root(Vdram_rnd__Syms* symsp, const char* v__name);
    ~Vdram_rnd___024root();
    VL_UNCOPYABLE(Vdram_rnd___024root);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);


#endif  // guard
