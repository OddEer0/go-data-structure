package list_stack_test

import (
	"fmt"
	liststack "github.com/OddEer0/go-data-structure/stack/list_stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListStack(t *testing.T) {

	t.Run("Should correct Push", func(t *testing.T) {
		stack := liststack.New[int]()
		assert.Equal(t, 0, stack.Size())

		stack.Push(1)
		assert.Equal(t, 1, stack.Size())

		stack.Push(2)
		assert.Equal(t, 2, stack.Size())

		stack.Push(3)
		assert.Equal(t, 3, stack.Size())
	})

	t.Run("Should correct IsEmpty", func(t *testing.T) {
		stack := liststack.New[int]()
		assert.Equal(t, 0, stack.Size())
		assert.True(t, stack.IsEmpty())

		stack.Push(1)
		assert.Equal(t, 1, stack.Size())
		assert.False(t, stack.IsEmpty())
	})

	t.Run("Should correct String", func(t *testing.T) {
		stack := liststack.New[int]()
		assert.Equal(t, "", stack.String())
		stack.Push(1)
		stack.Push(2)
		stack.Push(2)
		stack.Push(3)
		assert.Equal(t, "1 2 2 3", stack.String())
	})

	t.Run("Should correct Clear", func(t *testing.T) {
		stack := liststack.New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(2)
		assert.Equal(t, 3, stack.Size())
		stack.Clear()
		assert.Equal(t, 0, stack.Size())
	})

	t.Run("Should correct Peek", func(t *testing.T) {
		stack := liststack.New[int]()
		stack.Push(5)
		assert.Equal(t, 5, stack.Peek())
		assert.Equal(t, 1, stack.Size())

		stack.Push(9)
		assert.Equal(t, 9, stack.Peek())
		assert.Equal(t, 2, stack.Size())

		stack.Push(13)
		assert.Equal(t, 13, stack.Peek())
		assert.Equal(t, 3, stack.Size())
	})

	t.Run("Should correct Pop", func(t *testing.T) {
		stack := liststack.New[int]()
		stack.Push(5)
		stack.Push(9)
		stack.Push(13)

		assert.Equal(t, 3, stack.Size())
		assert.Equal(t, 13, stack.Pop())
		assert.Equal(t, 2, stack.Size())

		assert.Equal(t, 9, stack.Pop())
		assert.Equal(t, 1, stack.Size())

		assert.Equal(t, 5, stack.Pop())
		assert.Equal(t, 0, stack.Size())

		var val int
		assert.Equal(t, val, stack.Pop())
		assert.Equal(t, 0, stack.Size())
	})

	t.Run("Should correct Copy", func(t *testing.T) {
		stack := liststack.New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		copyStack := stack.Copy()

		assert.Equal(t, stack.Size(), copyStack.Size())
		assert.False(t, fmt.Sprintf("%p", stack) == fmt.Sprintf("%p", copyStack))

		copyStack.Pop()
		assert.Equal(t, 3, stack.Size())
		assert.Equal(t, 2, copyStack.Size())
	})
}
