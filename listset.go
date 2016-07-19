package orderedset

import (
	"container/list"
)

type ListSet struct {
	*list.List
}

func NewListSet() ListSet {
	return ListSet{
		List: list.New(),
	}
}

func (s ListSet) Get(item Item) Item {
	for i := s.Front(); i != nil; i = i.Next() {
		c := i.Value.(Item)
		if !c.Less(item) {
			if !item.Less(c) {
				return item
			}
			break
		}
	}
	return nil
}

func (s ListSet) ReplaceOrInsert(item Item) {
	if s.Back().Value.(Item).Less(item) {
		s.PushBack(item)
		return
	}
	for i := s.Front(); i != nil; i = i.Next() {
		c := i.Value.(Item)
		if !c.Less(item) {
			if !item.Less(c) {
				i.Value = item
			} else {
				s.InsertBefore(item, i)
			}
			break
		}
	}
}

func (s ListSet) Delete(item Item) {
	for i := s.Front(); i != nil; i = i.Next() {
		c := i.Value.(Item)
		if !c.Less(item) {
			if !item.Less(c) {
				s.Remove(i)
			}
			break
		}
	}
}

type listSetIterator struct {
	*list.Element
}

func (i listSetIterator) Value() Item {
	v := i.Element.Value
	if v != nil {
		return v.(Item)
	}
	return nil
}

func (i listSetIterator) Next() Item {
	i.Element = i.Element.Next()
	return i.Value()
}

func (i listSetIterator) Close() {
}

func (s ListSet) Ascend() Iterator {
	return listSetIterator{Element: s.List.Front()}
}
