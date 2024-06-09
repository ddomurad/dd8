// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vsprite_render.h for the primary calling header

#include "verilated.h"

#include "Vsprite_render__Syms.h"
#include "Vsprite_render___024root.h"

void Vsprite_render___024root___ctor_var_reset(Vsprite_render___024root* vlSelf);

Vsprite_render___024root::Vsprite_render___024root(Vsprite_render__Syms* symsp, const char* v__name)
    : VerilatedModule{v__name}
    , vlSymsp{symsp}
 {
    // Reset structure values
    Vsprite_render___024root___ctor_var_reset(this);
}

void Vsprite_render___024root::__Vconfigure(bool first) {
    if (false && first) {}  // Prevent unused
}

Vsprite_render___024root::~Vsprite_render___024root() {
}
