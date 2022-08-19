int8_t intCounter = 0;
int8_t latch;


void interrupt () {
    intCounter++;
    latch = REG_LATCH1;
}

// Incremtend by the main() function.
int8_t  musCounter = 0;

// Incremented by the interrupt() function.
volatile int8_t  lastLatch;


void main () {
  while(true) {
    musCounter++;
    
    // Only tick after interrupt ticks
    while (musCounter < intCounter) {
    }

    updateMusic(); // Feed the YM2151 or skip for pause.
	
    if (latch == 0xFF) continue;
    if (latch == lastLatch) continue;
    lastLatch = latch;

    // Forward to OKI
    if (latch & 0x80) {
      REG_OKI = 0x8  | latch;
      // 0x10 = Channel 1, 0x00 = Max volume.
      REG_OKI = 0x10 | 0x00; // TODO: Round-robin on channels.
    } else {
      setupMusicPlayback(latch & 0x70);
    }
  }
} 