  Sprites are represented by a number of 8 byte values
  xx xx yy yy nn nn aa aa
  where xxxx = x position
        yyyy = y position
        nnnn = tile number
        aaaa = attribute word
                    0x0001  colour
                    0x0002  colour
                    0x0004  colour
                    0x0008  colour
                    0x0010  colour
                    0x0020  X Flip
                    0x0040  Y Flip
                    0x0080  X & Y offset toggle (used in Marvel vs. Capcom.)
                    0x0100  X block size (in sprites)
                    0x0200  X block size
                    0x0400  X block size
                    0x0800  X block size
                    0x1000  Y block size (in sprites)
                    0x2000  Y block size
                    0x4000  Y block size
                    0x8000  Y block size