int8_t noopCounter = 0;

void updateMusic() {
  if (noopCounter) { // Does the YM2151 need a break?
    noopCounter--;   // 14ms increment.
    return;
  }
next_byte_code:
  int8_t bc = next();
  switch (bc) { // Use our custom bytecode (miniVGM).
    case DELAY : 
       noopCounter = bc; break;
    case MUSIC_NOTE : 
       REG_YM2151_CMD = next(); 
       REG_YM2151_DAT = next(); 
       goto next_byte_code;
    case MUSIC_SOUND : 
       REG_OKI = 0x8 | next();
       REG_OKI = next();
       goto next_byte_code;
    case END_SONG: // End of song
       stopMusic(); break;   
  }
}