// List of VGM bytecode to convert to miniVGM bytecode.
0x54 | aa dd | YM2151, write value dd to register aa
0x61 | nn nn | Wait n samples (0 to 1.49 seconds).
0x62 |       | Wait 735 samples (60th of a second)
0x63 |       | Wait 882 samples (50th of a second),
0x66 |       | End of sound data

0xC0 |bbaa dd| Sega PCM, write value dd to memory offset aabb.