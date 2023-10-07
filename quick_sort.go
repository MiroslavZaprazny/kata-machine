package main

import "fmt"

func quick_sort(arr *[8]int, lo int, hi int) {
    if lo >= hi {
        return
    }
    pivot := partition(arr, lo, hi)
    quick_sort(arr, lo, pivot - 1)
    quick_sort(arr, pivot + 1, hi)
}

func partition(arr *[8]int, lo int, hi int) int {
    pivot := arr[hi]
    idx := lo - 1;

    for i:= lo; i < hi; i++ {
        if arr[i] <= pivot {
            idx++
            tmp := arr[i]
            arr[i] = arr[idx]
            arr[idx] = tmp
        }
    }

    idx++
    arr[hi] = arr[idx]
    arr[idx] = pivot

    return idx
}

func main() {
    input := [8]int {9, 3, 7, 4, 69, 6, 420, 45}
    quick_sort(&input, 0, len(input) - 1)
    fmt.Println(input);
}
