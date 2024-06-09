// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design internal header
// See Vdma.h for the primary calling header

#ifndef VERILATED_VDMA___024UNIT_H_
#define VERILATED_VDMA___024UNIT_H_  // guard

#include "verilated.h"

class Vdma__Syms;

class Vdma___024unit final : public VerilatedModule {
  public:

    // INTERNAL VARIABLES
    Vdma__Syms* const vlSymsp;

    // CONSTRUCTORS
    Vdma___024unit(Vdma__Syms* symsp, const char* v__name);
    ~Vdma___024unit();
    VL_UNCOPYABLE(Vdma___024unit);

    // INTERNAL METHODS
    void __Vconfigure(bool first);
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);


#endif  // guard
