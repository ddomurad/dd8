// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vdma.h for the primary calling header

#include "verilated.h"

#include "Vdma___024root.h"

VL_ATTR_COLD void Vdma___024root___eval_static(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_static\n"); );
}

VL_ATTR_COLD void Vdma___024root___eval_initial(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_initial\n"); );
    // Body
    vlSelf->__Vtrigprevexpr_hb954e6fa__0 = ((IData)(vlSelf->reg_clk) 
                                            | (IData)(vlSelf->v_clk));
    vlSelf->__Vtrigrprev__TOP__reg_clk = vlSelf->reg_clk;
    vlSelf->__Vtrigrprev__TOP__v_clk = vlSelf->v_clk;
}

VL_ATTR_COLD void Vdma___024root___eval_final(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_final\n"); );
}

VL_ATTR_COLD void Vdma___024root___eval_triggers__stl(Vdma___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__stl(Vdma___024root* vlSelf);
#endif  // VL_DEBUG
VL_ATTR_COLD void Vdma___024root___eval_stl(Vdma___024root* vlSelf);

VL_ATTR_COLD void Vdma___024root___eval_settle(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_settle\n"); );
    // Init
    CData/*0:0*/ __VstlContinue;
    // Body
    vlSelf->__VstlIterCount = 0U;
    __VstlContinue = 1U;
    while (__VstlContinue) {
        __VstlContinue = 0U;
        Vdma___024root___eval_triggers__stl(vlSelf);
        if (vlSelf->__VstlTriggered.any()) {
            __VstlContinue = 1U;
            if (VL_UNLIKELY((0x64U < vlSelf->__VstlIterCount))) {
#ifdef VL_DEBUG
                Vdma___024root___dump_triggers__stl(vlSelf);
#endif
                VL_FATAL_MT("dma.v", 1, "", "Settle region did not converge.");
            }
            vlSelf->__VstlIterCount = ((IData)(1U) 
                                       + vlSelf->__VstlIterCount);
            Vdma___024root___eval_stl(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__stl(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___dump_triggers__stl\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VstlTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VstlTriggered.at(0U)) {
        VL_DBG_MSGF("         'stl' region trigger index 0 is active: Internal 'stl' trigger - first iteration\n");
    }
}
#endif  // VL_DEBUG

void Vdma___024root___ico_sequent__TOP__0(Vdma___024root* vlSelf);

VL_ATTR_COLD void Vdma___024root___eval_stl(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_stl\n"); );
    // Body
    if (vlSelf->__VstlTriggered.at(0U)) {
        Vdma___024root___ico_sequent__TOP__0(vlSelf);
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__ico(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___dump_triggers__ico\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VicoTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VicoTriggered.at(0U)) {
        VL_DBG_MSGF("         'ico' region trigger index 0 is active: Internal 'ico' trigger - first iteration\n");
    }
}
#endif  // VL_DEBUG

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__act(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___dump_triggers__act\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VactTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VactTriggered.at(0U)) {
        VL_DBG_MSGF("         'act' region trigger index 0 is active: @(posedge (reg_clk | v_clk))\n");
    }
    if (vlSelf->__VactTriggered.at(1U)) {
        VL_DBG_MSGF("         'act' region trigger index 1 is active: @(posedge reg_clk)\n");
    }
    if (vlSelf->__VactTriggered.at(2U)) {
        VL_DBG_MSGF("         'act' region trigger index 2 is active: @(posedge v_clk)\n");
    }
}
#endif  // VL_DEBUG

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__nba(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___dump_triggers__nba\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VnbaTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VnbaTriggered.at(0U)) {
        VL_DBG_MSGF("         'nba' region trigger index 0 is active: @(posedge (reg_clk | v_clk))\n");
    }
    if (vlSelf->__VnbaTriggered.at(1U)) {
        VL_DBG_MSGF("         'nba' region trigger index 1 is active: @(posedge reg_clk)\n");
    }
    if (vlSelf->__VnbaTriggered.at(2U)) {
        VL_DBG_MSGF("         'nba' region trigger index 2 is active: @(posedge v_clk)\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vdma___024root___ctor_var_reset(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___ctor_var_reset\n"); );
    // Body
    vlSelf->reg_clk = VL_RAND_RESET_I(1);
    vlSelf->reg_addr = VL_RAND_RESET_I(4);
    vlSelf->reg_data = VL_RAND_RESET_I(8);
    vlSelf->reg_cs = VL_RAND_RESET_I(1);
    vlSelf->reg_we = VL_RAND_RESET_I(1);
    vlSelf->reg_wait = VL_RAND_RESET_I(1);
    vlSelf->v_clk = VL_RAND_RESET_I(1);
    vlSelf->v_addr = VL_RAND_RESET_I(17);
    vlSelf->v_data = VL_RAND_RESET_I(8);
    vlSelf->v_re = VL_RAND_RESET_I(1);
    vlSelf->v_we = VL_RAND_RESET_I(1);
    vlSelf->v_blank = VL_RAND_RESET_I(1);
    vlSelf->row_cnt = VL_RAND_RESET_I(8);
    vlSelf->col_cnt = VL_RAND_RESET_I(9);
    vlSelf->cnt_done = VL_RAND_RESET_I(1);
    vlSelf->dma__DOT___src_offset = VL_RAND_RESET_I(17);
    vlSelf->dma__DOT___dst_offset = VL_RAND_RESET_I(17);
    vlSelf->dma__DOT___run_reg = VL_RAND_RESET_I(1);
    vlSelf->dma__DOT___sub_state = VL_RAND_RESET_I(1);
    vlSelf->dma__DOT___src_data = VL_RAND_RESET_I(8);
    vlSelf->dma__DOT___v_src_addr = VL_RAND_RESET_I(17);
    vlSelf->dma__DOT____VdfgTmp_h2394a40c__0 = 0;
    vlSelf->__Vtrigprevexpr_hb954e6fa__0 = VL_RAND_RESET_I(1);
    vlSelf->__Vtrigrprev__TOP__reg_clk = VL_RAND_RESET_I(1);
    vlSelf->__Vtrigrprev__TOP__v_clk = VL_RAND_RESET_I(1);
}
