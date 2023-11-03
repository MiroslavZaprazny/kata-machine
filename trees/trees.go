package trees

type BinaryNode struct {
    Value interface{}
    Right *BinaryNode
    Left *BinaryNode
}

func walk(node *BinaryNode, path *[]int64) *[]int64 {
    if node == nil {
        return path
    }

    val, ok := node.Value.(int)
    if !ok {
        panic("not a number")
    }
    *path = append(*path, int64(val))
    walk(node.Left, path)
    
    // in-order walk
    // *path = append(*path, int64(val))
    walk(node.Right, path)
    // post-order walk
    // *path = append(*path, int64(val))

    return path
}

func Pre_order_walk(head *BinaryNode) *[]int64 {
    return walk(head, &[]int64{})
}

func Bfs(head *BinaryNode, value interface{}) bool {
    //to make it more efficient use a proper queue not a slice
    q := []*BinaryNode {head}

    for len(q) != 0 {
        //shift
        curr := q[0]
        q = q[1:]
        
        if curr.Value == value {
            return true
        }

        if curr.Right != nil {
            q = append(q, curr.Right)
        }

        if curr.Left != nil {
            q = append(q, curr.Left)
        }
    }

    return false
}

func Dfs(head *BinaryNode, value interface{}) bool {
    if head == nil {
        return false
    }

    if head.Value == value {
        return true
    }

    if head.Value.(int) < value.(int) {
        return Dfs(head.Right, value)
    }

    return Dfs(head.Left, value)
}

func Compare(a *BinaryNode, b *BinaryNode) bool {
    if a == nil && b == nil {
        return true
    }
    if a == nil || b == nil {
        return false
    }
    if a.Value != b.Value {
        return false
    }

    return Compare(a.Left, b.Left) && Compare(a.Right, b.Right)
}
