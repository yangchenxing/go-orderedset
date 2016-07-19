package orderedset

type Item interface {
	Less(Item) bool
}

type OrderedSet interface {
	Get(Item) Item
	ReplaceOrInsert(Item)
	Delete(Item)
	Ascend() Iterator
}

type Iterator interface {
	Value() Item
	Next() Item
	Close()
}

func Union(a, b, res OrderedSet) {
	i := a.Ascend()
	j := b.Ascend()
	defer i.Close()
	defer j.Close()
	v := i.Value()
	w := j.Value()
	for v != nil && w != nil {
		if v.Less(w) {
			res.ReplaceOrInsert(v)
			v = i.Next()
		} else {
			res.ReplaceOrInsert(w)
			w = j.Next()
		}
	}
	for ; v != nil; v = i.Next() {
		res.ReplaceOrInsert(v)
	}
	for ; w != nil; w = j.Next() {
		res.ReplaceOrInsert(w)
	}
}

func Intersect(a, b, res OrderedSet) {
	i := a.Ascend()
	j := b.Ascend()
	defer i.Close()
	defer j.Close()
	v := i.Value()
	w := j.Value()
	for v != nil && w != nil {
		if v.Less(w) {
			v = i.Next()
		} else if w.Less(v) {
			w = j.Next()
		} else {
			res.ReplaceOrInsert(v)
			v = i.Next()
			w = j.Next()
		}
	}
}

func Complement(a, b, res OrderedSet) {
	i := a.Ascend()
	j := b.Ascend()
	defer i.Close()
	defer j.Close()
	v := i.Value()
	w := j.Value()
	for v != nil && w != nil {
		if v.Less(w) {
			res.ReplaceOrInsert(v)
			v = i.Next()
		} else if w.Less(v) {
			w = j.Next()
		} else {
			v = i.Next()
			w = j.Next()
		}
	}
}
