package rbtests

import (
	"testing"

	"github.com/OddEer0/go-data-structure/tree/redblacktree"
	"github.com/stretchr/testify/assert"
)

func TestBaseUtilityMethods(t *testing.T) {
	t.Run("Should correct insert many", func(t *testing.T) {
		args := make([]*redblacktree.Entry[int, int], 0, 100)
		tree := redblacktree.New[int, int]()
		for i := 100; i < 200; i++ {
			args = append(args, &redblacktree.Entry[int, int]{Key: i, Value: i})
		}
		tree.InsertMany(args...)
		assert.Equal(t, tree.Size(), 100)

		temp := make([]int, 0, 100)
		for i := 100; i < 200; i++ {
			temp = append(temp, i)
		}

		assert.Equal(t, temp, tree.Values())
	})

	t.Run("Should correct insert or update", func(t *testing.T) {
		tree := redblacktree.New[int, int]()
		for i := 100; i < 200; i++ {
			tree.InsertOrUpdate(i, i)
		}
		assert.Equal(t, tree.Size(), 100)

		temp := make([]int, 0, 100)
		for i := 100; i < 200; i++ {
			temp = append(temp, i)
		}
		assert.Equal(t, temp, tree.Values())

		for i := 150; i < 200; i++ {
			tree.InsertOrUpdate(i, i*2)
		}
		temp = make([]int, 0, 100)
		for i := 100; i < 150; i++ {
			temp = append(temp, i)
		}
		for i := 150; i < 200; i++ {
			temp = append(temp, i*2)
		}

		assert.Equal(t, temp, tree.Values())
	})

	t.Run("Should correct insert or update many", func(t *testing.T) {
		args := make([]*redblacktree.Entry[int, int], 0, 100)
		tree := redblacktree.New[int, int]()
		for i := 100; i < 200; i++ {
			args = append(args, &redblacktree.Entry[int, int]{Key: i, Value: i})
		}
		tree.InsertOrUpdateMany(args...)
		assert.Equal(t, tree.Size(), 100)
		temp := make([]int, 0, 100)
		for i := 100; i < 200; i++ {
			temp = append(temp, i)
		}
		assert.Equal(t, temp, tree.Values())

		args2 := make([]*redblacktree.Entry[int, int], 0, 60)
		temp = make([]int, 0, 100)
		for i := 100; i < 150; i++ {
			temp = append(temp, i)
		}
		for i := 150; i < 200; i++ {
			temp = append(temp, i*2)
			args2 = append(args2, &redblacktree.Entry[int, int]{Key: i, Value: i * 2})

		}
		tree.InsertOrUpdateMany(args2...)

		assert.Equal(t, temp, tree.Values())
	})

	t.Run("Should correct remove many", func(t *testing.T) {
		args := make([]*redblacktree.Entry[int, int], 0, 100)
		tree := redblacktree.New[int, int]()
		for i := 100; i < 200; i++ {
			args = append(args, &redblacktree.Entry[int, int]{Key: i, Value: i})
		}
		tree.InsertOrUpdateMany(args...)
		assert.Equal(t, tree.Size(), 100)

		temp := make([]int, 0, 50)
		res := make([]int, 0, 50)
		for i := 100; i < 150; i++ {
			temp = append(temp, i)
		}
		for i := 150; i < 200; i++ {
			res = append(res, i)
		}
		tree.RemoveMany(temp...)
		assert.Equal(t, tree.Size(), 50)
		assert.Equal(t, res, tree.Values())
	})
}
