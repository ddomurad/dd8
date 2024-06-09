// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vvga_sync.h for the primary calling header

#include "verilated.h"

#include "Vvga_sync___024root.h"

VL_INLINE_OPT void Vvga_sync___024root___ico_sequent__TOP__0(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___ico_sequent__TOP__0\n"); );
    // Body
    vlSelf->vga_sync__DOT__clk_int = ((IData)(vlSelf->clk) 
                                      & (IData)(vlSelf->en));
}

void Vvga_sync___024root___eval_ico(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_ico\n"); );
    // Body
    if (vlSelf->__VicoTriggered.at(0U)) {
        Vvga_sync___024root___ico_sequent__TOP__0(vlSelf);
    }
}

void Vvga_sync___024root___eval_act(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_act\n"); );
}

VL_INLINE_OPT void Vvga_sync___024root___nba_sequent__TOP__0(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___nba_sequent__TOP__0\n"); );
    // Body
    vlSelf->h_cnt = ((IData)(vlSelf->vga_sync__DOT__cnt_h_rst)
                      ? 0U : (0x3ffU & ((IData)(1U) 
                                        + (IData)(vlSelf->h_cnt))));
    vlSelf->h_sync = ((0x28fU < (IData)(vlSelf->h_cnt)) 
                      & (0x2f0U > (IData)(vlSelf->h_cnt)));
    vlSelf->vga_sync__DOT__cnt_h_rst = (0x320U == (IData)(vlSelf->h_cnt));
}

VL_INLINE_OPT void Vvga_sync___024root___nba_sequent__TOP__1(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___nba_sequent__TOP__1\n"); );
    // Body
    vlSelf->v_cnt = ((IData)(vlSelf->vga_sync__DOT__cnt_v_rst)
                      ? 0U : (0x3ffU & ((IData)(1U) 
                                        + (IData)(vlSelf->v_cnt))));
    vlSelf->v_sync = ((0x1e9U < (IData)(vlSelf->v_cnt)) 
                      & (0x1ecU > (IData)(vlSelf->v_cnt)));
    vlSelf->vga_sync__DOT__cnt_v_rst = (0x20dU == (IData)(vlSelf->v_cnt));
}

VL_INLINE_OPT void Vvga_sync___024root___nba_comb__TOP__0(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___nba_comb__TOP__0\n"); );
    // Body
    vlSelf->d_out = ((0x280U > (IData)(vlSelf->h_cnt)) 
                     & (0x1e0U > (IData)(vlSelf->v_cnt)));
    vlSelf->d_out_b = (1U & (~ (IData)(vlSelf->d_out)));
}

void Vvga_sync___024root___eval_nba(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_nba\n"); );
    // Body
    if (vlSelf->__VnbaTriggered.at(0U)) {
        Vvga_sync___024root___nba_sequent__TOP__0(vlSelf);
    }
    if (vlSelf->__VnbaTriggered.at(1U)) {
        Vvga_sync___024root___nba_sequent__TOP__1(vlSelf);
    }
    if ((vlSelf->__VnbaTriggered.at(0U) | vlSelf->__VnbaTriggered.at(1U))) {
        Vvga_sync___024root___nba_comb__TOP__0(vlSelf);
    }
}

void Vvga_sync___024root___eval_triggers__ico(Vvga_sync___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__ico(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG
void Vvga_sync___024root___eval_triggers__act(Vvga_sync___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__act(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG
#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__nba(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG

void Vvga_sync___024root___eval(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval\n"); );
    // Init
    CData/*0:0*/ __VicoContinue;
    VlTriggerVec<2> __VpreTriggered;
    IData/*31:0*/ __VnbaIterCount;
    CData/*0:0*/ __VnbaContinue;
    // Body
    vlSelf->__VicoIterCount = 0U;
    __VicoContinue = 1U;
    while (__VicoContinue) {
        __VicoContinue = 0U;
        Vvga_sync___024root___eval_triggers__ico(vlSelf);
        if (vlSelf->__VicoTriggered.any()) {
            __VicoContinue = 1U;
            if (VL_UNLIKELY((0x64U < vlSelf->__VicoIterCount))) {
#ifdef VL_DEBUG
                Vvga_sync___024root___dump_triggers__ico(vlSelf);
#endif
                VL_FATAL_MT("vga_sync.v", 10, "", "Input combinational region did not converge.");
            }
            vlSelf->__VicoIterCount = ((IData)(1U) 
                                       + vlSelf->__VicoIterCount);
            Vvga_sync___024root___eval_ico(vlSelf);
        }
    }
    __VnbaIterCount = 0U;
    __VnbaContinue = 1U;
    while (__VnbaContinue) {
        __VnbaContinue = 0U;
        vlSelf->__VnbaTriggered.clear();
        vlSelf->__VactIterCount = 0U;
        vlSelf->__VactContinue = 1U;
        while (vlSelf->__VactContinue) {
            vlSelf->__VactContinue = 0U;
            Vvga_sync___024root___eval_triggers__act(vlSelf);
            if (vlSelf->__VactTriggered.any()) {
                vlSelf->__VactContinue = 1U;
                if (VL_UNLIKELY((0x64U < vlSelf->__VactIterCount))) {
#ifdef VL_DEBUG
                    Vvga_sync___024root___dump_triggers__act(vlSelf);
#endif
                    VL_FATAL_MT("vga_sync.v", 10, "", "Active region did not converge.");
                }
                vlSelf->__VactIterCount = ((IData)(1U) 
                                           + vlSelf->__VactIterCount);
                __VpreTriggered.andNot(vlSelf->__VactTriggered, vlSelf->__VnbaTriggered);
                vlSelf->__VnbaTriggered.set(vlSelf->__VactTriggered);
                Vvga_sync___024root___eval_act(vlSelf);
            }
        }
        if (vlSelf->__VnbaTriggered.any()) {
            __VnbaContinue = 1U;
            if (VL_UNLIKELY((0x64U < __VnbaIterCount))) {
#ifdef VL_DEBUG
                Vvga_sync___024root___dump_triggers__nba(vlSelf);
#endif
                VL_FATAL_MT("vga_sync.v", 10, "", "NBA region did not converge.");
            }
            __VnbaIterCount = ((IData)(1U) + __VnbaIterCount);
            Vvga_sync___024root___eval_nba(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
void Vvga_sync___024root___eval_debug_assertions(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_debug_assertions\n"); );
    // Body
    if (VL_UNLIKELY((vlSelf->clk & 0xfeU))) {
        Verilated::overWidthError("clk");}
    if (VL_UNLIKELY((vlSelf->en & 0xfeU))) {
        Verilated::overWidthError("en");}
}
#endif  // VL_DEBUG
