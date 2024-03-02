package avltests

import (
	"testing"

	avltree "github.com/OddEer0/go-data-structure/pkg/avl_tree"
	"github.com/stretchr/testify/assert"
)

func TestTreeIterationMethods(t *testing.T) {
	t.Run("Should correct InOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{5, 6, 7, 8, 10, 13, 14, 15}
		tree.InOrderFunc(func(node *avltree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})

	t.Run("Should correct PreOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{10, 5, 7, 6, 8, 14, 13, 15}
		tree.PreOrderFunc(func(node *avltree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})

	t.Run("Should correct PostOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{6, 8, 7, 5, 13, 15, 14, 10}
		tree.PostOrderFunc(func(node *avltree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})
}
