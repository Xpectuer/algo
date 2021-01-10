package LRUCache

type ListNode struct {
	val  int
	next *ListNode
	prev *ListNode
	key  int
}

type LRUCache struct {
	head     *ListNode
	tail     *ListNode
	hashmap  map[int]*ListNode
	capacity int
	length   int
}

func Constructor(capacity int) LRUCache {

	var (
		head *ListNode
		tail *ListNode
	)
	// memberwise initiate
	cap := capacity
	// initiate hashmap
	hashmap := make(map[int]*ListNode, capacity)
	// initiate LinkedList
	head = &ListNode{val: -1, prev: nil, next: nil}
	tail = &ListNode{val: -1, prev: nil, next: nil}
	head.next = tail
	tail.prev = head

	return LRUCache{hashmap: hashmap, tail: tail, head: head, capacity: cap, length: 0}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.hashmap[key]; ok {
		// remove
		node := v
		node.prev.next = node.next
		node.next.prev = node.prev
		// set to first
		this.addFirst(node)
		// return the value
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// check if the key exists
	if this.Get(key) != -1 {
		// modify the val
		this.hashmap[key].val = value
		return
	}
	// check if the cap is reached
	if this.tail.prev.prev != nil && this.length == this.capacity {
		// remove the last in the LinkedList
		node := this.tail.prev
		this.tail.prev.prev.next = this.tail
		this.tail.prev = this.tail.prev.prev
		//call GC
		node.next = nil
		node.prev = nil
		// remove the last in the hashmap
		delete(this.hashmap, node.key)
		this.length--
	}
	// insert into first
	node := &ListNode{prev: nil, next: nil, val: value, key: key}
	// reusable
	this.addFirst(node)

	this.hashmap[key] = node
	// maintain the length
	this.length++
}

func (this *LRUCache) addFirst(node *ListNode) {
	if node != nil {
		node.next = this.head.next
		//  fmt.Println(this.head.next)
		this.head.next.prev = node
		this.head.next = node
		node.prev = this.head
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
