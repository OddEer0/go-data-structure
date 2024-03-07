package redblacktree

func (k *Entry[T, K]) Key() T {
	return k.key
}

func (k *Entry[T, K]) Value() K {
	return k.value
}
