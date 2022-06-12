#define GFXRAM __attribute__ ((section (".gfx_data")))
#define CPSA   __attribute__ ((section (".cpsa_reg")))  
#define CPSB   __attribute__ ((section (".cpsb_reg")))  

GFXRAM short palettes[6 * 32 * 16];




CPSA   short cpsa_reg[0x20] = {};
CPSB   short cpsb_reg[0x20] = {};