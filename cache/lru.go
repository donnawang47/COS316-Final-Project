package cache

// front 2 5 6 7  9 10 8 end

type Node struct {
	key   string
	value []byte
	prev  *Node
	next  *Node
}

type LinkedList struct {
	first *Node // least recent
	last  *Node // most recent
}

func (linkedList *LinkedList) AddToEnd(node *Node) {
	// linkedList.size += 1
	if linkedList.first == nil {
		linkedList.first = node
		linkedList.last = node
		return
	}
	lastNode := linkedList.last
	lastNode.next = node
	node.prev = lastNode
	node.next = nil
	linkedList.last = node
}

func (linkedList *LinkedList) RemoveNode(node *Node) {

	// If node to be deleted is head node
	if linkedList.first == node {
		linkedList.first = node.next
	}

	// 1 3
	// Change next only if node to be deleted
	// is NOT the last node
	if node.next != nil {
		node.next.prev = node.prev
	}

	// Change prev only if node to be deleted
	// is NOT the first node
	if node.prev != nil {
		node.prev.next = node.next
	}

	// Finally, free the memory occupied by del
	return

}

// An LRU is a fixed-size in-memory cache with least-recently-used eviction
type LRU struct {
	// whatever fields you want here
	maxBytes  int // max storage of cache/queue
	currBytes int
	hits      int
	misses    int
	nodeList  *LinkedList
	nodeMap   map[string]*Node
}

// NewLRU returns a pointer to a new LRU with a capacity to store limit bytes
func NewLru(limit int) *LRU {
	lru := new(LRU)
	lru.maxBytes = limit
	lru.currBytes = 0
	lru.hits = 0
	lru.misses = 0
	lru.nodeList = new(LinkedList)
	// lru.nodeList.size = 0
	lru.nodeMap = make(map[string]*Node)
	return lru
}

// MaxStorage returns the maximum number of bytes this LRU can store
func (lru *LRU) MaxStorage() int {
	return lru.maxBytes
}

// RemainingStorage returns the number of unused bytes available in this LRU
func (lru *LRU) RemainingStorage() int {
	return lru.maxBytes - lru.currBytes
}

// Get returns the value associated with the given key, if it exists.
// This operation counts as a "use" for that key-value pair
// ok is true if a value was found and false otherwise.
func (lru *LRU) Get(key string) (value []byte, ok bool) {

	// retrieve from map
	node, ok := lru.nodeMap[key]

	// update stats
	if ok {
		lru.hits += 1
	} else {
		lru.misses += 1
		return nil, false
	}

	// remove node from curr position
	lru.nodeList.RemoveNode(node)

	// move node to end (most recent)
	lru.nodeList.AddToEnd(node)

	return node.value, ok
}

// Remove removes and returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise
func (lru *LRU) Remove(key string) (value []byte, ok bool) {
	ok = false

	// get map
	node, ok := lru.nodeMap[key]

	if !ok {
		return nil, false
	}

	lru.nodeList.RemoveNode(node)

	// remove node from map
	delete(lru.nodeMap, key)

	// update curr storage
	nodeBytes := len(node.key) + len(node.value)
	lru.currBytes -= nodeBytes

	return node.value, ok
}

// 1

// Set associates the given value with the given key, possibly evicting values
// to make room. Returns true if the binding was added successfully, else false.
func (lru *LRU) Set(key string, value []byte) bool {

	item_size := len(key) + len(value)

	// not possible to cache current item even with eviction
	if item_size > lru.maxBytes {
		return false
	}

	// check if key already exists
	node, ok := lru.nodeMap[key]

	if ok {
		// subtract original size
		prev_size := len(node.key) + len(node.value)
		lru.currBytes -= prev_size
		// update value inside node
		node.value = value
		// remove from map
		delete(lru.nodeMap, key)
		// remove from linkedlist
		lru.nodeList.RemoveNode(node)

	} else {
		// create node with given key and val
		node = new(Node)
		node.key = key
		node.value = value
	}

	// Evict LRU nodes as needed to make for space for current item
	for item_size > lru.RemainingStorage() && len(lru.nodeMap) > 0 {
		evicted_node := lru.nodeList.first
		lru.nodeList.RemoveNode(evicted_node)

		// update curr size
		evicted_size := len(evicted_node.key) + len(evicted_node.value)
		lru.currBytes -= evicted_size
		// lru.nodeList.size -= 1

		// remove from map
		delete(lru.nodeMap, evicted_node.key)
	}

	// add to map
	lru.nodeMap[key] = node

	// add to linkedlist
	lru.nodeList.AddToEnd(node)

	lru.currBytes += item_size

	return true
}

// Len returns the number of bindings in the LRU.
func (lru *LRU) Len() int {
	return len(lru.nodeMap)
}

// Stats returns statistics about how many search hits and misses have occurred.
func (lru *LRU) Stats() *Stats {
	return &Stats{lru.hits, lru.misses}
}
