// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vvga_dump.h for the primary calling header

#ifndef VERILATED_VVGA_DUMP___024ROOT_H_
#define VERILATED_VVGA_DUMP___024ROOT_H_  // guard

#include "verilated.h"

class Vvga_dump__Syms;

class Vvga_dump___024root final : public VerilatedModule {
  public:

    // DESIGN SPECIFIC STATE
    VL_IN8(clk,0,0);
    VL_IN8(d_out,0,0);
    VL_IN8(pix_data,7,0);
    CData/*0:0*/ __Vtrigrprev__TOP__clk;
    CData/*0:0*/ __VactContinue;
    VL_OUT16(vga_data_b,8,0);
    SData/*8:0*/ vga_dump__DOT___vga_data;
    IData/*31:0*/ __VstlIterCount;
    IData/*31:0*/ __VicoIterCount;
    IData/*31:0*/ __VactIterCount;
    VlTriggerVec<1> __VstlTriggered;
    VlTriggerVec<1> __VicoTriggered;
    VlTriggerVec<1> __VactTriggered;
    VlTriggerVec<1> __VnbaTriggered;

    // INTERNAL VARIABLES
    Vvga_dump__Syms* const vlSymsp;

    // CONSTRUCTORS
    Vvga_dump___024root(Vvga_dump__Syms* symsp, const char* v__name);
    ~Vvga_dump___024root();
    VL_UNCOPYABLE(Vvga_dump___024root);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);


#endif  // guard
