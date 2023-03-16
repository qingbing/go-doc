#include <stdio.h>
#include <unistd.h>

int func(int n)
{
    int sum = 0, i;
    for (i = 0; i < n; i++)
    {
        sum += i;
    }
    return sum;
}

int main()
{
    printf("Program START\n");
    for (int i = 1; i <= 1000; i++)
    {
        printf("result[%d] = %dn\n", i, func(i));
        sleep(1);
    }
    return 0;
}