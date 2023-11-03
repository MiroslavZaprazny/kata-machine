package heap

type MinHeap struct {
    len int
    data []int
}

func (h *MinHeap) Insert(value int) {

}

func (h *MinHeap) Delete() int {

}

func (h *MinHeap) heapifyUp(idx int) {
    if idx == 0 {
        return
    }

    parent := h.parent(idx)
    parentValue := h.data[parent]
    currValue := h.data[idx]

    if currValue > parentValue {
        h.data[parent] = currValue
        h.data[idx] = parentValue
        h.heapifyUp(parent)
    }

}

func (h *MinHeap) parent(idx int) int {
    return (idx-1) / 2
}

func (h *MinHeap) leftChild(idx int) int {
    return idx * 2 + 1
}

func (h *MinHeap) rightChild(idx int) int {
    return idx * 2 + 2
}
