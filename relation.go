package remap

import (
	"sync"
)

type relation struct {
	mu   sync.RWMutex
	data map[string]map[interface{}]interface{}
}

// init 初始化
func (r *relation) New() {
	r.data = make(map[string]map[interface{}]interface{})
}

// GetIndex get index
// 创建索引
func (r *relation) GetIndex(indexName string) (map[interface{}]interface{}, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.data[indexName], r.data[indexName] != nil
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (r *relation) Range(f func(key, value interface{}) bool) {
	r.mu.RLock()
	copyMap := r.data
	r.mu.RUnlock()
	for k, v := range copyMap {
		if f(k, v) == false {
			break
		}
	}
}

// DeleteIndex delete index
// 删除索引
func (r *relation) DeleteIndex(indexName string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.data, indexName)
}

// DeleteKey delete index key
// 删除索引内容
func (r *relation) DeleteKey(key interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, indexData := range r.data {
		for newk := range indexData {
			if newk == key {
				delete(indexData, newk)
				break
			}
		}
	}
}

// StoneKey set index by key of value
// 设置索引的值
func (r *relation) StoneKey(key, value interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, indexData := range r.data {
		for dataKey := range indexData {
			if dataKey == key {
				indexData[dataKey] = value
				break
			}
		}
	}
}
