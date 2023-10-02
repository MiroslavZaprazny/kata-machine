struct node {
    void* value;
    struct node* next;
}node;

struct queue {
   struct node* head;
   struct node* tail;
   int length;
}queue;

void enqueue(struct queue queue, void* value) {
    struct node* node = {value};
    if (queue.length == 0) {
        queue.head = node;
        queue.tail = node;
        return;
    }

    struct node* tail = queue.tail;
    queue.tail->next = node;
    queue.tail = node;
    queue.length++;
}

void* deque(struct queue queue) {
    if (queue.length == 0) {
        return 0;
    }

    struct node* head = queue.head;
    queue.head = queue.head->value;
    queue.length--;

    return head->value;
}

void* peek(struct queue queue) {
    return queue.head->value;
}

int main(int argc, char *argv[])
{
    return 0;
}
