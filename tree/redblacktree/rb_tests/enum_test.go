package rbtests

import (
	"fmt"
	"testing"

	"github.com/OddEer0/go-data-structure/tree/redblacktree"
	"github.com/stretchr/testify/assert"
)

func TestEnumFunction(t *testing.T) {
	t.Run("Should correct Each enum", func(t *testing.T) {
		tree := initTree2()
		slK := tree.Keys()
		slV := tree.Values()

		i := 0
		tree.Each(func(key int, value int) {
			assert.Equal(t, slK[i], key)
			assert.Equal(t, slV[i], value)
			i++
		})
	})

	t.Run("Should correct EachLast", func(t *testing.T) {
		tree := initTree2()
		slK := tree.Keys()
		slV := tree.Values()

		i := 0
		tree.EachLast(func(key int, value int) {
			assert.Equal(t, slK[tree.Size()-i-1], key)
			assert.Equal(t, slV[tree.Size()-i-1], value)
			i++
		})
	})

	t.Run("Should correct Some", func(t *testing.T) {
		tree := initTree2()
		checkedValue := 14400
		has := tree.Some(func(key int, value int) bool {
			if value == checkedValue {
				return true
			}
			return false
		})
		assert.False(t, has)
		tree.Insert(checkedValue, checkedValue)
		has = tree.Some(func(key int, value int) bool {
			if value == checkedValue {
				return true
			}
			return false
		})
		assert.True(t, has)
	})

	t.Run("Should correct Every", func(t *testing.T) {
		tree := initTree2()
		maxValue := 10000
		check := tree.Every(func(key int, value int) bool {
			if value <= maxValue {
				return true
			}
			return false
		})
		assert.True(t, check)
		tree.Insert(maxValue+1, maxValue+1)
		check = tree.Every(func(key int, value int) bool {
			if value <= maxValue {
				return true
			}
			return false
		})
		assert.False(t, check)
	})

	t.Run("Should correct Map", func(t *testing.T) {
		tree := initTree2()

		newTree := tree.Map(func(key int, value int) (int, int) {
			return key * 2, value * 2
		})
		assert.False(t, fmt.Sprintf("%p", tree) == fmt.Sprintf("%p", newTree))
		slK := newTree.Keys()
		slV := newTree.Values()

		i := 0
		iterator := tree.Iterator()
		for iterator.Next() {
			assert.Equal(t, slK[i], iterator.Key()*2)
			assert.Equal(t, slV[i], iterator.Value()*2)
			i++
		}
	})

	t.Run("Should correct Filter", func(t *testing.T) {
		tree := redblacktree.New[int, int]()
		for i := 1; i <= 1000; i++ {
			tree.Insert(i, i)
		}
		newEvenTree := tree.Filter(func(key int, value int) bool {
			if key%2 == 0 {
				return true
			}
			return false
		})
		assert.False(t, fmt.Sprintf("%p", tree) == fmt.Sprintf("%p", newEvenTree))
		assert.Equal(t, 500, newEvenTree.Size())

		iterator := newEvenTree.Iterator()
		for iterator.Next() {
			assert.True(t, iterator.Key()%2 == 0)
		}
	})

	t.Run("Should correct Concat", func(t *testing.T) {
		tree := redblacktree.New[int, int]()
		tree2 := redblacktree.New[int, int]()
		tree3 := redblacktree.New[int, int]()

		for i := 1; i <= 1000; i++ {
			tree.Insert(i, i)
			tree2.Insert(1000+i, 1000+i)
			tree3.Insert(2000+i, 3000+i)
		}

		newTree := tree.Concat(tree2, tree3)
		assert.Equal(t, 1000, tree.Size())
		assert.Equal(t, 1000, tree2.Size())
		assert.Equal(t, 1000, tree3.Size())
		assert.Equal(t, 3000, newTree.Size())

		sl := tree.Values()
		sl = append(sl, tree2.Values()...)
		sl = append(sl, tree3.Values()...)

		iterator := newTree.Iterator()
		i := 0
		for iterator.Next() {
			assert.Equal(t, sl[i], iterator.Value())
			i++
		}
	})
}
