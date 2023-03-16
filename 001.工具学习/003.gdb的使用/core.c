#include <stdio.h>

int div(int total, int piece)
{
    return total / piece;
}

int main()
{
    int total = {100};
    for (int i = 0; i < 1000; i++)
    {
        // 当 i=100 时，函数内除数为0，将引发异常退出，从而产生 coredump
        printf("total: %d, piece: %d, div: %d", total, i, div(total, i));
    }
    return 0;
}