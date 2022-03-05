__at(0xF000) char YM2151_REG_CMD;
__at(0xF001) char YM2151_REG_DAT;

void interrupt() {
  // Copy latch values in circular buffer
}

void schedInterrupt() {
  // Schedule an interrupt in 4ms
  YM2151_REG_CMD = x;
  YM2151_REG_DAT = x;
}
