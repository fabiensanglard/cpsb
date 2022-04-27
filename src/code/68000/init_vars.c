extern char _etext, _data, _edata, _bss, _ebss;
char *src = &_etext;
char *dst = &_data;

while (dst < &_edata) // Copy ROM to DATA */
  *dst++ = *src++;

for (dst = &_bss; dst< &_ebss; dst++) // Zero bss.
  *dst = 0;