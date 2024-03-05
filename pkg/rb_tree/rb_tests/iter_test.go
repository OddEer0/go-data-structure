package rbtests

import (
	"testing"

	rbtree "github.com/OddEer0/go-data-structure/pkg/rb_tree"
	"github.com/stretchr/testify/assert"
)

func TestTreeIterationMethods(t *testing.T) {
	t.Run("Should correct InOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{700, 950, 960, 1000, 1030, 1050, 1100, 1150, 1200, 1250}
		tree.InOrderFunc(func(node *rbtree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})

	t.Run("Should correct PreOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1200, 1250}
		tree.PreOrderFunc(func(node *rbtree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})

	t.Run("Should correct PostOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{700, 960, 950, 1030, 1000, 1100, 1250, 1200, 1150, 1050}
		tree.PostOrderFunc(func(node *rbtree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})

	t.Run("Should not iterate empty tree", func(t *testing.T) {
		i := 0
		tree := rbtree.NewRBTree[int, int]()
		tree.PostOrderFunc(func(node *rbtree.Node[int, int]) {
			i++
		})

		tree.PreOrderFunc(func(node *rbtree.Node[int, int]) {
			i++
		})

		tree.InOrderFunc(func(node *rbtree.Node[int, int]) {
			i++
		})

		assert.Equal(t, i, 0)
	})
}
