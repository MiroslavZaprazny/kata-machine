package quick_sort

func Quick_sort(arr *[8]int, lo int, hi int) {
    if lo >= hi {
        return
    }
    pivot := partition(arr, lo, hi)
    Quick_sort(arr, lo, pivot - 1)
    Quick_sort(arr, pivot + 1, hi)
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

