// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table internal header
//
// Internal details; most calling programs do not need this header,
// unless using verilator public meta comments.

#ifndef VERILATED_VVGA_SYNC__SYMS_H_
#define VERILATED_VVGA_SYNC__SYMS_H_  // guard

#include "verilated.h"

// INCLUDE MODEL CLASS

#include "Vvga_sync.h"

// INCLUDE MODULE CLASSES
#include "Vvga_sync___024root.h"

// SYMS CLASS (contains all model state)
class Vvga_sync__Syms final : public VerilatedSyms {
  public:
    // INTERNAL STATE
    Vvga_sync* const __Vm_modelp;
    bool __Vm_didInit = false;

    // MODULE INSTANCE STATE
    Vvga_sync___024root            TOP;

    // CONSTRUCTORS
    Vvga_sync__Syms(VerilatedContext* contextp, const char* namep, Vvga_sync* modelp);
    ~Vvga_sync__Syms();

    // METHODS
    const char* name() { return TOP.name(); }
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);

#endif  // guard
