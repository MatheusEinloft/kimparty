package cmap

type Tuple[K comparable, V any] struct {
	Key K
	Val V
}

type ConcurrentMap[V any] interface {
	Get(key string) (value V, ok bool)
	Set(key string, value V)
	Remove(key string)
	Count() int
	IterWithKey() <-chan Tuple[string, V]
	Iter() <-chan V
	Values() []V
}
