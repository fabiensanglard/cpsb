# VSYNC interrupt handler, jumps to C function.
.align 4
VSync:
    jsr onVSync
	rte
