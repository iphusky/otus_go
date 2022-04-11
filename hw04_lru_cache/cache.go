package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*cacheItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*cacheItem, capacity),
	}
}

func (i *lruCache) Get(key Key) (interface{}, bool) {
	val, ok := i.items[key]

	if ok {
		i.queue.MoveToFront(i.queue.GetItem(val))
		return val.value, true
	}

	return nil, false
}

func (i *lruCache) Set(key Key, value interface{}) bool {
	val, ok := i.items[key]

	if ok {
		val.value = value
		elem := i.queue.GetItem(val)
		i.queue.MoveToFront(elem)
		return true
	}

	item := itemFactory(key, value)
	i.queue.PushFront(item)
	i.items[key] = item

	if i.queue.Len() > i.capacity {
		last := i.queue.Back()
		i.queue.Remove(last)

		ci, ok := last.Value.(*cacheItem)

		if ok {
			delete(i.items, ci.key)
		}
	}

	return false
}

func (i *lruCache) Clear() {
	i.queue = NewList()
	i.items = make(map[Key]*cacheItem, i.capacity)
}

func itemFactory(key Key, value interface{}) *cacheItem {
	return &cacheItem{
		key:   key,
		value: value,
	}
}
