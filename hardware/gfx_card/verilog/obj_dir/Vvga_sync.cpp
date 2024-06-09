// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Model implementation (design independent parts)

#include "Vvga_sync.h"
#include "Vvga_sync__Syms.h"

//============================================================
// Constructors

Vvga_sync::Vvga_sync(VerilatedContext* _vcontextp__, const char* _vcname__)
    : VerilatedModel{*_vcontextp__}
    , vlSymsp{new Vvga_sync__Syms(contextp(), _vcname__, this)}
    , clk{vlSymsp->TOP.clk}
    , en{vlSymsp->TOP.en}
    , h_sync{vlSymsp->TOP.h_sync}
    , v_sync{vlSymsp->TOP.v_sync}
    , d_out{vlSymsp->TOP.d_out}
    , d_out_b{vlSymsp->TOP.d_out_b}
    , h_cnt{vlSymsp->TOP.h_cnt}
    , v_cnt{vlSymsp->TOP.v_cnt}
    , rootp{&(vlSymsp->TOP)}
{
    // Register model with the context
    contextp()->addModel(this);
}

Vvga_sync::Vvga_sync(const char* _vcname__)
    : Vvga_sync(Verilated::threadContextp(), _vcname__)
{
}

//============================================================
// Destructor

Vvga_sync::~Vvga_sync() {
    delete vlSymsp;
}

//============================================================
// Evaluation function

#ifdef VL_DEBUG
void Vvga_sync___024root___eval_debug_assertions(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG
void Vvga_sync___024root___eval_static(Vvga_sync___024root* vlSelf);
void Vvga_sync___024root___eval_initial(Vvga_sync___024root* vlSelf);
void Vvga_sync___024root___eval_settle(Vvga_sync___024root* vlSelf);
void Vvga_sync___024root___eval(Vvga_sync___024root* vlSelf);

void Vvga_sync::eval_step() {
    VL_DEBUG_IF(VL_DBG_MSGF("+++++TOP Evaluate Vvga_sync::eval_step\n"); );
#ifdef VL_DEBUG
    // Debug assertions
    Vvga_sync___024root___eval_debug_assertions(&(vlSymsp->TOP));
#endif  // VL_DEBUG
    if (VL_UNLIKELY(!vlSymsp->__Vm_didInit)) {
        vlSymsp->__Vm_didInit = true;
        VL_DEBUG_IF(VL_DBG_MSGF("+ Initial\n"););
        Vvga_sync___024root___eval_static(&(vlSymsp->TOP));
        Vvga_sync___024root___eval_initial(&(vlSymsp->TOP));
        Vvga_sync___024root___eval_settle(&(vlSymsp->TOP));
    }
    // MTask 0 start
    VL_DEBUG_IF(VL_DBG_MSGF("MTask0 starting\n"););
    Verilated::mtaskId(0);
    VL_DEBUG_IF(VL_DBG_MSGF("+ Eval\n"););
    Vvga_sync___024root___eval(&(vlSymsp->TOP));
    // Evaluate cleanup
    Verilated::endOfThreadMTask(vlSymsp->__Vm_evalMsgQp);
    Verilated::endOfEval(vlSymsp->__Vm_evalMsgQp);
}

//============================================================
// Events and timing
bool Vvga_sync::eventsPending() { return false; }

uint64_t Vvga_sync::nextTimeSlot() {
    VL_FATAL_MT(__FILE__, __LINE__, "", "%Error: No delays in the design");
    return 0;
}

//============================================================
// Utilities

const char* Vvga_sync::name() const {
    return vlSymsp->name();
}

//============================================================
// Invoke final blocks

void Vvga_sync___024root___eval_final(Vvga_sync___024root* vlSelf);

VL_ATTR_COLD void Vvga_sync::final() {
    Vvga_sync___024root___eval_final(&(vlSymsp->TOP));
}

//============================================================
// Implementations of abstract methods from VerilatedModel

const char* Vvga_sync::hierName() const { return vlSymsp->name(); }
const char* Vvga_sync::modelName() const { return "Vvga_sync"; }
unsigned Vvga_sync::threads() const { return 1; }
