package avltests

import (
	"testing"

	avltree "github.com/OddEer0/go-data-structure/tree/avltree"
	"github.com/stretchr/testify/assert"
)

func TestTreeIterationMethods(t *testing.T) {
	t.Run("Should correct InOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{5, 10, 15, 20, 25, 30, 35, 40}
		tree.InOrderFunc(func(node *avltree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})

	t.Run("Should correct PreOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{25, 15, 10, 5, 20, 35, 30, 40}
		tree.PreOrderFunc(func(node *avltree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})

	t.Run("Should correct PostOrderIteration", func(t *testing.T) {
		tree := initTree()
		i := 0
		expectSl := []int{5, 10, 20, 15, 30, 40, 35, 25}
		tree.PostOrderFunc(func(node *avltree.Node[int, int]) {
			assert.Equal(t, node.Value(), expectSl[i])
			i++
		})
		assert.Equal(t, i, len(expectSl))
	})
}
