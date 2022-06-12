.area _GSINIT                ; Initialize global variables
gsinit:                      ; Copy values from ROM > RAM.
   ld  bc, #l__INITIALIZER
   ld  a, b
   or  a, c
   jr  Z, gsinit_next
   ld  de, #s__INITIALIZED
   ld  hl, #s__INITIALIZER
   ldir
gsinit_next:
   ret