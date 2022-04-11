package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	GetItem(key KeyVal) *ListItem
}

type KeyVal interface{}

type ListItem struct {
	Value KeyVal
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	items map[KeyVal]*ListItem
	first *ListItem
	last  *ListItem
}

func (l *list) Len() int {
	return len(l.items)
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{}
	item.Value = v

	if l.Len() > 0 {
		l.first.Prev = item
		item.Next = l.first
	}

	l.first = item

	if l.Len() == 0 {
		l.last = item
	}

	l.items[item.Value] = item

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{}
	item.Value = v

	if l.Len() > 0 {
		l.last.Next = item
		item.Prev = l.last
	}

	l.last = item

	if l.Len() == 0 {
		l.first = item
	}

	l.items[item.Value] = item

	return item
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) Remove(i *ListItem) {
	delete(l.items, i.Value)

	if i == l.first {
		l.first.Next.Prev = nil
		l.first = l.first.Next
		return
	}

	if i == l.last {
		l.last.Prev.Next = nil
		l.last = l.last.Prev
		return
	}

	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
}

func (l *list) GetItem(key KeyVal) *ListItem {
	item, ok := l.items[key]

	if ok {
		return item
	}

	return nil
}

func NewList() List {
	return &list{
		items: make(map[KeyVal]*ListItem),
	}
}
