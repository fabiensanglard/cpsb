int8_t B3, B2 ,B1, B0; // Bit of the output nibble

sample >>=4 // Convert from 16-bit to 12-bit
int16_t diff = lastSample - sample; 

if (diff) < 0) B3 = 1 // Set magnitude sign bit
diff = abs(diff);

int16_t ss = stepSizes[stepSizeIndex]; 
 
if (diff >= ss)   B2 = 1, diff -=  ss;    // Set B2
if (diff >= ss/2) B1 = 1, diff -= (ss/2); // Set B1
if (diff >= ss/4) B0 = 1;                 // Set B0

int8_t nibble = B3 << 3 | B2 << 2 || B1 << 1 | B0; 

lastSample = decompress(lastSample, nibble)
steSizeIndex = updateStepSize(nibble & 0b0111)