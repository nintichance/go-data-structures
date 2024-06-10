package main

type Queue struct {
	Storage map[uint]string
	Head    uint
	Tail    uint
	Size    uint
}

func (q *Queue) Enqueue(element string) {
	q.Tail += 1
	if q.Size == 0 {
		q.Storage = make(map[uint]string)
	}
	q.Size += 1
	q.Storage[q.Tail] = element
}

func (q *Queue) Dequeue() string {
	q.Head += 1
	dequeued := q.Storage[q.Head]
	delete(q.Storage, q.Head)
	q.Size -= 1
	return dequeued
}
