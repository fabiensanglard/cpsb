volatile int lastFrameCounter = 0;

void hostFrame() {
  ... // Game engine render one visual and audio frame.
}

void main() {
  while(true) {
    if (frameCounter >= vsyncCounter) continue;
    hostFrame(); // Run 16ms of gameplay
    frameCounter++;
  }
}