// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table internal header
//
// Internal details; most calling programs do not need this header,
// unless using verilator public meta comments.

#ifndef VERILATED_VSPRITE_TEST__SYMS_H_
#define VERILATED_VSPRITE_TEST__SYMS_H_  // guard

#include "verilated.h"

// INCLUDE MODEL CLASS

#include "Vsprite_test.h"

// INCLUDE MODULE CLASSES
#include "Vsprite_test___024root.h"

// SYMS CLASS (contains all model state)
class Vsprite_test__Syms final : public VerilatedSyms {
  public:
    // INTERNAL STATE
    Vsprite_test* const __Vm_modelp;
    bool __Vm_didInit = false;

    // MODULE INSTANCE STATE
    Vsprite_test___024root         TOP;

    // CONSTRUCTORS
    Vsprite_test__Syms(VerilatedContext* contextp, const char* namep, Vsprite_test* modelp);
    ~Vsprite_test__Syms();

    // METHODS
    const char* name() { return TOP.name(); }
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);

#endif  // guard
