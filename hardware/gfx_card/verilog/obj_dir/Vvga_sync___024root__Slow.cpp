// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vvga_sync.h for the primary calling header

#include "verilated.h"

#include "Vvga_sync__Syms.h"
#include "Vvga_sync___024root.h"

void Vvga_sync___024root___ctor_var_reset(Vvga_sync___024root* vlSelf);

Vvga_sync___024root::Vvga_sync___024root(Vvga_sync__Syms* symsp, const char* v__name)
    : VerilatedModule{v__name}
    , vlSymsp{symsp}
 {
    // Reset structure values
    Vvga_sync___024root___ctor_var_reset(this);
}

void Vvga_sync___024root::__Vconfigure(bool first) {
    if (false && first) {}  // Prevent unused
}

Vvga_sync___024root::~Vvga_sync___024root() {
}
