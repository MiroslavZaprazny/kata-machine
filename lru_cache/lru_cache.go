package lru_cache

import "errors"

type node struct {
    value int64
    next *node
    prev *node
}

type LruCache struct {
    capacity int64
    len int64
    cache map[int64]*node
    lookup map[*node]int64
    head *node
    tail *node
}

func New(capacity int64) *LruCache {
    return &LruCache{capacity: capacity, cache: make(map[int64]*node), lookup: make(map[*node]int64)}
}

func (lu *LruCache) Put(key int64, value int64) {
    node, err := lu.Get(key)
    if err != nil {
        lu.insert(key, value)
        return
    }

    node.value = value
    lu.cache[key] = node
}

func (lu *LruCache) insert(key int64, value int64) {
    node := &node{value: value}
    lu.detach(node)
    lu.prepend(node)
    lu.trim()

    lu.len++
    lu.cache[key] = node
    lu.lookup[node] = key
}

func (lu *LruCache) detach(node *node) {
    if node.prev != nil {
        node.prev.next = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    }

    if node == lu.head {
       lu.head = lu.head.next 
    }

    if node == lu.tail {
       lu.tail = lu.tail.prev
    }
    
    node.prev = nil
    node.next = nil
}

func (lu *LruCache) prepend(node *node) {
    if lu.head == nil {
        lu.head = node
        lu.tail = node

        return
    }

    node.next = lu.head
    lu.head.prev = node
    lu.head = node
}

func (lu *LruCache) trim() {
    if lu.len >= lu.capacity {
        return
    }

    key := lu.lookup[lu.tail]
    delete(lu.cache, key)
    delete(lu.lookup, lu.tail)
    lu.detach(lu.tail)
}

func (lu *LruCache) Get(key int64) (*node, error) {
    node, ok := lu.cache[key]

    if !ok {
        return nil, errors.New("Key not found")
    }
    lu.detach(node)
    lu.prepend(node)

    return node, nil
}
