CLK_PER_BYTE = 1
EMPTY_CLK_PER_BATCH = 2

VGA_CLK_T_MHZ = 25.175
VGA_RES = (800, 525)
IMG_RES = (256, 240)  # 2x2 pixels
IMG_CYCLES = IMG_RES[0]*IMG_RES[1]
IMG_CYCLES2 = IMG_RES[0]*IMG_RES[1]*2  # H-lines are repeated
DMA_FREE_CYCLES = (VGA_RES[0]*VGA_RES[1]) - IMG_CYCLES2

# src_addr_l, src_addr_h, dst_addr_l,
# dst_addr_h, width, height, mask, status = 8 bytes
# CPU runs x2 slower than the dma
DMA_FULL_SETUP_BYTES = 8
CPU_AVG_CLK_PER_INSTRUCTION = 4
CPU_CLK_MUL = 2
CPU_DMA_SETUP_CLK = CPU_CLK_MUL * \
    DMA_FULL_SETUP_BYTES*CPU_AVG_CLK_PER_INSTRUCTION


def get_time(clk, relative_to=0):
    n = relative_to/clk
    t_ns = (1/VGA_CLK_T_MHZ)*clk
    if relative_to == 0:
        return f"{clk} ({t_ns}ns)"
    return f"{clk} ({t_ns}ns)[{n} times]"


print("vga clk:", VGA_CLK_T_MHZ, "Mhz")
print("dma setup:", get_time(CPU_DMA_SETUP_CLK))
print("dma available cycles:", DMA_FREE_CYCLES)
print("relative to full screen:", VGA_RES, "px")
print()
_screen_fill = IMG_CYCLES + CPU_DMA_SETUP_CLK
print("whole screen fill:", get_time(_screen_fill, DMA_FREE_CYCLES), "@ one frame")
_screen_fill = IMG_RES[0]*32 + CPU_DMA_SETUP_CLK
print("256x32 fill:", get_time(_screen_fill, DMA_FREE_CYCLES), "@ one frame")
_32_lines_cycles = VGA_RES[0]*32 - IMG_RES[1]*32
print("256x32 fill:", get_time(_screen_fill, _32_lines_cycles), "@ 32 lines")
_1_fill = 1 + CPU_DMA_SETUP_CLK
print("1x1 fill:", get_time(_1_fill, _32_lines_cycles), "@ 32 lines")
_64_fill = 64 + CPU_DMA_SETUP_CLK
print("8x8 fill:", get_time(_64_fill, _32_lines_cycles), "@ 32 lines")
_1024_fill = 1024 + CPU_DMA_SETUP_CLK
print("32x32 fill:", get_time(_1024_fill, _32_lines_cycles), "@ 32 lines")
