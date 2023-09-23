#include <stdio.h>
int linear_search(int arr[], int value) {
    int len = (int) (sizeof(arr) / sizeof(arr[0]));
    for (int i = 0; i <= len; i++) {
        if (arr[i] == value) {
            return i;
        }
    }

    return -1;
}

int main(int argc, char *argv[])
{
    int arr[4] = {1, 2, 3, 4};
    printf("%d", linear_search(arr, 3));
    return 0;
}

