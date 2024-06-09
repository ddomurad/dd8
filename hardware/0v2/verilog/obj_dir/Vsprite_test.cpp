// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Model implementation (design independent parts)

#include "Vsprite_test.h"
#include "Vsprite_test__Syms.h"

//============================================================
// Constructors

Vsprite_test::Vsprite_test(VerilatedContext* _vcontextp__, const char* _vcname__)
    : VerilatedModel{*_vcontextp__}
    , vlSymsp{new Vsprite_test__Syms(contextp(), _vcname__, this)}
    , clk{vlSymsp->TOP.clk}
    , s_data{vlSymsp->TOP.s_data}
    , s_re{vlSymsp->TOP.s_re}
    , s_rs{vlSymsp->TOP.s_rs}
    , r_we{vlSymsp->TOP.r_we}
    , s_addr{vlSymsp->TOP.s_addr}
    , line_n{vlSymsp->TOP.line_n}
    , r_addr{vlSymsp->TOP.r_addr}
    , rootp{&(vlSymsp->TOP)}
{
    // Register model with the context
    contextp()->addModel(this);
}

Vsprite_test::Vsprite_test(const char* _vcname__)
    : Vsprite_test(Verilated::threadContextp(), _vcname__)
{
}

//============================================================
// Destructor

Vsprite_test::~Vsprite_test() {
    delete vlSymsp;
}

//============================================================
// Evaluation function

#ifdef VL_DEBUG
void Vsprite_test___024root___eval_debug_assertions(Vsprite_test___024root* vlSelf);
#endif  // VL_DEBUG
void Vsprite_test___024root___eval_static(Vsprite_test___024root* vlSelf);
void Vsprite_test___024root___eval_initial(Vsprite_test___024root* vlSelf);
void Vsprite_test___024root___eval_settle(Vsprite_test___024root* vlSelf);
void Vsprite_test___024root___eval(Vsprite_test___024root* vlSelf);

void Vsprite_test::eval_step() {
    VL_DEBUG_IF(VL_DBG_MSGF("+++++TOP Evaluate Vsprite_test::eval_step\n"); );
#ifdef VL_DEBUG
    // Debug assertions
    Vsprite_test___024root___eval_debug_assertions(&(vlSymsp->TOP));
#endif  // VL_DEBUG
    if (VL_UNLIKELY(!vlSymsp->__Vm_didInit)) {
        vlSymsp->__Vm_didInit = true;
        VL_DEBUG_IF(VL_DBG_MSGF("+ Initial\n"););
        Vsprite_test___024root___eval_static(&(vlSymsp->TOP));
        Vsprite_test___024root___eval_initial(&(vlSymsp->TOP));
        Vsprite_test___024root___eval_settle(&(vlSymsp->TOP));
    }
    // MTask 0 start
    VL_DEBUG_IF(VL_DBG_MSGF("MTask0 starting\n"););
    Verilated::mtaskId(0);
    VL_DEBUG_IF(VL_DBG_MSGF("+ Eval\n"););
    Vsprite_test___024root___eval(&(vlSymsp->TOP));
    // Evaluate cleanup
    Verilated::endOfThreadMTask(vlSymsp->__Vm_evalMsgQp);
    Verilated::endOfEval(vlSymsp->__Vm_evalMsgQp);
}

//============================================================
// Events and timing
bool Vsprite_test::eventsPending() { return false; }

uint64_t Vsprite_test::nextTimeSlot() {
    VL_FATAL_MT(__FILE__, __LINE__, "", "%Error: No delays in the design");
    return 0;
}

//============================================================
// Utilities

const char* Vsprite_test::name() const {
    return vlSymsp->name();
}

//============================================================
// Invoke final blocks

void Vsprite_test___024root___eval_final(Vsprite_test___024root* vlSelf);

VL_ATTR_COLD void Vsprite_test::final() {
    Vsprite_test___024root___eval_final(&(vlSymsp->TOP));
}

//============================================================
// Implementations of abstract methods from VerilatedModel

const char* Vsprite_test::hierName() const { return vlSymsp->name(); }
const char* Vsprite_test::modelName() const { return "Vsprite_test"; }
unsigned Vsprite_test::threads() const { return 1; }
