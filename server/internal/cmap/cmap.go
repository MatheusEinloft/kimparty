package cmap

import (
	conMap "github.com/orcaman/concurrent-map/v2"
)

type ConMap[V any] struct {
	in conMap.ConcurrentMap[string, V]
}

func (cm *ConMap[V]) Get(key string) (value V, ok bool) {
	return cm.in.Get(key)
}

func (cm *ConMap[V]) Set(key string, value V) {
	cm.in.Set(key, value)
}

func (cm *ConMap[V]) Remove(key string) {
	cm.in.Remove(key)

	if cm.in.Count()%200 == 0 {
		newMap := conMap.New[V]()

		for item := range cm.in.IterBuffered() {
			newMap.Set(item.Key, item.Val)
		}
	}
}

func (cm *ConMap[V]) Count() int {
	return cm.in.Count()
}

func (cm *ConMap[V]) IterWithKey() <-chan Tuple[string, V] {
	ch := make(chan Tuple[string, V])

	go func() {
		for item := range cm.in.IterBuffered() {
			ch <- Tuple[string, V]{item.Key, item.Val}
		}
		close(ch)
	}()

	return ch
}

func (cm *ConMap[V]) Iter() <-chan V {
	ch := make(chan V)

	go func() {
		for item := range cm.in.IterBuffered() {
			ch <- item.Val
		}
		close(ch)
	}()

	return ch
}

func (cm *ConMap[V]) Values() []V {
	var all []V

	for item := range cm.in.IterBuffered() {
		all = append(all, item.Val)
	}

	return all
}

func New[V any]() ConcurrentMap[V] {
	return &ConMap[V]{in: conMap.New[V]()}
}
