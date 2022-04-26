.extern	main

dc.l 0xFFF000, _boot,   Def, Def, Def, Def, Def, Def
dc.l      Def,   Def,   Def, Def, Def, Def, Def, Def
dc.l      Def,   Def,   Def, Def, Def, Def, Def, Def
dc.l      Def,   Def, VSync, Def, Def, Def, Def, Def
dc.l      Def,   Def,   Def, Def, Def, Def, Def, Def
dc.l      Def,   Def,   Def, Def, Def, Def, Def, Def
dc.l      Def,   Def,   Def, Def, Def, Def, Def, Def
dc.l      Def,   Def,   Def, Def, Def, Def, Def, Def

.align 4
Def:
	rte