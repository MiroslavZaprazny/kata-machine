package lru_cache

import "errors"

type node struct {
    value int64
    next *node
    prev *node
}

type LruCache struct {
    capacity int64
    cache map[int64]*node
    head *node
    tail *node
}

func New(capacity int64) *LruCache {
    return &LruCache{capacity: capacity, cache: make(map[int64]*node)}
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
    //todo check for overcapacity
    node := &node{value: value}
    if lu.head == nil {
        lu.head = node
        lu.tail = node
    } else {
        lu.tail.next = node
        node.prev = lu.tail
        lu.tail = node
    }

    lu.cache[key] = node
}

func (lu *LruCache) Get(key int64) (*node, error) {
    node, ok := lu.cache[key]

    if !ok {
        return nil, errors.New("Key not found")
    }
    node.next = lu.head
    lu.head.prev = node
    lu.head = node

    return node, nil
}
