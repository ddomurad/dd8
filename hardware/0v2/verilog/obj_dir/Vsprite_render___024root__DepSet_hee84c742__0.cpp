// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vsprite_render.h for the primary calling header

#include "verilated.h"

#include "Vsprite_render___024root.h"

VL_INLINE_OPT void Vsprite_render___024root___ico_sequent__TOP__0(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___ico_sequent__TOP__0\n"); );
    // Body
    vlSelf->r_addr = ((0x200U & ((~ ((IData)(vlSelf->line_n) 
                                     >> 1U)) << 9U)) 
                      | (0x1ffU & ((vlSelf->dma__DOT___spr_reg 
                                    >> 8U) + (IData)(vlSelf->dma__DOT___cpy_cnt))));
    vlSelf->s_we = ((IData)(vlSelf->off_src_line) & 
                    ((IData)(vlSelf->i_cs) & ((IData)(vlSelf->i_we) 
                                              & (3U 
                                                 == (IData)(vlSelf->i_addr)))));
    vlSelf->s_re = (1U & (~ (IData)(vlSelf->off_src_line)));
    vlSelf->s_data = (((IData)(vlSelf->off_src_line)
                        ? 0xffU : 0U) & (((IData)(vlSelf->off_src_line)
                                           ? (IData)(vlSelf->i_data)
                                           : 0U) & 
                                         ((IData)(vlSelf->off_src_line)
                                           ? 0xffU : 0U)));
    vlSelf->dma__DOT___y_diff = (0xffU & (((IData)(vlSelf->line_n) 
                                           >> 1U) - vlSelf->dma__DOT___spr_reg));
    vlSelf->i_wait = ((IData)(vlSelf->i_cs) & ((IData)(vlSelf->s_re) 
                                               & (3U 
                                                  == (IData)(vlSelf->i_addr))));
    vlSelf->r_we = ((4U == (IData)(vlSelf->dma__DOT___state)) 
                    & ((IData)(vlSelf->clk) & (0U != (IData)(vlSelf->s_data))));
    vlSelf->s_addr = ((IData)(vlSelf->off_src_line)
                       ? (IData)(vlSelf->dma__DOT___i_addr)
                       : ((4U == (IData)(vlSelf->dma__DOT___state))
                           ? ((0x7fc0U & (vlSelf->dma__DOT___spr_reg 
                                          >> 0x11U)) 
                              | ((0x38U & ((IData)(vlSelf->dma__DOT___y_diff) 
                                           << 3U)) 
                                 | (IData)(vlSelf->dma__DOT___cpy_cnt)))
                           : (((IData)(vlSelf->dma__DOT___spr_cnt) 
                               << 2U) | (3U & (IData)(vlSelf->dma__DOT___state)))));
}

void Vsprite_render___024root___eval_ico(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___eval_ico\n"); );
    // Body
    if (vlSelf->__VicoTriggered.at(0U)) {
        Vsprite_render___024root___ico_sequent__TOP__0(vlSelf);
    }
}

void Vsprite_render___024root___eval_act(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___eval_act\n"); );
}

VL_INLINE_OPT void Vsprite_render___024root___nba_sequent__TOP__0(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___nba_sequent__TOP__0\n"); );
    // Body
    if (((IData)(vlSelf->i_cs) & (IData)(vlSelf->i_we))) {
        if ((0U == (IData)(vlSelf->i_addr))) {
            vlSelf->dma__DOT___i_addr = ((0x7f00U & (IData)(vlSelf->dma__DOT___i_addr)) 
                                         | (IData)(vlSelf->i_data));
        } else if ((1U == (IData)(vlSelf->i_addr))) {
            vlSelf->dma__DOT___i_addr = ((0xffU & (IData)(vlSelf->dma__DOT___i_addr)) 
                                         | (0x7f00U 
                                            & ((IData)(vlSelf->i_data) 
                                               << 8U)));
        }
    }
}

VL_INLINE_OPT void Vsprite_render___024root___nba_sequent__TOP__1(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___nba_sequent__TOP__1\n"); );
    // Body
    if ((1U & (~ ((IData)(vlSelf->dma__DOT___state) 
                  >> 2U)))) {
        vlSelf->dma__DOT___spr_reg = ((2U & (IData)(vlSelf->dma__DOT___state))
                                       ? ((1U & (IData)(vlSelf->dma__DOT___state))
                                           ? ((0xffffffU 
                                               & vlSelf->dma__DOT___spr_reg) 
                                              | ((IData)(vlSelf->s_data) 
                                                 << 0x18U))
                                           : ((0xff00ffffU 
                                               & vlSelf->dma__DOT___spr_reg) 
                                              | ((IData)(vlSelf->s_data) 
                                                 << 0x10U)))
                                       : ((1U & (IData)(vlSelf->dma__DOT___state))
                                           ? ((0xffff00ffU 
                                               & vlSelf->dma__DOT___spr_reg) 
                                              | ((IData)(vlSelf->s_data) 
                                                 << 8U))
                                           : ((0xffffff00U 
                                               & vlSelf->dma__DOT___spr_reg) 
                                              | (IData)(vlSelf->s_data))));
    }
    vlSelf->dma__DOT___y_diff = (0xffU & (((IData)(vlSelf->line_n) 
                                           >> 1U) - vlSelf->dma__DOT___spr_reg));
}

VL_INLINE_OPT void Vsprite_render___024root___nba_sequent__TOP__2(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___nba_sequent__TOP__2\n"); );
    // Body
    if (((IData)(vlSelf->dma__DOT___last_rst) != (1U 
                                                  & (IData)(vlSelf->line_n)))) {
        if ((1U & (~ (IData)(vlSelf->line_n)))) {
            vlSelf->dma__DOT___state = 0U;
            vlSelf->dma__DOT___cpy_cnt = 0U;
            vlSelf->dma__DOT___spr_cnt = 0U;
        }
        vlSelf->dma__DOT___last_rst = (1U & (IData)(vlSelf->line_n));
    } else if ((7U < (IData)(vlSelf->dma__DOT___y_diff))) {
        vlSelf->dma__DOT___spr_cnt = (0x3fU & ((IData)(1U) 
                                               + (IData)(vlSelf->dma__DOT___spr_cnt)));
        vlSelf->dma__DOT___cpy_cnt = 0U;
        vlSelf->dma__DOT___state = 0U;
    } else if (vlSelf->dma__DOT___write_state) {
        if ((7U == (IData)(vlSelf->dma__DOT___cpy_cnt))) {
            vlSelf->dma__DOT___spr_cnt = (0x3fU & ((IData)(1U) 
                                                   + (IData)(vlSelf->dma__DOT___spr_cnt)));
            vlSelf->dma__DOT___cpy_cnt = 0U;
            vlSelf->dma__DOT___state = 0U;
        } else {
            vlSelf->dma__DOT___cpy_cnt = (7U & ((IData)(1U) 
                                                + (IData)(vlSelf->dma__DOT___cpy_cnt)));
        }
    } else {
        vlSelf->dma__DOT___state = (7U & ((IData)(1U) 
                                          + (IData)(vlSelf->dma__DOT___state)));
    }
    vlSelf->dma__DOT___write_state = (4U == (IData)(vlSelf->dma__DOT___state));
}

VL_INLINE_OPT void Vsprite_render___024root___nba_comb__TOP__0(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___nba_comb__TOP__0\n"); );
    // Body
    vlSelf->r_addr = ((0x200U & ((~ ((IData)(vlSelf->line_n) 
                                     >> 1U)) << 9U)) 
                      | (0x1ffU & ((vlSelf->dma__DOT___spr_reg 
                                    >> 8U) + (IData)(vlSelf->dma__DOT___cpy_cnt))));
}

VL_INLINE_OPT void Vsprite_render___024root___nba_comb__TOP__1(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___nba_comb__TOP__1\n"); );
    // Body
    vlSelf->s_addr = ((IData)(vlSelf->off_src_line)
                       ? (IData)(vlSelf->dma__DOT___i_addr)
                       : ((4U == (IData)(vlSelf->dma__DOT___state))
                           ? ((0x7fc0U & (vlSelf->dma__DOT___spr_reg 
                                          >> 0x11U)) 
                              | ((0x38U & ((IData)(vlSelf->dma__DOT___y_diff) 
                                           << 3U)) 
                                 | (IData)(vlSelf->dma__DOT___cpy_cnt)))
                           : (((IData)(vlSelf->dma__DOT___spr_cnt) 
                               << 2U) | (3U & (IData)(vlSelf->dma__DOT___state)))));
}

VL_INLINE_OPT void Vsprite_render___024root___nba_sequent__TOP__3(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___nba_sequent__TOP__3\n"); );
    // Body
    vlSelf->r_we = ((4U == (IData)(vlSelf->dma__DOT___state)) 
                    & ((IData)(vlSelf->clk) & (0U != (IData)(vlSelf->s_data))));
}

void Vsprite_render___024root___eval_nba(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___eval_nba\n"); );
    // Body
    if (vlSelf->__VnbaTriggered.at(2U)) {
        Vsprite_render___024root___nba_sequent__TOP__0(vlSelf);
    }
    if (vlSelf->__VnbaTriggered.at(0U)) {
        Vsprite_render___024root___nba_sequent__TOP__1(vlSelf);
    }
    if (vlSelf->__VnbaTriggered.at(1U)) {
        Vsprite_render___024root___nba_sequent__TOP__2(vlSelf);
    }
    if ((vlSelf->__VnbaTriggered.at(0U) | vlSelf->__VnbaTriggered.at(1U))) {
        Vsprite_render___024root___nba_comb__TOP__0(vlSelf);
    }
    if (((vlSelf->__VnbaTriggered.at(0U) | vlSelf->__VnbaTriggered.at(1U)) 
         | vlSelf->__VnbaTriggered.at(2U))) {
        Vsprite_render___024root___nba_comb__TOP__1(vlSelf);
    }
    if (vlSelf->__VnbaTriggered.at(1U)) {
        Vsprite_render___024root___nba_sequent__TOP__3(vlSelf);
    }
}

void Vsprite_render___024root___eval_triggers__ico(Vsprite_render___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_render___024root___dump_triggers__ico(Vsprite_render___024root* vlSelf);
#endif  // VL_DEBUG
void Vsprite_render___024root___eval_triggers__act(Vsprite_render___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_render___024root___dump_triggers__act(Vsprite_render___024root* vlSelf);
#endif  // VL_DEBUG
#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_render___024root___dump_triggers__nba(Vsprite_render___024root* vlSelf);
#endif  // VL_DEBUG

void Vsprite_render___024root___eval(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___eval\n"); );
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
        Vsprite_render___024root___eval_triggers__ico(vlSelf);
        if (vlSelf->__VicoTriggered.any()) {
            __VicoContinue = 1U;
            if (VL_UNLIKELY((0x64U < vlSelf->__VicoIterCount))) {
#ifdef VL_DEBUG
                Vsprite_render___024root___dump_triggers__ico(vlSelf);
#endif
                VL_FATAL_MT("sprite_render.v", 1, "", "Input combinational region did not converge.");
            }
            vlSelf->__VicoIterCount = ((IData)(1U) 
                                       + vlSelf->__VicoIterCount);
            Vsprite_render___024root___eval_ico(vlSelf);
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
            Vsprite_render___024root___eval_triggers__act(vlSelf);
            if (vlSelf->__VactTriggered.any()) {
                vlSelf->__VactContinue = 1U;
                if (VL_UNLIKELY((0x64U < vlSelf->__VactIterCount))) {
#ifdef VL_DEBUG
                    Vsprite_render___024root___dump_triggers__act(vlSelf);
#endif
                    VL_FATAL_MT("sprite_render.v", 1, "", "Active region did not converge.");
                }
                vlSelf->__VactIterCount = ((IData)(1U) 
                                           + vlSelf->__VactIterCount);
                __VpreTriggered.andNot(vlSelf->__VactTriggered, vlSelf->__VnbaTriggered);
                vlSelf->__VnbaTriggered.set(vlSelf->__VactTriggered);
                Vsprite_render___024root___eval_act(vlSelf);
            }
        }
        if (vlSelf->__VnbaTriggered.any()) {
            __VnbaContinue = 1U;
            if (VL_UNLIKELY((0x64U < __VnbaIterCount))) {
#ifdef VL_DEBUG
                Vsprite_render___024root___dump_triggers__nba(vlSelf);
#endif
                VL_FATAL_MT("sprite_render.v", 1, "", "NBA region did not converge.");
            }
            __VnbaIterCount = ((IData)(1U) + __VnbaIterCount);
            Vsprite_render___024root___eval_nba(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
void Vsprite_render___024root___eval_debug_assertions(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___eval_debug_assertions\n"); );
    // Body
    if (VL_UNLIKELY((vlSelf->clk & 0xfeU))) {
        Verilated::overWidthError("clk");}
    if (VL_UNLIKELY((vlSelf->line_n & 0xfe00U))) {
        Verilated::overWidthError("line_n");}
    if (VL_UNLIKELY((vlSelf->off_src_line & 0xfeU))) {
        Verilated::overWidthError("off_src_line");}
    if (VL_UNLIKELY((vlSelf->i_clk & 0xfeU))) {
        Verilated::overWidthError("i_clk");}
    if (VL_UNLIKELY((vlSelf->i_addr & 0xfcU))) {
        Verilated::overWidthError("i_addr");}
    if (VL_UNLIKELY((vlSelf->i_we & 0xfeU))) {
        Verilated::overWidthError("i_we");}
    if (VL_UNLIKELY((vlSelf->i_cs & 0xfeU))) {
        Verilated::overWidthError("i_cs");}
}
#endif  // VL_DEBUG
