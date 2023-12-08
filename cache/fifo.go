package cache

// Notes:
// empty string = valid key;
// empty slice and NIL = valid items;

// An FIFO is a fixed-size in-memory cache with first-in first-out eviction
type FIFO struct {
	// whatever fields you want here
	maxBytes  int // max storage of cache/queue
	currBytes int
	queue     []string          // queue containing keys
	keyValMap map[string][]byte // map containing key, val pairs -> for fast lookup of val given key
	hits      int
	misses    int
}

// NewFIFO returns a pointer to a new FIFO with a capacity to store limit bytes
func NewFifo(limit int) *FIFO {
	fifo := new(FIFO)
	fifo.maxBytes = limit
	fifo.currBytes = 0
	fifo.queue = make([]string, 0)
	fifo.keyValMap = make(map[string][]byte)
	return fifo
}

// MaxStorage returns the maximum number of bytes this FIFO can store
func (fifo *FIFO) MaxStorage() int {
	return fifo.maxBytes
}

// RemainingStorage returns the number of unused bytes available in this FIFO
func (fifo *FIFO) RemainingStorage() int {
	return fifo.maxBytes - fifo.currBytes
}

// Get returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise.
func (fifo *FIFO) Get(key string) (value []byte, ok bool) {
	value, ok = fifo.keyValMap[key]
	if ok {
		fifo.hits += 1
	} else {
		fifo.misses += 1
	}
	return value, ok
}

// Remove removes and returns the value associated with the given key, if it exists.
// ok is true if a value was found and false otherwise
func (fifo *FIFO) Remove(key string) (value []byte, ok bool) {
	ok = false
	// remove from queue -- iterate through queue to find index
	for i := 0; i < len(fifo.queue); i++ {
		if fifo.queue[i] == key {
			ok = true
			fifo.queue = append(fifo.queue[:i], fifo.queue[i+1:]...)
			break
		}
	}
	if !ok {
		return nil, false
	}
	// remove from map
	value = fifo.keyValMap[key]
	delete(fifo.keyValMap, key)

	// update curr storage
	keyValBytes := len(key) + len(value)
	fifo.currBytes -= keyValBytes

	return value, ok
}

// Set associates the given value with the given key, possibly evicting values
// to make room. Returns true if the binding was added successfully, else false.
func (fifo *FIFO) Set(key string, value []byte) bool {
	item_size := len(key) + len(value)
	// check if it is impossible to store item into cache
	if item_size > fifo.maxBytes {
		return false
	}

	// check if key already exists
	prev_val, exists := fifo.keyValMap[key]

	if exists {
		// subtract original size
		prev_size := len(key) + len(prev_val)
		fifo.currBytes -= prev_size

		delete(fifo.keyValMap, key)
	}

	// check if there's room
	for item_size > fifo.RemainingStorage() {
		// update curr size
		evictedKey := fifo.queue[0]
		evictedVal := fifo.keyValMap[evictedKey]
		evictedSize := len(evictedKey) + len(evictedVal)
		fifo.currBytes -= evictedSize

		// removes item
		fifo.queue = fifo.queue[1:]
		delete(fifo.keyValMap, evictedKey)
	}
	// add key
	fifo.keyValMap[key] = value
	if !exists {
		fifo.queue = append(fifo.queue, key)
	}

	fifo.currBytes += item_size

	return true
}

// Len returns the number of bindings in the FIFO.
func (fifo *FIFO) Len() int {
	return len(fifo.queue)
}

// Stats returns statistics about how many search hits and misses have occurred.
func (fifo *FIFO) Stats() *Stats {
	return &Stats{fifo.hits, fifo.misses}
}
