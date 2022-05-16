int8_t noopCounter = 0;

void updateMusic() {
  if (noopCounter) {
    noopCounter--;
    return;
  }

  int8_t bc = next();
  switch (bc) {
    case XX : noopCounter = bc; break;
    case YY : 
       REG_YM2151_CMD = next(); 
       REG_YM2151_DAT = next(); 
       break;
    case ZZ : 
       REG_OKI = 0x8 | next();
       REG_OKI = next();
  }
}