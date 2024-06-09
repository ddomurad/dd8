// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vsprite_test.h for the primary calling header

#include "verilated.h"

#include "Vsprite_test__Syms.h"
#include "Vsprite_test___024root.h"

void Vsprite_test___024root___ctor_var_reset(Vsprite_test___024root* vlSelf);

Vsprite_test___024root::Vsprite_test___024root(Vsprite_test__Syms* symsp, const char* v__name)
    : VerilatedModule{v__name}
    , vlSymsp{symsp}
 {
    // Reset structure values
    Vsprite_test___024root___ctor_var_reset(this);
}

void Vsprite_test___024root::__Vconfigure(bool first) {
    if (false && first) {}  // Prevent unused
}

Vsprite_test___024root::~Vsprite_test___024root() {
}
