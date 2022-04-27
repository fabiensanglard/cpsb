int vsyncCounter = 0;

void VySnc() {
  vsyncCounter++;
  flipGFXRAMPointers();
  writeToSoundLatch();
  readInputs();
}

int frameCounter == 0;

void hostFrame() {
  frameCounter++;
  // Game engine render one visual frame.
}

void main() {
  while(true) {
    if (frameCounter < vsyncCounter) continue;
    hostFrame();
  }
}