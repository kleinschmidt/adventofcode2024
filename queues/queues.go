package queues

import "iter"

type Elem[E comparable] struct {
	e    E
	next *Elem[E]
}

type Queue[E comparable] struct {
	first *Elem[E]
	last  *Elem[E]
}

func (queue *Queue[E]) Enqueue(e E) {
	elem := &Elem[E]{e, nil}
	if queue.first == nil {
		queue.first = elem
		queue.last = elem
	} else {
		queue.last.next = elem
		queue.last = queue.last.next
	}
}

func (queue *Queue[E]) Dequeue() *E {
	first := queue.first
	if first == nil {
		return nil
	}
	queue.first = first.next
	if first.next == nil {
		queue.last = nil
	}
	return &first.e
}

func (queue *Queue[E]) HasNext() bool {
	return queue.first != nil
}

// this is a bit iffy: it's destructive!
func (queue *Queue[E]) All() iter.Seq[E] {
	return func(yield func(E) bool) {
		for {
			next := queue.Dequeue()
			if next == nil || !yield(*next) {
				return
			}
		}
	}
}

func NewQueue[E comparable](arr []E) *Queue[E] {
	queue := new(Queue[E])
	for _, e := range arr {
		queue.Enqueue(e)
	}
	return queue
}
