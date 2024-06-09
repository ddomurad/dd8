// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table internal header
//
// Internal details; most calling programs do not need this header,
// unless using verilator public meta comments.

#ifndef VERILATED_VVGA_DUMP__SYMS_H_
#define VERILATED_VVGA_DUMP__SYMS_H_  // guard

#include "verilated.h"

// INCLUDE MODEL CLASS

#include "Vvga_dump.h"

// INCLUDE MODULE CLASSES
#include "Vvga_dump___024root.h"

// SYMS CLASS (contains all model state)
class Vvga_dump__Syms final : public VerilatedSyms {
  public:
    // INTERNAL STATE
    Vvga_dump* const __Vm_modelp;
    bool __Vm_didInit = false;

    // MODULE INSTANCE STATE
    Vvga_dump___024root            TOP;

    // CONSTRUCTORS
    Vvga_dump__Syms(VerilatedContext* contextp, const char* namep, Vvga_dump* modelp);
    ~Vvga_dump__Syms();

    // METHODS
    const char* name() { return TOP.name(); }
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);

#endif  // guard
