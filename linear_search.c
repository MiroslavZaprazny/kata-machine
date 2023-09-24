#include <stdio.h>
int linear_search(int arr[], int needle, int n)
{
    for (int i = 0; i <= n; i++) {
        if (arr[i] == needle) {
            return i;
        }
    }

    return -1;
}

int main(int argc, char *argv[])
{
    int arr[5] = {2, 1, 3, 8, 20};
    int n = (int) (sizeof(arr) / sizeof(arr[0]));
    printf("%d", linear_search(arr, 3, n));
    return 0;
}

