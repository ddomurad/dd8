// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vsprite_render.h for the primary calling header

#include "verilated.h"

#include "Vsprite_render__Syms.h"
#include "Vsprite_render___024root.h"

#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_render___024root___dump_triggers__ico(Vsprite_render___024root* vlSelf);
#endif  // VL_DEBUG

void Vsprite_render___024root___eval_triggers__ico(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___eval_triggers__ico\n"); );
    // Body
    vlSelf->__VicoTriggered.at(0U) = (0U == vlSelf->__VicoIterCount);
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vsprite_render___024root___dump_triggers__ico(vlSelf);
    }
#endif
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vsprite_render___024root___dump_triggers__act(Vsprite_render___024root* vlSelf);
#endif  // VL_DEBUG

void Vsprite_render___024root___eval_triggers__act(Vsprite_render___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vsprite_render__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vsprite_render___024root___eval_triggers__act\n"); );
    // Body
    vlSelf->__VactTriggered.at(0U) = ((IData)(vlSelf->clk) 
                                      & (~ (IData)(vlSelf->__Vtrigrprev__TOP__clk)));
    vlSelf->__VactTriggered.at(1U) = ((~ (IData)(vlSelf->clk)) 
                                      & (IData)(vlSelf->__Vtrigrprev__TOP__clk));
    vlSelf->__VactTriggered.at(2U) = ((IData)(vlSelf->i_clk) 
                                      & (~ (IData)(vlSelf->__Vtrigrprev__TOP__i_clk)));
    vlSelf->__Vtrigrprev__TOP__clk = vlSelf->clk;
    vlSelf->__Vtrigrprev__TOP__i_clk = vlSelf->i_clk;
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vsprite_render___024root___dump_triggers__act(vlSelf);
    }
#endif
}
