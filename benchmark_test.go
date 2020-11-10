package main

import (
	"fmt"
	"testing"

	rbtree "github.com/emirpasic/gods/trees/redblacktree"
)

func BenchmarkInsert_hashmap(b *testing.B) {
	hashmap := make(map[string]struct{})

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		hashmap[fmt.Sprintf("%d", n)] = struct{}{}
	}
}

func BenchmarkInsert_rbt(b *testing.B) {

	rbt := rbtree.NewWithStringComparator()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		rbt.Put(fmt.Sprintf("%d", n), nil)
	}
}

func BenchmarkGet_hashmap(b *testing.B) {
	hashmap := make(map[string]struct{})

	for n := 0; n < 100; n++ {
		hashmap[fmt.Sprintf("%d", n)] = struct{}{}
	}

	b.ResetTimer()
	i := 0
	for n := 0; n < b.N; n++ {
		if i == 100 {
			i = 0
		}
		_, found := hashmap[fmt.Sprintf("%d", i)]
		if !found {
			b.Fail()
		}

		i++
	}
}

func BenchmarkGet_rbt(b *testing.B) {
	rbt := rbtree.NewWithStringComparator()

	for n := 0; n < 100; n++ {
		rbt.Put(fmt.Sprintf("%d", n), nil)
	}

	b.ResetTimer()
	i := 0
	for n := 0; n < b.N; n++ {
		if i == 100 {
			i = 0
		}
		_, found := rbt.Get(fmt.Sprintf("%d", i))
		if !found {
			b.Fail()
		}

		i++
	}
}

func BenchmarkGetAll_hashmap(b *testing.B) {
	hashmap := make(map[string]struct{})

	for n := 0; n < 100; n++ {
		hashmap[fmt.Sprintf("%d", n)] = struct{}{}
	}

	keys := make([]string, 0, 100)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for key := range hashmap {
			keys = append(keys, key)
		}
	}
}

func BenchmarkGetAll_rbt(b *testing.B) {
	rbt := rbtree.NewWithStringComparator()

	for n := 0; n < 100; n++ {
		rbt.Put(fmt.Sprintf("%d", n), nil)
	}

	strKeys := make([]string, 0, 100)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		keys := rbt.Keys()
		for _, key := range keys {
			strKeys = append(strKeys, key.(string))
		}
	}
}

// func BenchmarkClear_hashmap(b *testing.B) {
// 	hashmap := make(map[string]struct{})

// 	for n := 0; n < 100; n++ {
// 		hashmap[fmt.Sprintf("%d", n)] = struct{}{}
// 	}

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		hashmap = nil
// 	}
// }

// func BenchmarkClear_rbt(b *testing.B) {

// 	rbt := rbtree.NewWithStringComparator()

// 	for n := 0; n < 100; n++ {
// 		rbt.Put(fmt.Sprintf("%d", n), nil)
// 	}

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		rbt.Clear()
// 	}
// }

