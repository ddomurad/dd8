// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Symbol table implementation internals

#include "Vdma__Syms.h"
#include "Vdma.h"
#include "Vdma___024root.h"

// FUNCTIONS
Vdma__Syms::~Vdma__Syms()
{
}

Vdma__Syms::Vdma__Syms(VerilatedContext* contextp, const char* namep, Vdma* modelp)
    : VerilatedSyms{contextp}
    // Setup internal state of the Syms class
    , __Vm_modelp{modelp}
    // Setup module instances
    , TOP{this, namep}
{
    // Configure time unit / time precision
    _vm_contextp__->timeunit(-12);
    _vm_contextp__->timeprecision(-12);
    // Setup each module's pointers to their submodules
    // Setup each module's pointer back to symbol table (for public functions)
    TOP.__Vconfigure(true);
}
