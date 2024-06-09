// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table internal header
//
// Internal details; most calling programs do not need this header,
// unless using verilator public meta comments.

#ifndef VERILATED_VSPRITE_RENDER__SYMS_H_
#define VERILATED_VSPRITE_RENDER__SYMS_H_  // guard

#include "verilated.h"

// INCLUDE MODEL CLASS

#include "Vsprite_render.h"

// INCLUDE MODULE CLASSES
#include "Vsprite_render___024root.h"

// SYMS CLASS (contains all model state)
class Vsprite_render__Syms final : public VerilatedSyms {
  public:
    // INTERNAL STATE
    Vsprite_render* const __Vm_modelp;
    bool __Vm_didInit = false;

    // MODULE INSTANCE STATE
    Vsprite_render___024root       TOP;

    // CONSTRUCTORS
    Vsprite_render__Syms(VerilatedContext* contextp, const char* namep, Vsprite_render* modelp);
    ~Vsprite_render__Syms();

    // METHODS
    const char* name() { return TOP.name(); }
} VL_ATTR_ALIGNED(VL_CACHE_LINE_BYTES);

#endif  // guard
