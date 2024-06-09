// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vvga_sync.h for the primary calling header

#ifndef VERILATED_VVGA_SYNC___024ROOT_H_
#define VERILATED_VVGA_SYNC___024ROOT_H_  // guard

#include "verilated.h"

class Vvga_sync__Syms;

class Vvga_sync___024root final : public VerilatedModule {
  public:

    // DESIGN SPECIFIC STATE
    CData/*0:0*/ vga_sync__DOT__cnt_h_rst;
    CData/*0:0*/ vga_sync__DOT__clk_int;
    VL_IN8(clk,0,0);
    VL_IN8(en,0,0);
    VL_OUT8(h_sync,0,0);
    VL_OUT8(v_sync,0,0);
    VL_OUT8(d_out,0,0);
    VL_OUT8(d_out_b,0,0);
    CData/*0:0*/ vga_sync__DOT__cnt_v_rst;
    CData/*0:0*/ __Vtrigrprev__TOP__vga_sync__DOT__clk_int;
    CData/*0:0*/ __Vtrigrprev__TOP__vga_sync__DOT__cnt_h_rst;
    CData/*0:0*/ __VactContinue;
    VL_OUT16(h_cnt,9,0);
    VL_OUT16(v_cnt,9,0);
    IData/*31:0*/ __VstlIterCount;
    IData/*31:0*/ __VicoIterCount;
    IData/*31:0*/ __VactIterCount;
    VlTriggerVec<1> __VstlTriggered;
    VlTriggerVec<1> __VicoTriggered;
    VlTriggerVec<2> __VactTriggered;
    VlTriggerVec<2> __VnbaTriggered;

    // INTERNAL VARIABLES
    Vvga_sync__Syms* const vlSymsp;

    // CONSTRUCTORS
    Vvga_sync___024root(Vvga_sync__Syms* symsp, const char* v__name);
    ~Vvga_sync___024root();
    VL_UNCOPYABLE(Vvga_sync___024root);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);


#endif  // guard
