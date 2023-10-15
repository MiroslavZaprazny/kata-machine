// package main

import (
	"errors"
	"fmt"
)

type Node struct {
    Value interface{}
    Next *Node
    Prev *Node
}

type LinkedList struct {
    length int
    head *Node
    tail *Node
}

func (li *LinkedList) Prepend(item interface{}) {
    node := &Node{Value: item}
    li.length++

    if li.head == nil {
        li.head = node
        li.tail = node
        return
    }
    
    node.Next = li.head
    li.head.Prev = node
    li.head = node
}

func (li *LinkedList) Append(item interface{}) {
    li.length++
    node := &Node{Value: item}
    if li.tail == nil {
        li.tail = node
        li.head = node
        return
    }

    node.Prev = li.tail
    li.tail.Next = node
    li.tail = node
}

func (li *LinkedList) InsertAt(item interface{}, idx int) error {
    if idx > li.length {
        return errors.New(fmt.Sprintf("Length lower %d then idx %d", li.length, idx))
    } else if idx == li.length && idx == 0 {
        li.Append(item)
        return nil
    }

    curr := li.head
    for i := 0; i < idx; i++ {
        curr = curr.Next
    } 

    li.length++
    node := &Node{Value: item}

    node.Next = curr
    node.Prev = curr.Prev
    curr.Prev = node
    node.Prev.Next = node

    return nil
}

func (li *LinkedList) Remove(item interface{}) (interface{}, error) {
    curr := li.head
    for i := 0; i < li.length; i++ {
        if curr.Value == item {
            break;
        }
        curr = curr.Next
    }

    if curr == nil {
        return nil, errors.New("Item not found")
    }

    node := li.removeNode(curr)

    return node.Value, nil
}

func (li *LinkedList) Get(idx int) (interface{}, error) {
    curr, err := li.getAt(idx)
    if err != nil {
        panic(err.Error())
    }

    return curr.Value, nil
}

func (li *LinkedList) RemoveAt(idx int) (interface{}, error) {
    curr, err := li.getAt(idx)
    if err != nil {
        return nil, err
    }
    node := li.removeNode(curr)

    return node.Value, nil
}

func (li *LinkedList) getAt(idx int) (*Node, error) {
    curr := li.head
    for i := 0; i < idx; i ++ {
        curr = curr.Next
    }
    if curr == nil {
       return nil, errors.New(fmt.Sprintf("Nothing was found at index %d", idx))
    }

    return curr, nil
}

func (li *LinkedList) removeNode(node *Node) *Node {
    li.length--
    if li.length == 0 {
        li.head = nil
        li.tail = nil

        return node
    }

    if node.Prev.Next != nil {
        node.Prev.Next = node.Next
    }
    if node.Next.Prev != nil {
        node.Next.Prev = node.Prev
    }

    node.Next = nil
    node.Prev = nil

    return node
}

func main() {
    list := LinkedList{}
    list.Prepend(5)
    list.Prepend(10)
    list.InsertAt(1, 1)
    fmt.Printf("%+v\n", list)
    fmt.Printf("%+v\n", list.head.Value)
    fmt.Printf("%+v\n", list.tail.Value)
    val, err := list.Get(1)
    if err != nil {
        panic(err)
    }
    fmt.Print(val)
}

