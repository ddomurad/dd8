// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vvga_sync.h for the primary calling header

#include "verilated.h"

#include "Vvga_sync___024root.h"

VL_ATTR_COLD void Vvga_sync___024root___eval_static(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_static\n"); );
}

VL_ATTR_COLD void Vvga_sync___024root___eval_initial(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_initial\n"); );
    // Body
    vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__clk_int 
        = vlSelf->vga_sync__DOT__clk_int;
    vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__cnt_h_rst 
        = vlSelf->vga_sync__DOT__cnt_h_rst;
}

VL_ATTR_COLD void Vvga_sync___024root___eval_final(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_final\n"); );
}

VL_ATTR_COLD void Vvga_sync___024root___eval_triggers__stl(Vvga_sync___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__stl(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___eval_stl(Vvga_sync___024root* vlSelf);

VL_ATTR_COLD void Vvga_sync___024root___eval_settle(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_settle\n"); );
    // Init
    CData/*0:0*/ __VstlContinue;
    // Body
    vlSelf->__VstlIterCount = 0U;
    __VstlContinue = 1U;
    while (__VstlContinue) {
        __VstlContinue = 0U;
        Vvga_sync___024root___eval_triggers__stl(vlSelf);
        if (vlSelf->__VstlTriggered.any()) {
            __VstlContinue = 1U;
            if (VL_UNLIKELY((0x64U < vlSelf->__VstlIterCount))) {
#ifdef VL_DEBUG
                Vvga_sync___024root___dump_triggers__stl(vlSelf);
#endif
                VL_FATAL_MT("vga_sync.v", 10, "", "Settle region did not converge.");
            }
            vlSelf->__VstlIterCount = ((IData)(1U) 
                                       + vlSelf->__VstlIterCount);
            Vvga_sync___024root___eval_stl(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__stl(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___dump_triggers__stl\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VstlTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VstlTriggered.at(0U)) {
        VL_DBG_MSGF("         'stl' region trigger index 0 is active: Internal 'stl' trigger - first iteration\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vvga_sync___024root___stl_sequent__TOP__0(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___stl_sequent__TOP__0\n"); );
    // Body
    vlSelf->h_sync = ((0x28fU < (IData)(vlSelf->h_cnt)) 
                      & (0x2f0U > (IData)(vlSelf->h_cnt)));
    vlSelf->v_sync = ((0x1e9U < (IData)(vlSelf->v_cnt)) 
                      & (0x1ecU > (IData)(vlSelf->v_cnt)));
    vlSelf->vga_sync__DOT__cnt_h_rst = (0x320U == (IData)(vlSelf->h_cnt));
    vlSelf->vga_sync__DOT__cnt_v_rst = (0x20dU == (IData)(vlSelf->v_cnt));
    vlSelf->vga_sync__DOT__clk_int = ((IData)(vlSelf->clk) 
                                      & (IData)(vlSelf->en));
    vlSelf->d_out = ((0x280U > (IData)(vlSelf->h_cnt)) 
                     & (0x1e0U > (IData)(vlSelf->v_cnt)));
    vlSelf->d_out_b = (1U & (~ (IData)(vlSelf->d_out)));
}

VL_ATTR_COLD void Vvga_sync___024root___eval_stl(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_stl\n"); );
    // Body
    if (vlSelf->__VstlTriggered.at(0U)) {
        Vvga_sync___024root___stl_sequent__TOP__0(vlSelf);
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__ico(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___dump_triggers__ico\n"); );
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
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__act(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___dump_triggers__act\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VactTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VactTriggered.at(0U)) {
        VL_DBG_MSGF("         'act' region trigger index 0 is active: @(posedge vga_sync.clk_int)\n");
    }
    if (vlSelf->__VactTriggered.at(1U)) {
        VL_DBG_MSGF("         'act' region trigger index 1 is active: @(posedge vga_sync.cnt_h_rst)\n");
    }
}
#endif  // VL_DEBUG

#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__nba(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___dump_triggers__nba\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VnbaTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VnbaTriggered.at(0U)) {
        VL_DBG_MSGF("         'nba' region trigger index 0 is active: @(posedge vga_sync.clk_int)\n");
    }
    if (vlSelf->__VnbaTriggered.at(1U)) {
        VL_DBG_MSGF("         'nba' region trigger index 1 is active: @(posedge vga_sync.cnt_h_rst)\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vvga_sync___024root___ctor_var_reset(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___ctor_var_reset\n"); );
    // Body
    vlSelf->clk = VL_RAND_RESET_I(1);
    vlSelf->en = VL_RAND_RESET_I(1);
    vlSelf->h_sync = VL_RAND_RESET_I(1);
    vlSelf->v_sync = VL_RAND_RESET_I(1);
    vlSelf->d_out = VL_RAND_RESET_I(1);
    vlSelf->d_out_b = VL_RAND_RESET_I(1);
    vlSelf->h_cnt = VL_RAND_RESET_I(10);
    vlSelf->v_cnt = VL_RAND_RESET_I(10);
    vlSelf->vga_sync__DOT__cnt_h_rst = VL_RAND_RESET_I(1);
    vlSelf->vga_sync__DOT__cnt_v_rst = VL_RAND_RESET_I(1);
    vlSelf->vga_sync__DOT__clk_int = VL_RAND_RESET_I(1);
    vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__clk_int = VL_RAND_RESET_I(1);
    vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__cnt_h_rst = VL_RAND_RESET_I(1);
}
