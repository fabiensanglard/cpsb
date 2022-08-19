void waitYM2151() {
  while (REG_YM2151_DAT == 0x80) {
    // Wait until YM2151 is ready for write
  }
}

void interrupt() {
  // Read latches here
}

void schedInterrupt() { // Schedule an interrupt in 4ms
  waitYM2151 ();
  REG_YM2151_ADR = 0x10; // Register Timer A 8 MSB
  REG_YM2151_DAT = 0xC8; // 0b11001000

  waitYM2151 ();
  REG_YM2151_ADR = 0x11; // Register Timer A 2 LSB
  REG_YM2151_DAT = 0x00; // 0b00
}
