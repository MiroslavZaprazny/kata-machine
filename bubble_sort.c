#include <stdio.h>

void bubble_sort(int arr[], int n) {
    for (int i = 0; i < n; i++) {
        for (int j = 0; j <= n-1-i; j++) {
            if (arr[j] >= arr[j+1]) {
                int tmp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = tmp;
            }
        }
    }

    for (int i = 0; i < n; i++) {
        printf("%d\n", arr[i]);
    }
}

int main(int argc, char *argv[])
{
    int arr[5] = {5, 8, 1, 2, 10};
    int n = sizeof(arr) / sizeof(arr[0]);
    bubble_sort(arr, n);

    return 0;
}
