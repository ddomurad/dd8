// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Model implementation (design independent parts)

#include "Vdram_rnd.h"
#include "Vdram_rnd__Syms.h"

//============================================================
// Constructors

Vdram_rnd::Vdram_rnd(VerilatedContext* _vcontextp__, const char* _vcname__)
    : VerilatedModel{*_vcontextp__}
    , vlSymsp{new Vdram_rnd__Syms(contextp(), _vcname__, this)}
    , clk{vlSymsp->TOP.clk}
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

Vdram_rnd::Vdram_rnd(const char* _vcname__)
    : Vdram_rnd(Verilated::threadContextp(), _vcname__)
{
}

//============================================================
// Destructor

Vdram_rnd::~Vdram_rnd() {
    delete vlSymsp;
}

//============================================================
// Evaluation function

#ifdef VL_DEBUG
void Vdram_rnd___024root___eval_debug_assertions(Vdram_rnd___024root* vlSelf);
#endif  // VL_DEBUG
void Vdram_rnd___024root___eval_static(Vdram_rnd___024root* vlSelf);
void Vdram_rnd___024root___eval_initial(Vdram_rnd___024root* vlSelf);
void Vdram_rnd___024root___eval_settle(Vdram_rnd___024root* vlSelf);
void Vdram_rnd___024root___eval(Vdram_rnd___024root* vlSelf);

void Vdram_rnd::eval_step() {
    VL_DEBUG_IF(VL_DBG_MSGF("+++++TOP Evaluate Vdram_rnd::eval_step\n"); );
#ifdef VL_DEBUG
    // Debug assertions
    Vdram_rnd___024root___eval_debug_assertions(&(vlSymsp->TOP));
#endif  // VL_DEBUG
    if (VL_UNLIKELY(!vlSymsp->__Vm_didInit)) {
        vlSymsp->__Vm_didInit = true;
        VL_DEBUG_IF(VL_DBG_MSGF("+ Initial\n"););
        Vdram_rnd___024root___eval_static(&(vlSymsp->TOP));
        Vdram_rnd___024root___eval_initial(&(vlSymsp->TOP));
        Vdram_rnd___024root___eval_settle(&(vlSymsp->TOP));
    }
    // MTask 0 start
    VL_DEBUG_IF(VL_DBG_MSGF("MTask0 starting\n"););
    Verilated::mtaskId(0);
    VL_DEBUG_IF(VL_DBG_MSGF("+ Eval\n"););
    Vdram_rnd___024root___eval(&(vlSymsp->TOP));
    // Evaluate cleanup
    Verilated::endOfThreadMTask(vlSymsp->__Vm_evalMsgQp);
    Verilated::endOfEval(vlSymsp->__Vm_evalMsgQp);
}

//============================================================
// Events and timing
bool Vdram_rnd::eventsPending() { return false; }

uint64_t Vdram_rnd::nextTimeSlot() {
    VL_FATAL_MT(__FILE__, __LINE__, "", "%Error: No delays in the design");
    return 0;
}

//============================================================
// Utilities

const char* Vdram_rnd::name() const {
    return vlSymsp->name();
}

//============================================================
// Invoke final blocks

void Vdram_rnd___024root___eval_final(Vdram_rnd___024root* vlSelf);

VL_ATTR_COLD void Vdram_rnd::final() {
    Vdram_rnd___024root___eval_final(&(vlSymsp->TOP));
}

//============================================================
// Implementations of abstract methods from VerilatedModel

const char* Vdram_rnd::hierName() const { return vlSymsp->name(); }
const char* Vdram_rnd::modelName() const { return "Vdram_rnd"; }
unsigned Vdram_rnd::threads() const { return 1; }
