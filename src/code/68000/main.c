void hostFrame() {
  frameCounter++;
  ... // Game engine render one visual frame.
}

void main() {
  while(true) {
    if (frameCounter < vsyncCounter) continue;
    hostFrame();
  }
}