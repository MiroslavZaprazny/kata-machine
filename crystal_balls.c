#include <math.h>
#include <stdio.h>

int crysal_balls(int arr[], int n) {
    int jump = floor(sqrt(n));
    int i = 0;

    for(; i < n; i+=jump) {
        if (arr[i] == 1) {
            break;
        }
    }

    i -= jump;

    for(int j = 0; j < jump && i < n; j++, i++) {
        if (arr[i] == 1) {
            return i;
        }
    }

    return -1;
}

int main(int argc, char *argv[])
{
    int arr[17] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1};
    int n = (int) (sizeof(arr) / sizeof(arr[0]));
    printf("%d", crysal_balls(arr, n));

    return 0;
}
