;----------------------------
; Initialize global variables
; Copy values from ROM > RAM.
;----------------------------
.area _GSINIT
gsinit:
   ld  bc, #l__INITIALIZER
   ld  a, b
   or  a, c
   jr  Z, gsinit_next
   ld  de, #s__INITIALIZED
   ld  hl, #s__INITIALIZER
   ldir
gsinit_next:
   .area _GSFINAL
   ret