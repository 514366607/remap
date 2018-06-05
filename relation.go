package remap

import (
	"sync"
)

type indexF struct {
	data *Map
	f    func(k, v interface{}) bool
}

type relation struct {
	mu   sync.RWMutex
	data map[string]*indexF
}

// init 初始化
func (r *relation) New() {
	r.data = make(map[string]*indexF)
}

// GetIndex get index
// 创建索引
func (r *relation) GetIndex(indexName string) (*Map, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	returnData, ok := r.data[indexName]
	if ok == false {
		return nil, false
	}
	return returnData.data, true
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

// delete
// 删除原数据，需要处理删除索引内容
func (r *relation) delete(k, v interface{}) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, indexData := range r.data {
		indexData.data.Delete(k)
	}
}

// StoneKey set index by key of value
// 设置索引的值
func (r *relation) StoneKey(k, v interface{}) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, indexData := range r.data {
		if indexData.f(k, v) == true {
			indexData.data.Store(k, v)
		} else {
			// 删除旧的索引
			indexData.data.Delete(k)
		}
	}
}
