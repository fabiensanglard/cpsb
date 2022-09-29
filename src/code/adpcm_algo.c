int transitionTable[8] = {-1, -1, -1, -1, 2, 4, 6, 8};

int stepSizes[49] = {
  16,   17,   19,   21,   23,   25,   28,   31,   34,   37,  
  41,   45,   50,   55,   60,   66,   73,   80,   88,   97,  
 107,  118,  130,  143,  157,  173,  190,  209,  230,  253,  
 279,  307,  337,  371,  408,  449,  494,  544,  598,  658,  
 724,  796,  876,  963, 1060, 1166, 1282, 1411, 1552};

int stepSizeIndex = 0; // Initial value (0) points to 16
int16_t lastSample = 0;

int8_t compress(int16_t sample) {
  int8_t B3 = 0, B2 = 0, B1 = 0, B0 = 0; // Bit of the output nibble

  sample >>= 4; // Convert from 16-bit to 12-bit
  int16_t diff = sample - lastSample; 

  if (diff < 0) B3 = 1; // Set magnitude sign bit
  diff = abs(diff);

  int16_t ss = stepSizes[stepSizeIndex]; 
 
  if (diff >= ss)   B2 = 1, diff -=  ss;    // Set B2
  if (diff >= ss/2) B1 = 1, diff -= (ss/2); // Set B1
  if (diff >= ss/4) B0 = 1;                 // Set B0

  int8_t nibble = B3 << 3 | B2 << 2 | B1 << 1 | B0; 

  // Keep track of the value upon decompression
  lastSample = decompress(lastSample, nibble, stepSizeIndex); 

  stepSizeIndex += transitionTable[nibble & 0x7];
  return nibble;
}
