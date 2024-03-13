package leetcode

type DLink struct {
	prev  *DLink
	next  *DLink
	value int
	key   int
}

// create dummy head
func createDlink() *DLink {
	return &DLink{
		prev:  nil,
		next:  nil,
		value: -1,
	}
}

func insertHead(head, node *DLink) *DLink {
	if head == nil || node == nil {
		return nil
	}
	node.next = head.next
	node.prev = head
	if head.next != nil {
		head.next.prev = node
	}
	head.next = node
	return node
}

func popTail(head *DLink) *DLink {
	if head.next == nil {
		return nil
	}
	p := head.next
	for p.next != nil {
		p = p.next
	}
	p.prev.next = p.next
	p.prev = nil
	return p
}

func removeNode(node *DLink) *DLink {
	if node == nil {
		return node
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = nil
	return node
}

type LRUCache struct {
	lru  *DLink
	data map[int]*DLink
	cap  int
}

func Constructor(capacity int) LRUCache {
	_data := make(map[int]*DLink, capacity)
	return LRUCache{
		cap: capacity,
		lru: &DLink{
			prev:  nil,
			next:  nil,
			value: -1,
			key:   -1,
		},
		data: _data,
	}
}

func (this *LRUCache) Get(key int) int {
	if link, ok := this.data[key]; ok {
		_l := removeNode(link)
		insertHead(this.lru, _l)
		return link.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if link, ok := this.data[key]; ok {
		_l := removeNode(link)
		_l.value = value
		insertHead(this.lru, _l)
	} else {
		if len(this.data) >= this.cap {
			_l := popTail(this.lru)
			if _l != nil {
				delete(this.data, _l.key)
			}
		}
		_h := insertHead(this.lru, &DLink{
			prev:  nil,
			next:  nil,
			value: value,
			key:   key,
		})
		this.data[key] = _h
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
