#include <math.h>
#include <stdio.h>

int binary_search(int arr[], int needle, int n)
{
    int lo = 0;
    int hi = n;

    do {
        int m = floor((hi + lo) / 2);
        if (arr[m] == needle) {
            return m;
        } else if(arr[m] > needle) {
            hi = m - 1;
        } else {
            lo = m + 1;
        }
    } while (lo <= hi);

    return -1;
}

int main(int argc, char *argv[])
{
    int arr[10] = {1, 2, 3, 4, 5, 6, 7, 8, 20, 30};
    int n = sizeof(arr) / sizeof(arr[0]);

    printf("%d", binary_search(arr, 0, n));

    return 0;
}
