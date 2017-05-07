// Copyright ® 2017 Ian G. Tayler <ian.g.tayler@gmail.com>
// Distribute according to the LICENSE.
#include <stdint.h>

int SimplRecFib (int n)
{
    if (n < 2) {
        return 1;
    }
    return SimplRecFib(n-1) + SimplRecFib(n-2);
}

uint32_t StdintRecFib (int n)
{
    if (n < 2) {
        return 1;
    } 
    return StdintRecFib(n-1) + StdintRecFib(n-2);
}