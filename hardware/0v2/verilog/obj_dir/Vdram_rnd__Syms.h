// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table internal header
//
// Internal details; most calling programs do not need this header,
// unless using verilator public meta comments.

#ifndef VERILATED_VDRAM_RND__SYMS_H_
#define VERILATED_VDRAM_RND__SYMS_H_  // guard

#include "verilated.h"

// INCLUDE MODEL CLASS

#include "Vdram_rnd.h"

// INCLUDE MODULE CLASSES
#include "Vdram_rnd___024root.h"

// SYMS CLASS (contains all model state)
class Vdram_rnd__Syms final : public VerilatedSyms {
  public:
    // INTERNAL STATE
    Vdram_rnd* const __Vm_modelp;
    bool __Vm_didInit = false;

    // MODULE INSTANCE STATE
    Vdram_rnd___024root            TOP;

    // CONSTRUCTORS
    Vdram_rnd__Syms(VerilatedContext* contextp, const char* namep, Vdram_rnd* modelp);
    ~Vdram_rnd__Syms();

    // METHODS
    const char* name() { return TOP.name(); }
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);

#endif  // guard
