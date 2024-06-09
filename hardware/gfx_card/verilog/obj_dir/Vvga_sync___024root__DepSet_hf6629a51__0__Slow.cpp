// Verilated -*- C++ -*-
// DESCRIPTION: Verilator output: Design implementation internals
// See Vvga_sync.h for the primary calling header

#include "verilated.h"

#include "Vvga_sync__Syms.h"
#include "Vvga_sync___024root.h"

#ifdef VL_DEBUG
VL_ATTR_COLD void Vvga_sync___024root___dump_triggers__stl(Vvga_sync___024root* vlSelf);
#endif  // VL_DEBUG

VL_ATTR_COLD void Vvga_sync___024root___eval_triggers__stl(Vvga_sync___024root* vlSelf) {
    if (false && vlSelf) {}  // Prevent unused
    Vvga_sync__Syms* const __restrict vlSymsp VL_ATTR_UNUSED = vlSelf->vlSymsp;
    VL_DEBUG_IF(VL_DBG_MSGF("+    Vvga_sync___024root___eval_triggers__stl\n"); );
    // Body
    vlSelf->__VstlTriggered.at(0U) = (0U == vlSelf->__VstlIterCount);
#ifdef VL_DEBUG
    if (VL_UNLIKELY(vlSymsp->_vm_contextp__->debug())) {
        Vvga_sync___024root___dump_triggers__stl(vlSelf);
    }
#endif
}
