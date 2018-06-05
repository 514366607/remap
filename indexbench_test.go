package remap

import (
	"fmt"
	"testing"
)

//go test -test.bench=".*"
func Benchmark_CreteIndex(b *testing.B) {
	var m Map
	m.Store("test", "test")
	m.Store("test1", "test")
	m.Store("test2", "test")
	m.Store("test3", "test")
	m.Store("test4", "test")
	for i := 0; i < b.N; i++ { //use b.N for looping
		m.CreateIndex(fmt.Sprintf("test%d", i), func(k, v interface{}) bool {
			if v == "test" {
				return true
			}
			return false
		})
	}
}

func Benchmark_MultiIndexGet(b *testing.B) {
	var m Map
	var m2 Map
	m2.Store("test1", "test")
	m2.Store("test2", "test")
	m2.Store("test3", "test")
	m2.Store("test4", "test")
	m.Store(1, &m2)
	m.CreateIndex("test", func(k, v interface{}) bool {
		if v == "test" {
			return true
		}
		return false
	})
	for i := 0; i < b.N; i++ { //use b.N for looping
		m.Index.GetIndex("test")
	}
}
func Benchmark_DeleteIndex(b *testing.B) {
	var m Map
	var m2 Map
	m2.Store("test1", "test")
	m2.Store("test2", "test")
	m2.Store("test3", "test")
	m2.Store("test4", "test")
	m.Store(1, &m2)
	m.CreateIndex("test", func(k, v interface{}) bool {
		if v == "test" {
			return true
		}
		return false
	})
	for i := 0; i < b.N; i++ { //use b.N for looping
		m2.Delete("test4")
	}
}

func Benchmark_MultiIndexDelete(b *testing.B) {
	var m Map
	var m2 Map
	m2.Store("test1", "test")
	m2.Store("test2", "test")
	m2.Store("test3", "test")
	m2.Store("test4", "test")
	m.Store(1, &m2)
	m.CreateIndex("test", func(k, v interface{}) bool {
		if v == "test" {
			return true
		}
		return false
	})
	m2.Delete("test4")
	for i := 0; i < b.N; i++ { //use b.N for looping
		i, _ := m.Index.GetIndex("test")
		if i.Len() != 3 {
			b.Errorf("数量不对%d", i.Len())
		}
	}
}

func Benchmark_IndexRange(b *testing.B) {
	var m Map
	var m2 Map
	m2.Store("test1", "test")
	m2.Store("test2", "test")
	m2.Store("test3", "test")
	m2.Store("test4", "test")
	m.Store(1, &m2)
	m.CreateIndex("test", func(k, v interface{}) bool {
		if v == "test" {
			return true
		}
		return false
	})
	m2.Delete("test4")
	for i := 0; i < b.N; i++ { //use b.N for looping
		m.Index.Range(func(k, v interface{}) bool {
			return true
		})
	}
}
