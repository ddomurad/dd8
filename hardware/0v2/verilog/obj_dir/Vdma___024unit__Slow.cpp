// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vdma.h for the primary calling header

#include "verilated.h"

#include "Vdma__Syms.h"
#include "Vdma___024unit.h"

void Vdma___024unit___ctor_var_reset(Vdma___024unit* vlSelf);

Vdma___024unit::Vdma___024unit(Vdma__Syms* symsp, const char* v__name)
    : VerilatedModule{v__name}
    , vlSymsp{symsp}
 {
    // Reset structure values
    Vdma___024unit___ctor_var_reset(this);
}

void Vdma___024unit::__Vconfigure(bool first) {
    if (false && first) {}  // Prevent unused
}

Vdma___024unit::~Vdma___024unit() {
}
