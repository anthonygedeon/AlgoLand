package main

import (
	"testing"
)

func TestEnqueue(t *testing.T) {

	t.Run("enqueue to list", func(t *testing.T) {
		queue := NewQueue()

		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Enqueue(4)

		want := []int{1, 2, 3, 4}

		if queue.list[2] != want[2] {
			t.Errorf("got %v, want %v", queue.list, want)
		}
	})

	t.Run("enqueue different types", func(t *testing.T) {
		queue := NewQueue()

		queue.Enqueue(true)
		queue.Enqueue(false)
		queue.Enqueue(1)

		want := []interface{}{true, false, 1}

		if queue.list[2] != want[2] {
			t.Errorf("got %v, want %v", queue.list, want)
		}
	})
}

func TestDequeue(t *testing.T) {

	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)

	queue.Dequeue()
	queue.Dequeue()

	if len(queue.list) != 0 {
		t.Errorf("queue is not empty: %v", queue.list)
	}

}

func TestFront(t *testing.T) {
	t.Run("get front of queue", func(t *testing.T) {
		queue := NewQueue()

		queue.Enqueue('A')
		queue.Enqueue('B')

		want := 'A'

		if queue.Front() != want {
			t.Errorf("got %v, want %v", queue.Front(), want)
		}
	})

	t.Run("get front of queue when dequeued", func(t *testing.T) {
		queue := NewQueue()

		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Enqueue(4)

		queue.Dequeue()

	})
}
