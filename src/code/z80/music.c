int8_t noopCounter = 0;

void updateMusic() {
  if (noopCounter) {
    noopCounter--;
    return;
  }

  int8_t bc = next();
  switch (bc) {
    case DELAY : 
       noopCounter = bc; break;
    case MUSIC_NOTE : 
       REG_YM2151_CMD = next(); 
       REG_YM2151_DAT = next(); 
       break;
    case MUSIC_SOUND : 
       REG_OKI = 0x8 | next();
       REG_OKI = next();
       break;
    case 0: // End of Song
       break;   
  }
}