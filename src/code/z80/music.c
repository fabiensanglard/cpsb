int8_t pauseCounter = 0;

void updateMusic() {
  if (pauseCounter) {
  	pauseCounter--;
  	return;
  }

  int8_t bc = next();
  switch (bc) {
    case XX : pauseCounter = bc; break;
    case YY : 
  	   REG_YM2151_CMD = next(); 
  	   REG_YM2151_DAT = next(); 
  	   break;
    case ZZ : 
       REG_OKI = 0x8 | next();
       REG_OKI = next();
  }
}