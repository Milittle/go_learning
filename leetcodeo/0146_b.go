package leetcodeo

type LRUCache struct {
	capacity   int
	Keys       map[int]*Node
	Head, Tail *Node
}

type Node struct {
	K, V      int
	pre, next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		Keys:     map[int]*Node{},
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.V
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.V = value
		this.Remove(node)
		this.Add(node)
	} else {
		node := &Node{
			K:    key,
			V:    value,
			pre:  nil,
			next: nil,
		}
		this.Keys[key] = node
		this.Add(node)
	}

	if len(this.Keys) > this.capacity {
		delete(this.Keys, this.Tail.K)
		this.Remove(this.Tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.pre = nil
	node.next = this.Head
	if this.Head != nil {
		this.Head.pre = node
	}
	this.Head = node

	if this.Tail == nil {
		this.Tail = node
		//this.Tail.next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.Head {
		this.Head = node.next
		//if node.next != nil {
		//	node.next.pre = nil
		//}
		//node.next = nil
		//return
	}
	if node == this.Tail {
		this.Tail = node.pre
		//node.pre.next = nil
		//node.pre = nil
		return
	}
	node.pre.next = node.next
	node.next.pre = node.pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
