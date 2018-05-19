package remap

import (
	"testing"
)

//go test -test.bench=".*"
func Benchmark_Load(b *testing.B) {
	var m Map
	m.Store("test", "test")
	for i := 0; i < b.N; i++ { //use b.N for looping
		m.Load("test")
	}
}

func Benchmark_Store(b *testing.B) {
	var m Map
	for i := 0; i < b.N; i++ { //use b.N for looping
		m.Store("test", "test")
	}
}

func Benchmark_LoadOrStore(b *testing.B) {
	var m Map
	m.Store("test", "test")
	for i := 0; i < b.N; i++ { //use b.N for looping
		v, ok := m.LoadOrStore("test", "test2")
		if ok == false {
			b.Error(" Store ERROR")
		}

		if v != "test" {
			b.Error("Store Set Value ERROR")
		}

	}
}

func Benchmark_Delete(b *testing.B) {
	var m Map
	for i := 0; i < b.N; i++ { //use b.N for looping
		m.Store("test", "test")
		m.Delete("test")
	}
}

func Benchmark_Range(b *testing.B) {
	var m Map
	m.Store("test1", "test")
	m.Store("test2", "test")
	m.Store("test3", "test")
	m.Store("test4", "test")
	m.Store("test5", "test")
	m.Store("test6", "test")
	m.Store("test7", "test")
	m.Store("test8", "test")
	for i := 0; i < b.N; i++ { //use b.N for looping
		m.Range(func(k, v interface{}) bool {
			return true
		})
	}
}
