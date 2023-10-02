type QNode<T> = {
    value: T,
    next?: QNode<T>,
}

class Queue<T> {
    public length: number;
    public head?: QNode<T>;
    public tail?: QNode<T>;

    constructor() {
        this.head = this.tail = undefined;
        this.length = 0; 
    }

    enqueue(item: T): void {
        this.length++;
        let node = {value: item} as QNode<T>;
        if (!this.tail) {
            this.tail = this.head = node;
            return;
        }

        this.tail.next = node;
        this.tail = node;
    }

    deque(): T | undefined {
        if (!this.head) {
            return undefined;
        }

        let head = this.head;
        this.head = this.head.next;
        this.length--;

        return head.value;
    }

    peek(): T | undefined {
        return this.head?.value;
    }
}

const queue = new Queue<number>();
queue.enqueue(5);
queue.enqueue(6);
queue.enqueue(8);
console.log(queue.peek());
console.log(queue.deque());
console.log(queue.length);

console.log(queue.peek());
console.log(queue.deque());
console.log(queue.length);

console.log(queue.peek());
console.log(queue.deque());
console.log(queue.length);

console.log(queue.peek());
console.log(queue.deque());
console.log(queue.length);
