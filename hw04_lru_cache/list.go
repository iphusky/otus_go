package hw04lrucache

var cacheSize int = 10

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List
	items []*ListItem
}

type LRUCache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

func (l *list) Len() int {
	return len(l.items)
}

func (l *list) Front() *ListItem {
	if len(l.items) > 0 {
		return l.items[0]
	}
	return nil
}

func (l *list) Back() *ListItem {
	if len(l.items) > 0 {
		return l.items[len(l.items)-1]
	}
	return nil
}

func (l *list) PushFront(v interface{}) *ListItem {

	item := &ListItem{}
	item.Value = v

	if len(l.items) == 0 {
		l.items = append(l.items, item)
		return item
	}

	head := l.items[0]
	head.Prev = item
	item.Next = head

	newItems := make([]*ListItem, len(l.items)+1)
	newItems[0] = item
	copy(newItems[1:], l.items)
	l.items = newItems

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {

	item := &ListItem{}
	item.Value = v

	if len(l.items) == 0 {
		l.items = append(l.items, item)
		return item
	}

	tail := l.items[len(l.items)-1]
	tail.Next = item
	item.Prev = tail

	l.items = append(l.items, item)

	return item
}

func (l *list) MoveToFront(i *ListItem) {
	for _, value := range l.items {
		if value == i {
			newItems := make([]*ListItem, len(l.items))
			newItems[0] = &ListItem{
				Value: value.Value,
				Prev:  nil,
				Next:  l.items[0],
			}
			j := 1
			for count := 0; count < len(l.items); count++ {
				if i != l.items[count] {
					newItems[j] = l.items[count]
					newItems[j-1].Next = newItems[j]
					newItems[j].Prev = newItems[j-1]
					j++
				}
			}
			newItems[len(newItems)-1].Next = nil
			l.items = newItems
		}
	}
}

func (l *list) Remove(i *ListItem) {
	for key, value := range l.items {
		if value == i {
			last := len(l.items) - 1
			l.items[key] = l.items[last]
			l.items[last] = nil
			l.items = l.items[:last]
		}
	}
}

func NewList() List {
	return new(list)
}
