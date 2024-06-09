// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vdma.h for the primary calling header

#include "verilated.h"

#include "Vdma___024root.h"

VL_INLINE_OPT void Vdma___024root___ico_sequent__TOP__0(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___ico_sequent__TOP__0\n"); );
    // Init
    CData/*0:0*/ dma__DOT____VdfgTmp_hbe67ea25__0;
    dma__DOT____VdfgTmp_hbe67ea25__0 = 0;
    // Body
    vlSelf->reg_wait = ((IData)(vlSelf->dma__DOT___run_reg) 
                        & (IData)(vlSelf->reg_cs));
    dma__DOT____VdfgTmp_hbe67ea25__0 = ((IData)(vlSelf->dma__DOT___run_reg) 
                                        & ((IData)(vlSelf->dma__DOT___sub_state) 
                                           & (IData)(vlSelf->v_blank)));
    vlSelf->dma__DOT___v_src_addr = ((0x1ff00U & ((
                                                   (vlSelf->dma__DOT___src_offset 
                                                    >> 8U) 
                                                   + (IData)(vlSelf->col_cnt)) 
                                                  << 8U)) 
                                     | (0xffU & (vlSelf->dma__DOT___src_offset 
                                                 + (IData)(vlSelf->row_cnt))));
    vlSelf->dma__DOT____VdfgTmp_h2394a40c__0 = ((IData)(vlSelf->dma__DOT___run_reg) 
                                                & (IData)(vlSelf->v_blank));
    vlSelf->v_data = (((IData)(dma__DOT____VdfgTmp_hbe67ea25__0)
                        ? 0xffU : 0U) & (((IData)(dma__DOT____VdfgTmp_hbe67ea25__0)
                                           ? (IData)(vlSelf->dma__DOT___src_data)
                                           : 0U) & 
                                         ((IData)(dma__DOT____VdfgTmp_hbe67ea25__0)
                                           ? 0xffU : 0U)));
    vlSelf->v_re = ((~ (IData)(vlSelf->dma__DOT___sub_state)) 
                    & (IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0));
    vlSelf->v_we = ((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0) 
                    & (IData)(vlSelf->dma__DOT___sub_state));
    vlSelf->v_addr = (((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0)
                        ? 0x1ffffU : 0U) & (((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0)
                                              ? ((IData)(vlSelf->dma__DOT___sub_state)
                                                  ? vlSelf->dma__DOT___v_src_addr
                                                  : 
                                                 (vlSelf->dma__DOT___v_src_addr 
                                                  + vlSelf->dma__DOT___dst_offset))
                                              : 0U) 
                                            & ((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0)
                                                ? 0x1ffffU
                                                : 0U)));
}

void Vdma___024root___eval_ico(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_ico\n"); );
    // Body
    if (vlSelf->__VicoTriggered.at(0U)) {
        Vdma___024root___ico_sequent__TOP__0(vlSelf);
    }
}

void Vdma___024root___eval_act(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_act\n"); );
}

VL_INLINE_OPT void Vdma___024root___nba_sequent__TOP__0(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___nba_sequent__TOP__0\n"); );
    // Body
    if (((IData)(vlSelf->dma__DOT___run_reg) & (IData)(vlSelf->v_blank))) {
        if (vlSelf->dma__DOT___sub_state) {
            vlSelf->dma__DOT___src_data = vlSelf->v_data;
        }
        vlSelf->dma__DOT___sub_state = (1U & (~ (IData)(vlSelf->dma__DOT___sub_state)));
    } else if ((1U & (~ (IData)(vlSelf->dma__DOT___run_reg)))) {
        vlSelf->dma__DOT___sub_state = 0U;
    }
}

VL_INLINE_OPT void Vdma___024root___nba_sequent__TOP__1(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___nba_sequent__TOP__1\n"); );
    // Body
    if (((IData)(vlSelf->reg_cs) & (IData)(vlSelf->reg_we))) {
        if ((1U & (~ ((IData)(vlSelf->reg_addr) >> 3U)))) {
            if ((4U & (IData)(vlSelf->reg_addr))) {
                if ((1U & (~ ((IData)(vlSelf->reg_addr) 
                              >> 1U)))) {
                    vlSelf->dma__DOT___dst_offset = 
                        ((1U & (IData)(vlSelf->reg_addr))
                          ? ((0xffffU & vlSelf->dma__DOT___dst_offset) 
                             | (0x10000U & ((IData)(vlSelf->reg_data) 
                                            << 0x10U)))
                          : ((0x100ffU & vlSelf->dma__DOT___dst_offset) 
                             | ((IData)(vlSelf->reg_data) 
                                << 8U)));
                }
            } else if ((2U & (IData)(vlSelf->reg_addr))) {
                if ((1U & (IData)(vlSelf->reg_addr))) {
                    vlSelf->dma__DOT___dst_offset = 
                        ((0x1ff00U & vlSelf->dma__DOT___dst_offset) 
                         | (IData)(vlSelf->reg_data));
                }
            }
            if ((1U & (~ ((IData)(vlSelf->reg_addr) 
                          >> 2U)))) {
                if ((2U & (IData)(vlSelf->reg_addr))) {
                    if ((1U & (~ (IData)(vlSelf->reg_addr)))) {
                        vlSelf->dma__DOT___src_offset 
                            = ((0xffffU & vlSelf->dma__DOT___src_offset) 
                               | (0x10000U & ((IData)(vlSelf->reg_data) 
                                              << 0x10U)));
                    }
                } else {
                    vlSelf->dma__DOT___src_offset = 
                        ((1U & (IData)(vlSelf->reg_addr))
                          ? ((0x100ffU & vlSelf->dma__DOT___src_offset) 
                             | ((IData)(vlSelf->reg_data) 
                                << 8U)) : ((0x1ff00U 
                                            & vlSelf->dma__DOT___src_offset) 
                                           | (IData)(vlSelf->reg_data)));
                }
            }
        }
    }
    vlSelf->dma__DOT___v_src_addr = ((0x1ff00U & ((
                                                   (vlSelf->dma__DOT___src_offset 
                                                    >> 8U) 
                                                   + (IData)(vlSelf->col_cnt)) 
                                                  << 8U)) 
                                     | (0xffU & (vlSelf->dma__DOT___src_offset 
                                                 + (IData)(vlSelf->row_cnt))));
}

VL_INLINE_OPT void Vdma___024root___nba_sequent__TOP__2(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___nba_sequent__TOP__2\n"); );
    // Body
    if ((((IData)(vlSelf->reg_cs) & (IData)(vlSelf->reg_we)) 
         & (0xaU == (IData)(vlSelf->reg_addr)))) {
        vlSelf->dma__DOT___run_reg = (1U & ((IData)(vlSelf->reg_data) 
                                            >> 1U));
    } else if (vlSelf->cnt_done) {
        vlSelf->dma__DOT___run_reg = 0U;
    }
    vlSelf->reg_wait = ((IData)(vlSelf->dma__DOT___run_reg) 
                        & (IData)(vlSelf->reg_cs));
    vlSelf->dma__DOT____VdfgTmp_h2394a40c__0 = ((IData)(vlSelf->dma__DOT___run_reg) 
                                                & (IData)(vlSelf->v_blank));
}

VL_INLINE_OPT void Vdma___024root___nba_comb__TOP__0(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___nba_comb__TOP__0\n"); );
    // Init
    CData/*0:0*/ dma__DOT____VdfgTmp_hbe67ea25__0;
    dma__DOT____VdfgTmp_hbe67ea25__0 = 0;
    // Body
    dma__DOT____VdfgTmp_hbe67ea25__0 = ((IData)(vlSelf->dma__DOT___run_reg) 
                                        & ((IData)(vlSelf->dma__DOT___sub_state) 
                                           & (IData)(vlSelf->v_blank)));
    vlSelf->v_re = ((~ (IData)(vlSelf->dma__DOT___sub_state)) 
                    & (IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0));
    vlSelf->v_we = ((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0) 
                    & (IData)(vlSelf->dma__DOT___sub_state));
    vlSelf->v_data = (((IData)(dma__DOT____VdfgTmp_hbe67ea25__0)
                        ? 0xffU : 0U) & (((IData)(dma__DOT____VdfgTmp_hbe67ea25__0)
                                           ? (IData)(vlSelf->dma__DOT___src_data)
                                           : 0U) & 
                                         ((IData)(dma__DOT____VdfgTmp_hbe67ea25__0)
                                           ? 0xffU : 0U)));
}

VL_INLINE_OPT void Vdma___024root___nba_comb__TOP__1(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___nba_comb__TOP__1\n"); );
    // Body
    vlSelf->v_addr = (((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0)
                        ? 0x1ffffU : 0U) & (((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0)
                                              ? ((IData)(vlSelf->dma__DOT___sub_state)
                                                  ? vlSelf->dma__DOT___v_src_addr
                                                  : 
                                                 (vlSelf->dma__DOT___v_src_addr 
                                                  + vlSelf->dma__DOT___dst_offset))
                                              : 0U) 
                                            & ((IData)(vlSelf->dma__DOT____VdfgTmp_h2394a40c__0)
                                                ? 0x1ffffU
                                                : 0U)));
}

void Vdma___024root___eval_nba(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_nba\n"); );
    // Body
    if (vlSelf->__VnbaTriggered.at(2U)) {
        Vdma___024root___nba_sequent__TOP__0(vlSelf);
    }
    if (vlSelf->__VnbaTriggered.at(1U)) {
        Vdma___024root___nba_sequent__TOP__1(vlSelf);
    }
    if (vlSelf->__VnbaTriggered.at(0U)) {
        Vdma___024root___nba_sequent__TOP__2(vlSelf);
    }
    if ((vlSelf->__VnbaTriggered.at(0U) | vlSelf->__VnbaTriggered.at(2U))) {
        Vdma___024root___nba_comb__TOP__0(vlSelf);
    }
    if (((vlSelf->__VnbaTriggered.at(0U) | vlSelf->__VnbaTriggered.at(1U)) 
         | vlSelf->__VnbaTriggered.at(2U))) {
        Vdma___024root___nba_comb__TOP__1(vlSelf);
    }
}

void Vdma___024root___eval_triggers__ico(Vdma___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__ico(Vdma___024root* vlSelf);
#endif  // VL_DEBUG
void Vdma___024root___eval_triggers__act(Vdma___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__act(Vdma___024root* vlSelf);
#endif  // VL_DEBUG
#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__nba(Vdma___024root* vlSelf);
#endif  // VL_DEBUG

void Vdma___024root___eval(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval\n"); );
    // Init
    CData/*0:0*/ __VicoContinue;
    VlTriggerVec<3> __VpreTriggered;
    IData/*31:0*/ __VnbaIterCount;
    CData/*0:0*/ __VnbaContinue;
    // Body
    vlSelf->__VicoIterCount = 0U;
    __VicoContinue = 1U;
    while (__VicoContinue) {
        __VicoContinue = 0U;
        Vdma___024root___eval_triggers__ico(vlSelf);
        if (vlSelf->__VicoTriggered.any()) {
            __VicoContinue = 1U;
            if (VL_UNLIKELY((0x64U < vlSelf->__VicoIterCount))) {
#ifdef VL_DEBUG
                Vdma___024root___dump_triggers__ico(vlSelf);
#endif
                VL_FATAL_MT("dma.v", 1, "", "Input combinational region did not converge.");
            }
            vlSelf->__VicoIterCount = ((IData)(1U) 
                                       + vlSelf->__VicoIterCount);
            Vdma___024root___eval_ico(vlSelf);
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
            Vdma___024root___eval_triggers__act(vlSelf);
            if (vlSelf->__VactTriggered.any()) {
                vlSelf->__VactContinue = 1U;
                if (VL_UNLIKELY((0x64U < vlSelf->__VactIterCount))) {
#ifdef VL_DEBUG
                    Vdma___024root___dump_triggers__act(vlSelf);
#endif
                    VL_FATAL_MT("dma.v", 1, "", "Active region did not converge.");
                }
                vlSelf->__VactIterCount = ((IData)(1U) 
                                           + vlSelf->__VactIterCount);
                __VpreTriggered.andNot(vlSelf->__VactTriggered, vlSelf->__VnbaTriggered);
                vlSelf->__VnbaTriggered.set(vlSelf->__VactTriggered);
                Vdma___024root___eval_act(vlSelf);
            }
        }
        if (vlSelf->__VnbaTriggered.any()) {
            __VnbaContinue = 1U;
            if (VL_UNLIKELY((0x64U < __VnbaIterCount))) {
#ifdef VL_DEBUG
                Vdma___024root___dump_triggers__nba(vlSelf);
#endif
                VL_FATAL_MT("dma.v", 1, "", "NBA region did not converge.");
            }
            __VnbaIterCount = ((IData)(1U) + __VnbaIterCount);
            Vdma___024root___eval_nba(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
void Vdma___024root___eval_debug_assertions(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_debug_assertions\n"); );
    // Body
    if (VL_UNLIKELY((vlSelf->reg_clk & 0xfeU))) {
        Verilated::overWidthError("reg_clk");}
    if (VL_UNLIKELY((vlSelf->reg_addr & 0xf0U))) {
        Verilated::overWidthError("reg_addr");}
    if (VL_UNLIKELY((vlSelf->reg_cs & 0xfeU))) {
        Verilated::overWidthError("reg_cs");}
    if (VL_UNLIKELY((vlSelf->reg_we & 0xfeU))) {
        Verilated::overWidthError("reg_we");}
    if (VL_UNLIKELY((vlSelf->v_clk & 0xfeU))) {
        Verilated::overWidthError("v_clk");}
    if (VL_UNLIKELY((vlSelf->v_addr & 0xfffe0000U))) {
        Verilated::overWidthError("v_addr");}
    if (VL_UNLIKELY((vlSelf->v_blank & 0xfeU))) {
        Verilated::overWidthError("v_blank");}
    if (VL_UNLIKELY((vlSelf->col_cnt & 0xfe00U))) {
        Verilated::overWidthError("col_cnt");}
    if (VL_UNLIKELY((vlSelf->cnt_done & 0xfeU))) {
        Verilated::overWidthError("cnt_done");}
}
#endif  // VL_DEBUG
