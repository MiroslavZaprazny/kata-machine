type SNode<T> = {
    value: T,
    prev?: SNode<T>,
}

class Stack<T> {
    public length: number;
    public head?: SNode<T>; 

    construct() {
        this.head = undefined;
        this.length = 0;
    }

    push(item: T): void {
        this.length++;
        let node = {value: item} as SNode<T>;
        if (!this.head) {
            this.head = node;
            return;
        }

        node.prev = this.head;
        this.head = node;
    }

    pop(): T | undefined {
        if (!this.head) {
            return undefined;
        }

        let head = this.head;
        this.head = head.prev;

        return head.value;
    }

    peek(): T | undefined {
        return this.head?.value;
    }
}

let stack = new Stack<number>();

stack.push(5);
stack.push(10);

console.log(stack.peek());
console.log(stack.pop());

console.log(stack.peek());
console.log(stack.pop());
