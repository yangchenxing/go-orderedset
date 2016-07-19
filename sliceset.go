package orderedset

type SliceSet struct {
	len   int
	items []Item
}

func NewSliceSet(cap int) SliceSet {
	return SliceSet{
		items: make([]Item, cap),
	}
}

func (s SliceSet) Get(item Item) Item {
	for i := 0; i < s.len; i++ {
		c := s.items[i]
		if !c.Less(item) {
			if !item.Less(c) {
				return s.items[i]
			}
		}
	}
	return nil
}

func (s SliceSet) ReplaceOrInsert(item Item) {
	if s.len == 0 || s.items[s.len-1].Less(item) {
		s.items = append(s.items, item)
		return
	}
	for i := 0; i < s.len; i++ {
		c := s.items[i]
		if !s.items[i].Less(item) {
			if !item.Less(c) {
				s.items[i] = item
			} else if s.len+1 <= len(s.items) {
				for j := s.len; j > i; j-- {
					s.items[j] = s.items[j-1]
				}
				s.items[i] = item
				s.len++
			}
			break
		}
	}
}

func (s SliceSet) Delete(item Item) {
	for i := 0; i < s.len; i++ {
		c := s.items[i]
		if !c.Less(item) {
			if !item.Less(c) {
				for j := s.len - 1; j > i; j-- {
					s.items[j-1] = s.items[j]
				}
			}
			break
		}
	}
}

type sliceSetIterator struct {
	s SliceSet
	i int
}

func (i sliceSetIterator) Value() Item {
	return i.s.items[i.i]
}

func (i sliceSetIterator) Next() Item {
	if i.i < i.s.len-1 {
		i.i++
	}
	return i.s.items[i.i]
}

func (i sliceSetIterator) Close() {
}

func (s SliceSet) Ascend() Iterator {
	return sliceSetIterator{
		s: s,
	}
}
