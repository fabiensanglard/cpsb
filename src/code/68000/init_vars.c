// These are defined by the linker via the linker script
extern char _etext, _data, _edata, _bss, _ebss;

char *src = &_etext;
char *dst = &_data;

// Copy ROM to DATA
void copyDATA() {
  while (dst < &_edata) {
    *dst++ = *src++;
  }
}

// Zero BSS.
void clearBSS() {
  for (dst = &_bss; dst< &_ebss; dst++) {
    *dst = 0;
  }
}
