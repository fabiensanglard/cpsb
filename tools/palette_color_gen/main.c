#include <stdio.h>

int main()
{
    
    for (int b = 0; b <= 0xf ; b++) {
        int bright = 0x0f + (b  << 1);
        printf("%X = %X0000FF\n", b, 0x0f * 0x11 * bright / 0x2d);
    }
;

    return 0;
}