volatile int vsyncCounter = 0;
volatile int frameCounter = 0;

// Called every 16ms
void VSync() {

  if (frameCounter != lastFrameCounter) {
    flipGFXRAMPointers(); // Flip GFX SCROLLs and OBJs.
    writeSoundLatch();    // dequeue and write latch
    readInputs();
    lastCounter = frameCounter;
  }

  vsyncCounter++; // Unlock the main loop.
}