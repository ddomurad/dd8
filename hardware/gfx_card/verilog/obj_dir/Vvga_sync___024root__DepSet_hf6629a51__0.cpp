// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vvga_sync.h for the primary calling header

#include "verilated.h"

#include "Vvga_sync__Syms.h"
#include "Vvga_sync___024root.h"

#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__ico(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG

void Vvga_sync___024root___eval_triggers__ico(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_triggers__ico\n"); );
    // Body
    vlSelf->__VicoTriggered.at(0U) = (0U == vlSelf->__VicoIterCount);
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vvga_sync___024root___dump_triggers__ico(vlSelf);
    }
#endif
}

#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__act(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG

void Vvga_sync___024root___eval_triggers__act(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_triggers__act\n"); );
    // Body
    vlSelf->__VactTriggered.at(0U) = ((IData)(vlSelf->vga_sync__DOT__clk_int) 
                                      & (~ (IData)(vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__clk_int)));
    vlSelf->__VactTriggered.at(1U) = ((IData)(vlSelf->vga_sync__DOT__cnt_h_rst) 
                                      & (~ (IData)(vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__cnt_h_rst)));
    vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__clk_int 
        = vlSelf->vga_sync__DOT__clk_int;
    vlSelf->__Vtrigrprev__TOP__vga_sync__DOT__cnt_h_rst 
        = vlSelf->vga_sync__DOT__cnt_h_rst;
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vvga_sync___024root___dump_triggers__act(vlSelf);
    }
#endif
}
