int vsyncCounter = 0;

void VSync() {
  vsyncCounter++;
  flipGFXRAMPointers();
  writeToSoundLatch();
  readInputs();
}

int frameCounter == 0;
