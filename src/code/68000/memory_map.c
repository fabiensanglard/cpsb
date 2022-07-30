#define ALIGN(X) __attribute__ ((aligned (X)))
#define GFXRAM __attribute__ ((section (".gfx_data")))
#define CPSA   __attribute__ ((section (".cpsa_reg")))  
#define CPSB   __attribute__ ((section (".cpsb_reg")))  
// ... All other SECTIONS here.

GFXRAM ALIGN(256) short palettes[6 * 32 * 16];
CPSA   short cpsa_reg[0x20] = {};
CPSB   short cpsb_reg[0x20] = {};
// ... All memory mapped data structures here.
