package rbtests

import (
	"sort"
	"testing"

	"github.com/OddEer0/go-data-structure/pkg/redblacktree"
	"github.com/stretchr/testify/assert"
)

func TestKeysValuesEntries(t *testing.T) {
	t.Run("Should correct keys values entries and other variations", func(t *testing.T) {
		tree := initTree()
		tree.InOrderFunc(func(node *redblacktree.Node[int, int]) {
			tree.Update(node.Key(), node.Value()*2)
		})
		res := []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1200, 1250}
		res2 := []int{700, 960, 950, 1030, 1000, 1100, 1250, 1200, 1150, 1050}
		res3 := []int{1050, 1000, 950, 700, 960, 1030, 1150, 1100, 1200, 1250}
		sort.Ints(res3)
		assert.Equal(t, res, tree.PreOrderKeys())
		assert.Equal(t, res2, tree.PostOrderKeys())
		assert.Equal(t, res3, tree.Keys())

		for i := range res {
			res[i] *= 2
			res2[i] *= 2
			res3[i] *= 2
		}

		assert.Equal(t, res, tree.PreOrderValues())
		assert.Equal(t, res2, tree.PostOrderValues())
		assert.Equal(t, res3, tree.Values())

		eres := tree.PreOrderEntries()
		eres2 := tree.PostOrderEntries()
		eres3 := tree.Entries()

		for i := range eres {
			assert.Equal(t, eres[i].Key, res[i]/2)
			assert.Equal(t, eres[i].Value, res[i])
			assert.Equal(t, eres2[i].Key, res2[i]/2)
			assert.Equal(t, eres2[i].Value, res2[i])
			assert.Equal(t, eres3[i].Key, res3[i]/2)
			assert.Equal(t, eres3[i].Value, res3[i])
		}
	})
}
