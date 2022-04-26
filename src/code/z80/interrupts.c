void waitYM2151() {
  while (REG_YM2151_DAT == 0x80) {
    // Wait until YM2151 is ready for write
  }
}

void interrupt() {
  // Read latches here
}

void schedInterrupt() {
  // Schedule an interrupt in 4ms
  waitYM2151 ();
  REG_YM2151_CMD = 0xC8;
  REG_YM2151_DAT = 0x00;
}
