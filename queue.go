package main

type Queue struct {
	list []interface{}
	head int
	tail int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(t interface{}) {
	q.tail++
	q.list = append(q.list, t)
}

func (q *Queue) Dequeue() {
	q.head++
	q.list = q.list[1:]
}

func (q *Queue) Front() interface{} {
	return q.list[q.head]
}
