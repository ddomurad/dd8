// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Model implementation (design independent parts)

#include "Vsprite_render.h"
#include "Vsprite_render__Syms.h"

//============================================================
// Constructors

Vsprite_render::Vsprite_render(VerilatedContext* _vcontextp__, const char* _vcname__)
    : VerilatedModel{*_vcontextp__}
    , vlSymsp{new Vsprite_render__Syms(contextp(), _vcname__, this)}
    , clk{vlSymsp->TOP.clk}
    , i_clk{vlSymsp->TOP.i_clk}
    , s_data{vlSymsp->TOP.s_data}
    , s_re{vlSymsp->TOP.s_re}
    , s_we{vlSymsp->TOP.s_we}
    , off_src_line{vlSymsp->TOP.off_src_line}
    , r_we{vlSymsp->TOP.r_we}
    , i_addr{vlSymsp->TOP.i_addr}
    , i_data{vlSymsp->TOP.i_data}
    , i_we{vlSymsp->TOP.i_we}
    , i_cs{vlSymsp->TOP.i_cs}
    , i_wait{vlSymsp->TOP.i_wait}
    , s_addr{vlSymsp->TOP.s_addr}
    , line_n{vlSymsp->TOP.line_n}
    , r_addr{vlSymsp->TOP.r_addr}
    , rootp{&(vlSymsp->TOP)}
{
    // Register model with the context
    contextp()->addModel(this);
}

Vsprite_render::Vsprite_render(const char* _vcname__)
    : Vsprite_render(Verilated::threadContextp(), _vcname__)
{
}

//============================================================
// Destructor

Vsprite_render::~Vsprite_render() {
    delete vlSymsp;
}

//============================================================
// Evaluation function

#ifdef VL_DEBUG
void Vsprite_render___024root___eval_debug_assertions(Vsprite_render___024root* vlSelf);
#endif  // VL_DEBUG
void Vsprite_render___024root___eval_static(Vsprite_render___024root* vlSelf);
void Vsprite_render___024root___eval_initial(Vsprite_render___024root* vlSelf);
void Vsprite_render___024root___eval_settle(Vsprite_render___024root* vlSelf);
void Vsprite_render___024root___eval(Vsprite_render___024root* vlSelf);

void Vsprite_render::eval_step() {
    VL_DEBUG_IF(VL_DBG_MSGF("+++++TOP Evaluate Vsprite_render::eval_step\n"); );
#ifdef VL_DEBUG
    // Debug assertions
    Vsprite_render___024root___eval_debug_assertions(&(vlSymsp->TOP));
#endif  // VL_DEBUG
    if (VL_UNLIKELY(!vlSymsp->__Vm_didInit)) {
        vlSymsp->__Vm_didInit = true;
        VL_DEBUG_IF(VL_DBG_MSGF("+ Initial\n"););
        Vsprite_render___024root___eval_static(&(vlSymsp->TOP));
        Vsprite_render___024root___eval_initial(&(vlSymsp->TOP));
        Vsprite_render___024root___eval_settle(&(vlSymsp->TOP));
    }
    // MTask 0 start
    VL_DEBUG_IF(VL_DBG_MSGF("MTask0 starting\n"););
    Verilated::mtaskId(0);
    VL_DEBUG_IF(VL_DBG_MSGF("+ Eval\n"););
    Vsprite_render___024root___eval(&(vlSymsp->TOP));
    // Evaluate cleanup
    Verilated::endOfThreadMTask(vlSymsp->__Vm_evalMsgQp);
    Verilated::endOfEval(vlSymsp->__Vm_evalMsgQp);
}

//============================================================
// Events and timing
bool Vsprite_render::eventsPending() { return false; }

uint64_t Vsprite_render::nextTimeSlot() {
    VL_FATAL_MT(__FILE__, __LINE__, "", "%Error: No delays in the design");
    return 0;
}

//============================================================
// Utilities

const char* Vsprite_render::name() const {
    return vlSymsp->name();
}

//============================================================
// Invoke final blocks

void Vsprite_render___024root___eval_final(Vsprite_render___024root* vlSelf);

VL_ATTR_COLD void Vsprite_render::final() {
    Vsprite_render___024root___eval_final(&(vlSymsp->TOP));
}

//============================================================
// Implementations of abstract methods from VerilatedModel

const char* Vsprite_render::hierName() const { return vlSymsp->name(); }
const char* Vsprite_render::modelName() const { return "Vsprite_render"; }
unsigned Vsprite_render::threads() const { return 1; }
