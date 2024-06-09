// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vdma.h for the primary calling header

#include "verilated.h"

#include "Vdma__Syms.h"
#include "Vdma___024root.h"

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__ico(Vdma___024root* vlSelf);
#endif  // VL_DEBUG

void Vdma___024root___eval_triggers__ico(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_triggers__ico\n"); );
    // Body
    vlSelf->__VicoTriggered.at(0U) = (0U == vlSelf->__VicoIterCount);
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vdma___024root___dump_triggers__ico(vlSelf);
    }
#endif
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vdma___024root___dump_triggers__act(Vdma___024root* vlSelf);
#endif  // VL_DEBUG

void Vdma___024root___eval_triggers__act(Vdma___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vdma__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vdma___024root___eval_triggers__act\n"); );
    // Body
    CData/*0:0*/ __Vtrigcurrexpr_hb954e6fa__0;
    __Vtrigcurrexpr_hb954e6fa__0 = 0;
    __Vtrigcurrexpr_hb954e6fa__0 = ((IData)(vlSelf->reg_clk) 
                                    | (IData)(vlSelf->v_clk));
    vlSelf->__VactTriggered.at(0U) = ((IData)(__Vtrigcurrexpr_hb954e6fa__0) 
                                      & (~ (IData)(vlSelf->__Vtrigprevexpr_hb954e6fa__0)));
    vlSelf->__VactTriggered.at(1U) = ((IData)(vlSelf->reg_clk) 
                                      & (~ (IData)(vlSelf->__Vtrigrprev__TOP__reg_clk)));
    vlSelf->__VactTriggered.at(2U) = ((IData)(vlSelf->v_clk) 
                                      & (~ (IData)(vlSelf->__Vtrigrprev__TOP__v_clk)));
    vlSelf->__Vtrigprevexpr_hb954e6fa__0 = __Vtrigcurrexpr_hb954e6fa__0;
    vlSelf->__Vtrigrprev__TOP__reg_clk = vlSelf->reg_clk;
    vlSelf->__Vtrigrprev__TOP__v_clk = vlSelf->v_clk;
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vdma___024root___dump_triggers__act(vlSelf);
    }
#endif
}
