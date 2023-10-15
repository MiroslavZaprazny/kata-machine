package main

import "fmt"

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

func pre_order_walk(head *BinaryNode) *[]int64 {
    return walk(head, &[]int64{})
}

func bfs(head *BinaryNode, value interface{}) bool {
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

func main() {
    tree := &BinaryNode{
        Value: 20,
        Right: &BinaryNode{
            Value: 50,
            Right: &BinaryNode{
                Value: 30,
                Right: nil,
                Left: nil,
            },
            Left: &BinaryNode{
                Value: 100,
                Right: nil,
                Left: nil,
            },
        },
        Left: &BinaryNode{
            Value: 200,
            Right: &BinaryNode{
                Value: 300,
                Right: nil,
                Left: nil,
            },
            Left: &BinaryNode{
                Value: 600,
                Right: nil,
                Left: nil,
            },
        },
    }

    path := pre_order_walk(tree)
    fmt.Println(bfs(tree, 200))

    fmt.Print(path)
}
