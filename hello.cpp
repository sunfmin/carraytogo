#include <stdint.h>
#include <stdio.h>

extern "C" {
    void fillArray1(uint8_t *ret);
    void fillArray2(uint8_t **ret);
}


void fillArray1(uint8_t *ret)
{
    printf("fillArray1\n");

    ret[0] = 2;
    ret[1] = 3;

    return;
}

void fillArray2(uint8_t **ret)
{
    printf("fillArray2\n");

    // uint8_t* value = (uint8_t*) 0x0004;
    // printf("%s", value);

    uint8_t e1[] = {1, 2};
    uint8_t e2[] = {3, 4};

    ret[0] = e1;
    ret[1] = e2;

    return;
}

