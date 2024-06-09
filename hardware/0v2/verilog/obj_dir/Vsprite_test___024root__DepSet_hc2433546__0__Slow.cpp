// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vsprite_test.h for the primary calling header

#include "verilated.h"

#include "Vsprite_test___024root.h"

VL_ATTR_COLD void Vsprite_test___024root___eval_static(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___eval_static\n"); );
}

VL_ATTR_COLD void Vsprite_test___024root___eval_initial(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___eval_initial\n"); );
    // Body
    vlSelf->__Vtrigrprev__TOP__clk = vlSelf->clk;
}

VL_ATTR_COLD void Vsprite_test___024root___eval_final(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___eval_final\n"); );
}

VL_ATTR_COLD void Vsprite_test___024root___eval_triggers__stl(Vsprite_test___024root* vlSelf);
#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_test___024root___dump_triggers__stl(Vsprite_test___024root* vlSelf);
#endif  // VL_DEBUG
VL_ATTR_COLD void Vsprite_test___024root___eval_stl(Vsprite_test___024root* vlSelf);

VL_ATTR_COLD void Vsprite_test___024root___eval_settle(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___eval_settle\n"); );
    // Init
    CData/*0:0*/ __VstlContinue;
    // Body
    vlSelf->__VstlIterCount = 0U;
    __VstlContinue = 1U;
    while (__VstlContinue) {
        __VstlContinue = 0U;
        Vsprite_test___024root___eval_triggers__stl(vlSelf);
        if (vlSelf->__VstlTriggered.any()) {
            __VstlContinue = 1U;
            if (VL_UNLIKELY((0x64U < vlSelf->__VstlIterCount))) {
#ifdef VL_DEBUG
                Vsprite_test___024root___dump_triggers__stl(vlSelf);
#endif
                VL_FATAL_MT("sprite_test.v", 1, "", "Settle region did not converge.");
            }
            vlSelf->__VstlIterCount = ((IData)(1U) 
                                       + vlSelf->__VstlIterCount);
            Vsprite_test___024root___eval_stl(vlSelf);
        }
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_test___024root___dump_triggers__stl(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___dump_triggers__stl\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VstlTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VstlTriggered.at(0U)) {
        VL_DBG_MSGF("         'stl' region trigger index 0 is active: Internal 'stl' trigger - first iteration\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vsprite_test___024root___stl_sequent__TOP__0(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___stl_sequent__TOP__0\n"); );
    // Body
    vlSelf->dma__DOT___write_state = (4U == (IData)(vlSelf->dma__DOT___state));
    vlSelf->r_addr = ((0x200U & ((IData)(vlSelf->line_n) 
                                 << 8U)) | (0x1ffU 
                                            & ((vlSelf->dma__DOT___spr_reg 
                                                >> 8U) 
                                               + (IData)(vlSelf->dma__DOT___cpy_cnt))));
    vlSelf->r_we = ((4U == (IData)(vlSelf->dma__DOT___state)) 
                    & ((IData)(vlSelf->clk) & (0U != (IData)(vlSelf->s_data))));
    vlSelf->dma__DOT___y_diff = (0xffU & (((IData)(vlSelf->line_n) 
                                           >> 1U) - vlSelf->dma__DOT___spr_reg));
    vlSelf->s_addr = ((4U == (IData)(vlSelf->dma__DOT___state))
                       ? ((0x7fc0U & (vlSelf->dma__DOT___spr_reg 
                                      >> 0x11U)) | 
                          ((0x38U & ((IData)(vlSelf->dma__DOT___y_diff) 
                                     << 3U)) | (IData)(vlSelf->dma__DOT___cpy_cnt)))
                       : (((IData)(vlSelf->dma__DOT___spr_cnt) 
                           << 2U) | (3U & (IData)(vlSelf->dma__DOT___state))));
}

VL_ATTR_COLD void Vsprite_test___024root___eval_stl(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___eval_stl\n"); );
    // Body
    if (vlSelf->__VstlTriggered.at(0U)) {
        Vsprite_test___024root___stl_sequent__TOP__0(vlSelf);
    }
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_test___024root___dump_triggers__ico(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___dump_triggers__ico\n"); );
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
VL_ATTR_COLD void Vsprite_test___024root___dump_triggers__act(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___dump_triggers__act\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VactTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VactTriggered.at(0U)) {
        VL_DBG_MSGF("         'act' region trigger index 0 is active: @(posedge clk)\n");
    }
    if (vlSelf->__VactTriggered.at(1U)) {
        VL_DBG_MSGF("         'act' region trigger index 1 is active: @(negedge clk)\n");
    }
}
#endif  // VL_DEBUG

#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_test___024root___dump_triggers__nba(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___dump_triggers__nba\n"); );
    // Body
    if ((1U & (~ (IData)(vlSelf->__VnbaTriggered.any())))) {
        VL_DBG_MSGF("         No triggers active\n");
    }
    if (vlSelf->__VnbaTriggered.at(0U)) {
        VL_DBG_MSGF("         'nba' region trigger index 0 is active: @(posedge clk)\n");
    }
    if (vlSelf->__VnbaTriggered.at(1U)) {
        VL_DBG_MSGF("         'nba' region trigger index 1 is active: @(negedge clk)\n");
    }
}
#endif  // VL_DEBUG

VL_ATTR_COLD void Vsprite_test___024root___ctor_var_reset(Vsprite_test___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_test__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_test___024root___ctor_var_reset\n"); );
    // Body
    vlSelf->clk = VL_RAND_RESET_I(1);
    vlSelf->s_addr = VL_RAND_RESET_I(15);
    vlSelf->s_data = VL_RAND_RESET_I(8);
    vlSelf->s_re = VL_RAND_RESET_I(1);
    vlSelf->s_rs = VL_RAND_RESET_I(1);
    vlSelf->line_n = VL_RAND_RESET_I(9);
    vlSelf->r_addr = VL_RAND_RESET_I(10);
    vlSelf->r_we = VL_RAND_RESET_I(1);
    vlSelf->dma__DOT___state = VL_RAND_RESET_I(3);
    vlSelf->dma__DOT___spr_cnt = VL_RAND_RESET_I(6);
    vlSelf->dma__DOT___spr_reg = VL_RAND_RESET_I(32);
    vlSelf->dma__DOT___cpy_cnt = VL_RAND_RESET_I(3);
    vlSelf->dma__DOT___last_rst = VL_RAND_RESET_I(1);
    vlSelf->dma__DOT___y_diff = VL_RAND_RESET_I(8);
    vlSelf->dma__DOT___write_state = VL_RAND_RESET_I(1);
    vlSelf->__Vtrigrprev__TOP__clk = VL_RAND_RESET_I(1);
}
