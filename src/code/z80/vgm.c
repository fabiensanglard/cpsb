// List of VGM bytecode converted to miniVGM bytecode.
0x54 | aa dd | Write dd to YM2151 register aa.     -> MUSIC_NOTE
0x61 | nn nn | Wait n samples (0 to 1.49 seconds). -> DELAY
0x62 |       | Wait 735 samples (60th of a second) -> DELAY
0x63 |       | Wait 882 samples (50th of a second).-> DELAY
0x66 |       | End of sound data.                  -> END_SONG
0xC0 |bbaa dd| Write dd to Sega PCM register aabb. -> MUSIC_SOUND