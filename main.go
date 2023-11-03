package main

import (
	"fmt"
	"kata-machine/linked_list"
	"kata-machine/quick_sort"
	"kata-machine/trees"
)

func main() {
    tree := &trees.BinaryNode{
        Value: 20,
        Right: &trees.BinaryNode{
            Value: 50,
            Right: &trees.BinaryNode{
                Value: 30,
                Right: nil,
                Left: nil,
            },
            Left: &trees.BinaryNode{
                Value: 100,
                Right: nil,
                Left: nil,
            },
        },
        Left: &trees.BinaryNode{
            Value: 200,
            Right: &trees.BinaryNode{
                Value: 300,
                Right: nil,
                Left: nil,
            },
            Left: &trees.BinaryNode{
                Value: 600,
                Right: nil,
                Left: nil,
            },
        },
    }

    tree2 := &trees.BinaryNode{
        Value: 20,
        Right: &trees.BinaryNode{
            Value: 50,
            Right: &trees.BinaryNode{
                Value: 30,
                Right: nil,
                Left: nil,
            },
            Left: &trees.BinaryNode{
                Value: 100,
                Right: nil,
                Left: nil,
            },
        },
        Left: &trees.BinaryNode{
            Value: 200,
            Right: &trees.BinaryNode{
                Value: 300,
                Right: nil,
                Left: nil,
            },
            Left: &trees.BinaryNode{
                Value: 600,
                Right: nil,
                Left: nil,
            },
        },
    }

    path := trees.Pre_order_walk(tree)
    fmt.Println(trees.Bfs(tree, 200))
    fmt.Println(path)
    fmt.Println(trees.Compare(tree, tree2))

    input := [8]int {9, 3, 7, 4, 69, 6, 420, 45}
    quick_sort.Quick_sort(&input, 0, len(input) - 1)
    fmt.Println(input);

    list := linked_list.LinkedList{}
    list.Prepend(5)
    list.Prepend(10)
    list.InsertAt(1, 1)
    fmt.Printf("%+v\n", list)
    fmt.Printf("%+v\n", list.Head.Value)
    fmt.Printf("%+v\n", list.Tail.Value)
    val, err := list.Get(1)
    if err != nil {
        panic(err)
    }
    fmt.Print(val)
}
