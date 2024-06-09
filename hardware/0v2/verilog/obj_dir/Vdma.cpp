// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Model implementation (design independent parts)

#include "Vdma.h"
#include "Vdma__Syms.h"

//============================================================
// Constructors

Vdma::Vdma(VerilatedContext* _vcontextp__, const char* _vcname__)
    : VerilatedModel{*_vcontextp__}
    , vlSymsp{new Vdma__Syms(contextp(), _vcname__, this)}
    , reg_clk{vlSymsp->TOP.reg_clk}
    , v_clk{vlSymsp->TOP.v_clk}
    , reg_addr{vlSymsp->TOP.reg_addr}
    , reg_data{vlSymsp->TOP.reg_data}
    , reg_cs{vlSymsp->TOP.reg_cs}
    , reg_we{vlSymsp->TOP.reg_we}
    , reg_wait{vlSymsp->TOP.reg_wait}
    , v_data{vlSymsp->TOP.v_data}
    , v_re{vlSymsp->TOP.v_re}
    , v_we{vlSymsp->TOP.v_we}
    , v_blank{vlSymsp->TOP.v_blank}
    , row_cnt{vlSymsp->TOP.row_cnt}
    , cnt_done{vlSymsp->TOP.cnt_done}
    , col_cnt{vlSymsp->TOP.col_cnt}
    , v_addr{vlSymsp->TOP.v_addr}
    , rootp{&(vlSymsp->TOP)}
{
    // Register model with the context
    contextp()->addModel(this);
}

Vdma::Vdma(const char* _vcname__)
    : Vdma(Verilated::threadContextp(), _vcname__)
{
}

//============================================================
// Destructor

Vdma::~Vdma() {
    delete vlSymsp;
}

//============================================================
// Evaluation function

#ifdef VL_DEBUG
void Vdma___024root___eval_debug_assertions(Vdma___024root* vlSelf);
#endif  // VL_DEBUG
void Vdma___024root___eval_static(Vdma___024root* vlSelf);
void Vdma___024root___eval_initial(Vdma___024root* vlSelf);
void Vdma___024root___eval_settle(Vdma___024root* vlSelf);
void Vdma___024root___eval(Vdma___024root* vlSelf);

void Vdma::eval_step() {
    VL_DEBUG_IF(VL_DBG_MSGF("+++++TOP Evaluate Vdma::eval_step\n"); );
#ifdef VL_DEBUG
    // Debug assertions
    Vdma___024root___eval_debug_assertions(&(vlSymsp->TOP));
#endif  // VL_DEBUG
    if (VL_UNLIKELY(!vlSymsp->__Vm_didInit)) {
        vlSymsp->__Vm_didInit = true;
        VL_DEBUG_IF(VL_DBG_MSGF("+ Initial\n"););
        Vdma___024root___eval_static(&(vlSymsp->TOP));
        Vdma___024root___eval_initial(&(vlSymsp->TOP));
        Vdma___024root___eval_settle(&(vlSymsp->TOP));
    }
    // MTask 0 start
    VL_DEBUG_IF(VL_DBG_MSGF("MTask0 starting\n"););
    Verilated::mtaskId(0);
    VL_DEBUG_IF(VL_DBG_MSGF("+ Eval\n"););
    Vdma___024root___eval(&(vlSymsp->TOP));
    // Evaluate cleanup
    Verilated::endOfThreadMTask(vlSymsp->__Vm_evalMsgQp);
    Verilated::endOfEval(vlSymsp->__Vm_evalMsgQp);
}

//============================================================
// Events and timing
bool Vdma::eventsPending() { return false; }

uint64_t Vdma::nextTimeSlot() {
    VL_FATAL_MT(__FILE__, __LINE__, "", "%Error: No delays in the design");
    return 0;
}

//============================================================
// Utilities

const char* Vdma::name() const {
    return vlSymsp->name();
}

//============================================================
// Invoke final blocks

void Vdma___024root___eval_final(Vdma___024root* vlSelf);

VL_ATTR_COLD void Vdma::final() {
    Vdma___024root___eval_final(&(vlSymsp->TOP));
}

//============================================================
// Implementations of abstract methods from VerilatedModel

const char* Vdma::hierName() const { return vlSymsp->name(); }
const char* Vdma::modelName() const { return "Vdma"; }
unsigned Vdma::threads() const { return 1; }
