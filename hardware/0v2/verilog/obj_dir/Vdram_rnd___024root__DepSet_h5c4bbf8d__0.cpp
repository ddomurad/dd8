// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vdram_rnd.h for the primary calling header

#include "verilated.h"

#include "Vdram_rnd___024root.h"

void Vdram_rnd___024root___eval_act(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_act\n"); );
}

VL_INLINE_OPT void Vdram_rnd___024root___nba_sequent__TOP__0(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___nba_sequent__TOP__0\n"); );
    // Init
    SData/*9:0*/ __Vdly__v_cnt;
    __Vdly__v_cnt = 0;
    SData/*9:0*/ __Vdly__h_cnt;
    __Vdly__h_cnt = 0;
    // Body
    __Vdly__h_cnt = vlSelf->h_cnt;
    __Vdly__v_cnt = vlSelf->v_cnt;
    if ((0x320U == (IData)(vlSelf->h_cnt))) {
        __Vdly__v_cnt = (0x3ffU & ((IData)(1U) + (IData)(vlSelf->v_cnt)));
        __Vdly__h_cnt = 0U;
    } else {
        __Vdly__h_cnt = (0x3ffU & ((IData)(1U) + (IData)(vlSelf->h_cnt)));
    }
    if ((0x20dU == (IData)(vlSelf->v_cnt))) {
        __Vdly__v_cnt = 0U;
    }
    vlSelf->v_cnt = __Vdly__v_cnt;
    vlSelf->h_cnt = __Vdly__h_cnt;
    vlSelf->v_sync = ((0x1eaU > (IData)(vlSelf->v_cnt)) 
                      | (0x1ebU < (IData)(vlSelf->v_cnt)));
    vlSelf->h_sync = ((0x290U > (IData)(vlSelf->h_cnt)) 
                      | (0x2efU < (IData)(vlSelf->h_cnt)));
    vlSelf->d_out = ((0x280U > (IData)(vlSelf->h_cnt)) 
                     & (0x1e0U > (IData)(vlSelf->v_cnt)));
    vlSelf->d_out_b = (1U & (~ (IData)(vlSelf->d_out)));
}

void Vdram_rnd___024root___eval_nba(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_nba\n"); );
    // Body
    if (vlSelf->__VnbaTriggered.at(0U)) {
        Vdram_rnd___024root___nba_sequent__TOP__0(vlSelf);
    }
}

void Vdram_rnd___024root___eval_triggers__act(Vdram_rnd___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vdram_rnd___024root___dump_triggers__act(Vdram_rnd___024root* vlSelf);
#endif  // VL_DEBUG
#ifdef VL_DEBUG
VL_ATTR_COLD void Vdram_rnd___024root___dump_triggers__nba(Vdram_rnd___024root* vlSelf);
#endif  // VL_DEBUG

void Vdram_rnd___024root___eval(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval\n"); );
    // Init
    VlTriggerVec<1> __VpreTriggered;
    IData/*31:0*/ __VnbaIterCount;
    CData/*0:0*/ __VnbaContinue;
    // Body
    __VnbaIterCount = 0U;
    __VnbaContinue = 1U;
    while (__VnbaContinue) {
        __VnbaContinue = 0U;
        vlSelf->__VnbaTriggered.clear();
        vlSelf->__VactIterCount = 0U;
        vlSelf->__VactContinue = 1U;
        while (vlSelf->__VactContinue) {
            vlSelf->__VactContinue = 0U;
            Vdram_rnd___024root___eval_triggers__act(vlSelf);
            if (vlSelf->__VactTriggered.any()) {
                vlSelf->__VactContinue = 1U;
                if (VL_UNLIKELY((0x64U < vlSelf->__VactIterCount))) {
#ifdef VL_DEBUG
                    Vdram_rnd___024root___dump_triggers__act(vlSelf);
#endif
                    VL_FATAL_MT("dram_rnd.v", 1, "", "Active region did not converge.");
                }
                vlSelf->__VactIterCount = ((IData)(1U) 
                                           + vlSelf->__VactIterCount);
                __VpreTriggered.andNot(vlSelf->__VactTriggered, vlSelf->__VnbaTriggered);
                vlSelf->__VnbaTriggered.set(vlSelf->__VactTriggered);
                Vdram_rnd___024root___eval_act(vlSelf);
            }
        }
        if (vlSelf->__VnbaTriggered.any()) {
            __VnbaContinue = 1U;
            if (VL_UNLIKELY((0x64U < __VnbaIterCount))) {
#ifdef VL_DEBUG
                Vdram_rnd___024root___dump_triggers__nba(vlSelf);
#endif
                VL_FATAL_MT("dram_rnd.v", 1, "", "NBA region did not converge.");
            }
            __VnbaIterCount = ((IData)(1U) + __VnbaIterCount);
            Vdram_rnd___024root___eval_nba(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
void Vdram_rnd___024root___eval_debug_assertions(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_debug_assertions\n"); );
    // Body
    if (VL_UNLIKELY((vlSelf->clk & 0xfeU))) {
        Verilated::overWidthError("clk");}
}
#endif  // VL_DEBUG
