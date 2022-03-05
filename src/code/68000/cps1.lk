OUTPUT_FORMAT("elf32-m68k", "elf32-m68k", "elf32-m68k")
OUTPUT_ARCH(m68k)
ENTRY(_start)

MEMORY
{
  rom (rx)    : ORIGIN = 0x000000, LENGTH = 0x200000
  gfx_ram(rw) : ORIGIN = 0x900000, LENGTH = 0x2FFFF
  cpsa_reg(rw): ORIGIN = 0x800100, LENGTH = 0x40
  cpsb_reg(rw): ORIGIN = 0x800140, LENGTH = 0x40
  ram (rw)    : ORIGIN = 0xFF0000, LENGTH = 0xFFFF
}

SECTIONS {
  .text : {
    *(.text)
    *(.text.*)
    . = ALIGN(4);
  } > rom

  .rodata : {
    *(.rodata)
    *(.rodata.*)
    . = ALIGN(4);
  } > rom

  .gfx_data : {
  } > gfx_ram

  .cpsa_reg : {
  } > cpsa_reg

  .cpsb_reg : {
  } > cpsb_reg

  .bss : {
    __bss_start = .;
    *(.bss)
    *(.bss.*)
    _end = .;
    . = ALIGN(4);
  } > ram

  .data : {
    *(.data)
    *(.data.*)
    . = ALIGN(4);
  } > ram
}