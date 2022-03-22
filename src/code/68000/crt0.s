.globl	_start

.extern	main

dc.l 0xFFF000, _boot,   Def, Def, Def, Def, Def, Def
dc.l Def,        Def,   Def, Def, Def, Def, Def, Def
dc.l Def,        Def,   Def, Def, Def, Def, Def, Def
dc.l Def,        Def, VSync, Def, Def, Def, Def, Def
dc.l Def,        Def,   Def, Def, Def, Def, Def, Def
dc.l Def,        Def,   Def, Def, Def, Def, Def, Def
dc.l Def,        Def,   Def, Def, Def, Def, Def, Def
dc.l Def,        Def,   Def, Def, Def, Def, Def, Def

.align 4
Def:
	rte
	
.align 4
VSync:
    jsr onVSync
	rte

.align 4
_boot:
	* Enable interrupts
	move.w	#0x2000, sr

	* Jump to C main()
	jbsr	main

EndLoop:
	bra.s	EndLoop
