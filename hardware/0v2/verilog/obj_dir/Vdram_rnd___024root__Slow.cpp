// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vdram_rnd.h for the primary calling header

#include "verilated.h"

#include "Vdram_rnd__Syms.h"
#include "Vdram_rnd___024root.h"

void Vdram_rnd___024root___ctor_var_reset(Vdram_rnd___024root* vlSelf);

Vdram_rnd___024root::Vdram_rnd___024root(Vdram_rnd__Syms* symsp, const char* v__name)
    : VerilatedModule{v__name}
    , vlSymsp{symsp}
 {
    // Reset structure values
    Vdram_rnd___024root___ctor_var_reset(this);
}

void Vdram_rnd___024root::__Vconfigure(bool first) {
    if (false && first) {}  // Prevent unused
}

Vdram_rnd___024root::~Vdram_rnd___024root() {
}