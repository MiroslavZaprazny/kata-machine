package linked_list

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
    Head *Node
    Tail *Node
}

func (li *LinkedList) Prepend(item interface{}) {
    node := &Node{Value: item}
    li.length++

    if li.Head == nil {
        li.Head = node
        li.Tail = node
        return
    }
    
    node.Next = li.Head
    li.Head.Prev = node
    li.Head = node
}

func (li *LinkedList) Append(item interface{}) {
    li.length++
    node := &Node{Value: item}
    if li.Tail == nil {
        li.Tail = node
        li.Head = node
        return
    }

    node.Prev = li.Tail
    li.Tail.Next = node
    li.Tail = node
}

func (li *LinkedList) InsertAt(item interface{}, idx int) error {
    if idx > li.length {
        return errors.New(fmt.Sprintf("Length lower %d then idx %d", li.length, idx))
    } else if idx == li.length && idx == 0 {
        li.Append(item)
        return nil
    }

    curr := li.Head
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
    curr := li.Head
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
    curr := li.Head
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
        li.Head = nil
        li.Tail = nil

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

