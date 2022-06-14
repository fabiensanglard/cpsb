extern char _etext, _data, _edata, _bss, _ebss;
char *src = &_etext;
char *dst = &_data;

// Copy ROM to DATA */
while (dst < &_edata) {
  *dst++ = *src++;
}

// Zero bss.
for (dst = &_bss; dst< &_ebss; dst++) {
  *dst = 0;
}
