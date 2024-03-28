package circular_queue_test

import (
	circularqueue "github.com/OddEer0/go-data-structure/queue/circular_queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	t.Run("Should correct Unshift", func(t *testing.T) {
		queue := circularqueue.New[int](4)
		assert.Equal(t, 0, queue.Size())
		queue.Unshift(1)
		assert.Equal(t, 1, queue.Size())
		queue.Unshift(2)
		assert.Equal(t, 2, queue.Size())
		queue.Unshift(3)
		assert.Equal(t, 3, queue.Size())
		queue.Unshift(4)
		assert.Equal(t, 4, queue.Size())
		queue.Unshift(5)
		assert.Equal(t, 4, queue.Size())
	})

	t.Run("Should correct Peek", func(t *testing.T) {
		queue := circularqueue.New[int](4)
		var res int
		assert.Equal(t, res, queue.Peek())
		queue.Unshift(1)
		assert.Equal(t, 1, queue.Peek())
		queue.Unshift(2)
		assert.Equal(t, 1, queue.Peek())
		queue.Unshift(3)
		assert.Equal(t, 1, queue.Peek())
		queue.Unshift(4)
		assert.Equal(t, 1, queue.Peek())
	})

	t.Run("Should correct Shift", func(t *testing.T) {
		queue := circularqueue.New[int](4)
		var res int
		assert.Equal(t, res, queue.Shift())
		queue.Unshift(1)
		queue.Unshift(2)
		assert.Equal(t, 1, queue.Shift())
		assert.Equal(t, 2, queue.Peek())
		assert.Equal(t, 1, queue.Size())
		queue.Unshift(3)
		queue.Unshift(4)
		queue.Unshift(5)
		assert.Equal(t, 2, queue.Shift())
		assert.Equal(t, 3, queue.Peek())
		assert.Equal(t, 3, queue.Size())
		queue.Unshift(6)
		assert.Equal(t, 3, queue.Shift())
		assert.Equal(t, 4, queue.Peek())
		assert.Equal(t, 3, queue.Size())

		assert.Equal(t, 4, queue.Shift())
		assert.Equal(t, 5, queue.Peek())
		assert.Equal(t, 2, queue.Size())

		assert.Equal(t, 5, queue.Shift())
		assert.Equal(t, 6, queue.Peek())
		assert.Equal(t, 1, queue.Size())

		assert.Equal(t, 6, queue.Shift())
		assert.Equal(t, res, queue.Peek())
		assert.Equal(t, 0, queue.Size())
	})

	t.Run("Should correct IsEmpty", func(t *testing.T) {
		queue := circularqueue.New[int](4)
		assert.True(t, queue.IsEmpty())
		queue.Unshift(1)
		queue.Unshift(2)
		queue.Unshift(3)
		queue.Unshift(4)
		assert.False(t, queue.IsEmpty())
		queue.Shift()
		queue.Shift()
		queue.Unshift(5)
		queue.Unshift(6)
		assert.False(t, queue.IsEmpty())
		queue.Shift()
		queue.Shift()
		queue.Shift()
		queue.Shift()
		assert.True(t, queue.IsEmpty())
	})

	t.Run("Should correct String", func(t *testing.T) {
		queue := circularqueue.New[int](4)
		assert.Equal(t, "", queue.String())
		queue.Unshift(1)
		queue.Unshift(2)
		queue.Unshift(3)
		queue.Unshift(4)
		assert.Equal(t, "1 2 3 4", queue.String())
		queue.Shift()
		queue.Shift()
		queue.Unshift(5)
		queue.Unshift(6)
		assert.Equal(t, "3 4 5 6", queue.String())
	})

	t.Run("Should correct Clear", func(t *testing.T) {
		queue := circularqueue.New[int](4)
		queue.Unshift(1)
		queue.Unshift(2)
		queue.Unshift(3)
		queue.Unshift(4)
		queue.Clear()
		assert.Equal(t, 0, queue.Size())
		assert.True(t, queue.IsEmpty())
	})
}
