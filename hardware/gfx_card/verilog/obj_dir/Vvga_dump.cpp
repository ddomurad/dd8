// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Model implementation (design independent parts)

#include "Vvga_dump.h"
#include "Vvga_dump__Syms.h"

//============================================================
// Constructors

Vvga_dump::Vvga_dump(VerilatedContext* _vcontextp__, const char* _vcname__)
    : VerilatedModel{*_vcontextp__}
    , vlSymsp{new Vvga_dump__Syms(contextp(), _vcname__, this)}
    , clk{vlSymsp->TOP.clk}
    , d_out{vlSymsp->TOP.d_out}
    , pix_data{vlSymsp->TOP.pix_data}
    , vga_data_b{vlSymsp->TOP.vga_data_b}
    , rootp{&(vlSymsp->TOP)}
{
    // Register model with the context
    contextp()->addModel(this);
}

Vvga_dump::Vvga_dump(const char* _vcname__)
    : Vvga_dump(Verilated::threadContextp(), _vcname__)
{
}

//============================================================
// Destructor

Vvga_dump::~Vvga_dump() {
    delete vlSymsp;
}

//============================================================
// Evaluation function

#ifdef VL_DEBUG
void Vvga_dump___024root___eval_debug_assertions(Vvga_dump___024root* vlSelf);
#endif  // VL_DEBUG
void Vvga_dump___024root___eval_static(Vvga_dump___024root* vlSelf);
void Vvga_dump___024root___eval_initial(Vvga_dump___024root* vlSelf);
void Vvga_dump___024root___eval_settle(Vvga_dump___024root* vlSelf);
void Vvga_dump___024root___eval(Vvga_dump___024root* vlSelf);

void Vvga_dump::eval_step() {
    VL_DEBUG_IF(VL_DBG_MSGF("+++++TOP Evaluate Vvga_dump::eval_step\n"); );
#ifdef VL_DEBUG
    // Debug assertions
    Vvga_dump___024root___eval_debug_assertions(&(vlSymsp->TOP));
#endif  // VL_DEBUG
    if (VL_UNLIKELY(!vlSymsp->__Vm_didInit)) {
        vlSymsp->__Vm_didInit = true;
        VL_DEBUG_IF(VL_DBG_MSGF("+ Initial\n"););
        Vvga_dump___024root___eval_static(&(vlSymsp->TOP));
        Vvga_dump___024root___eval_initial(&(vlSymsp->TOP));
        Vvga_dump___024root___eval_settle(&(vlSymsp->TOP));
    }
    // MTask 0 start
    VL_DEBUG_IF(VL_DBG_MSGF("MTask0 starting\n"););
    Verilated::mtaskId(0);
    VL_DEBUG_IF(VL_DBG_MSGF("+ Eval\n"););
    Vvga_dump___024root___eval(&(vlSymsp->TOP));
    // Evaluate cleanup
    Verilated::endOfThreadMTask(vlSymsp->__Vm_evalMsgQp);
    Verilated::endOfEval(vlSymsp->__Vm_evalMsgQp);
}

//============================================================
// Events and timing
bool Vvga_dump::eventsPending() { return false; }

uint64_t Vvga_dump::nextTimeSlot() {
    VL_FATAL_MT(__FILE__, __LINE__, "", "%Error: No delays in the design");
    return 0;
}

//============================================================
// Utilities

const char* Vvga_dump::name() const {
    return vlSymsp->name();
}

//============================================================
// Invoke final blocks

void Vvga_dump___024root___eval_final(Vvga_dump___024root* vlSelf);

VL_ATTR_COLD void Vvga_dump::final() {
    Vvga_dump___024root___eval_final(&(vlSymsp->TOP));
}

//============================================================
// Implementations of abstract methods from VerilatedModel

const char* Vvga_dump::hierName() const { return vlSymsp->name(); }
const char* Vvga_dump::modelName() const { return "Vvga_dump"; }
unsigned Vvga_dump::threads() const { return 1; }
