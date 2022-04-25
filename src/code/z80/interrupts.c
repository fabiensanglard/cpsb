// FUNCTION to WAIT FOR YM2151 to be ready.





void interrupt() {
  // Copy latch values in circular buffer
}

void schedInterrupt() {
  // Schedule an interrupt in 4ms
  YM2151_REG_CMD = x;
  YM2151_REG_DAT = x;
}
