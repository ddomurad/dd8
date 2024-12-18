// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vdram_rnd.h for the primary calling header

#include "verilated.h"

#include "Vdram_rnd___024root.h"

VL_ATTR_COLD void Vdram_rnd___024root___eval_static(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_static\n"); );
}

VL_ATTR_COLD void Vdram_rnd___024root___eval_initial(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_initial\n"); );
    // Body
    vlSelf->__Vtrigrprev__TOP__clk = vlSelf->clk;
}

VL_ATTR_COLD void Vdram_rnd___024root___eval_final(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_final\n"); );
}

VL_ATTR_COLD void Vdram_rnd___024root___eval_triggers__stl(Vdram_rnd___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vdram_rnd___024root___dump_triggers__stl(Vdram_rnd___024root* vlSelf);
#endif  // VL_DEBUG
VL_ATTR_COLD void Vdram_rnd___024root___eval_stl(Vdram_rnd___024root* vlSelf);

VL_ATTR_COLD void Vdram_rnd___024root___eval_settle(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_settle\n"); );
    // Init
    CData/*0:0*/ __VstlContinue;
    // Body
    vlSelf->__VstlIterCount = 0U;
    __VstlContinue = 1U;
    while (__VstlContinue) {
        __VstlContinue = 0U;
        Vdram_rnd___024root___eval_triggers__stl(vlSelf);
        if (vlSelf->__VstlTriggered.any()) {
            __VstlContinue = 1U;
            if (VL_UNLIKELY((0x64U < vlSelf->__VstlIterCount))) {
#ifdef VL_DEBUG
                Vdram_rnd___024root___dump_triggers__stl(vlSelf);
#endif
                VL_FATAL_MT("dram_rnd.v", 1, "", "Settle region did not converge.");
            }
            vlSelf->__VstlIterCount = ((IData)(1U) 
                                       + vlSelf->__VstlIterCount);
            Vdram_rnd___024root___eval_stl(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdram_rnd___024root___dump_triggers__stl(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___dump_triggers__stl\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VstlTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VstlTriggered.at(0U)) {
        VL_DBG_MSGF("         'stl' region trigger index 0 is active: Internal 'stl' trigger - first iteration\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vdram_rnd___024root___stl_sequent__TOP__0(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___stl_sequent__TOP__0\n"); );
    // Body
    vlSelf->h_sync = ((0x290U > (IData)(vlSelf->h_cnt)) 
                      | (0x2efU < (IData)(vlSelf->h_cnt)));
    vlSelf->v_sync = ((0x1eaU > (IData)(vlSelf->v_cnt)) 
                      | (0x1ebU < (IData)(vlSelf->v_cnt)));
    vlSelf->d_out = ((0x280U > (IData)(vlSelf->h_cnt)) 
                     & (0x1e0U > (IData)(vlSelf->v_cnt)));
    vlSelf->d_out_b = (1U & (~ (IData)(vlSelf->d_out)));
}

VL_ATTR_COLD void Vdram_rnd___024root___eval_stl(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___eval_stl\n"); );
    // Body
    if (vlSelf->__VstlTriggered.at(0U)) {
        Vdram_rnd___024root___stl_sequent__TOP__0(vlSelf);
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdram_rnd___024root___dump_triggers__act(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___dump_triggers__act\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VactTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VactTriggered.at(0U)) {
        VL_DBG_MSGF("         'act' region trigger index 0 is active: @(posedge clk)\n");
    }
}
#endif  // VL_DEBUG

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdram_rnd___024root___dump_triggers__nba(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___dump_triggers__nba\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VnbaTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VnbaTriggered.at(0U)) {
        VL_DBG_MSGF("         'nba' region trigger index 0 is active: @(posedge clk)\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vdram_rnd___024root___ctor_var_reset(Vdram_rnd___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdram_rnd__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdram_rnd___024root___ctor_var_reset\n"); );
    // Body
    vlSelf->clk = VL_RAND_RESET_I(1);
    vlSelf->h_sync = VL_RAND_RESET_I(1);
    vlSelf->v_sync = VL_RAND_RESET_I(1);
    vlSelf->d_out = VL_RAND_RESET_I(1);
    vlSelf->d_out_b = VL_RAND_RESET_I(1);
    vlSelf->h_cnt = VL_RAND_RESET_I(10);
    vlSelf->v_cnt = VL_RAND_RESET_I(10);
    vlSelf->__Vtrigrprev__TOP__clk = VL_RAND_RESET_I(1);
}
