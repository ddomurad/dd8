// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table internal header
//
// Internal details; most calling programs do not need this header,
// unless using verilator public meta comments.

#ifndef VERILATED_VDMA__SYMS_H_
#define VERILATED_VDMA__SYMS_H_  // guard

#include "verilated.h"

// INCLUDE MODEL CLASS

#include "Vdma.h"

// INCLUDE MODULE CLASSES
#include "Vdma___024root.h"

// SYMS CLASS (contains all model state)
class Vdma__Syms final : public VerilatedSyms {
  public:
    // INTERNAL STATE
    Vdma* const __Vm_modelp;
    bool __Vm_didInit = false;

    // MODULE INSTANCE STATE
    Vdma___024root                 TOP;

    // CONSTRUCTORS
    Vdma__Syms(VerilatedContext* contextp, const char* namep, Vdma* modelp);
    ~Vdma__Syms();

    // METHODS
    const char* name() { return TOP.name(); }
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);

#endif  // guard
