.globl	_start

.extern	main

dc.l 0xFFF000, _start,    Def, Def, Def, Def, Def, Def
dc.l Def,         Def,    Def, Def, Def, Def, Def, Def
dc.l Def,         Def,    Def, Def, Def, Def, Def, Def
dc.l Def,         Def, VBlank, Def, Def, Def, Def, Def
dc.l Def,         Def,    Def, Def, Def, Def, Def, Def
dc.l Def,         Def,    Def, Def, Def, Def, Def, Def
dc.l Def,         Def,    Def, Def, Def, Def, Def, Def
dc.l Def,         Def,    Def, Def, Def, Def, Def, Def

.align 4
Def:
	rte
	
.align 4
VBlank:
    jsr onVSync
	rte

.align 4
_start:
	* Enable interrupts
	move.w	#0x2000, sr

	* Jump to mainloop
	jbsr	main

EndLoop:
	bra.s	EndLoop
