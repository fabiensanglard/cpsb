char *YM2151_REG_CMD = (char*)0xF000;
char *YM2151_REG_DAT = (char*)0xF001;

void interrupt() {
  // Copy latch values in circular buffer
}

void schedInterrupt() {
  // Schedule an interrupt in 4ms
  *YM2151_REG_CMD = x;
  *YM2151_REG_DAT = x;
}
