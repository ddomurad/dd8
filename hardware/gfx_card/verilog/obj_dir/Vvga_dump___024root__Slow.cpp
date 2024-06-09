// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vvga_dump.h for the primary calling header

#include "verilated.h"

#include "Vvga_dump__Syms.h"
#include "Vvga_dump___024root.h"

void Vvga_dump___024root___ctor_var_reset(Vvga_dump___024root* vlSelf);

Vvga_dump___024root::Vvga_dump___024root(Vvga_dump__Syms* symsp, const char* v__name)
    : VerilatedModule{v__name}
    , vlSymsp{symsp}
 {
    // Reset structure values
    Vvga_dump___024root___ctor_var_reset(this);
}

void Vvga_dump___024root::__Vconfigure(bool first) {
    if (false && first) {}  // Prevent unused
}

Vvga_dump___024root::~Vvga_dump___024root() {
}
