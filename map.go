// Copyright 2018 hible
// relationmap 兼容 sync.Map 方法

package remap

import (
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Map map空结构
type Map struct {
	isMake         atomic.Value
	parent         *Map
	mu             sync.RWMutex
	Index          relation
	data           map[interface{}]interface{}
	itemExpiration time.Duration
	janitor        *janitor
	itemTime       *itemTime
}

// New 初始化
func (m *Map) New() {
	m.data = make(map[interface{}]interface{})

	m.itemTime = &itemTime{}
	m.itemTime.new()

	m.Index.New()
}

func (m *Map) tryMake() {
	flag := m.isMake.Load()
	if flag == nil {
		m.mu.Lock()
		m.New()
		m.isMake.Store(true)
		m.mu.Unlock()
	}

}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	m.tryMake()

	if m.itemExpiration > 0 {
		if m.itemTime.check(key) == false {
			return nil, false
		}
	}

	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.data[key], m.data[key] != nil
}

// Store sets the value for a key.
func (m *Map) Store(key, value interface{}) {
	m.tryMake()

	m.mu.Lock()
	defer m.mu.Unlock()

	d, ok := value.(*Map)
	if ok == true {
		d.parent = m
	}

	// 修改索引里的值
	m.Index.StoneKey(key, value)
	if m.parent != nil {
		// 父级的索引也要处理下
		m.parent.Index.StoneKey(key, value)
	}

	m.data[key] = value

	if m.itemExpiration > 0 {
		m.itemTime.update(key, m.itemExpiration)
	}

	return
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
// 如果值存在就直接返回值，如果不存在就设置为传入的值
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	m.tryMake()

	m.mu.RLock()
	actual, loaded = m.data[key]
	m.mu.RUnlock()
	if loaded == true {
		return actual, loaded
	}
	m.Store(key, value)
	return value, loaded
}

// Delete deletes the value for a key.
func (m *Map) Delete(key interface{}) {
	m.tryMake()

	m.mu.Lock()
	defer m.mu.Unlock()

	if m.data[key] == nil {
		return
	}

	// 处理索引数据
	m.Index.delete(key, m.data[key])
	if m.parent != nil {
		m.parent.Index.delete(key, m.data[key])
	}

	delete(m.data, key)

	if m.itemExpiration > 0 {
		m.itemTime.del(key)
	}

	return
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
func (m *Map) Range(f func(key, value interface{}) bool) {
	m.tryMake()

	m.mu.RLock()
	for k, v := range m.data {

		if m.itemExpiration > 0 {
			if m.itemTime.check(k) == false {
				continue
			}
		}

		m.mu.RUnlock()
		if f(k, v) == false {
			m.mu.RLock()
			break
		}
		m.mu.RLock()
	}
	m.mu.RUnlock()
}

// Len MapLen
func (m *Map) Len() int {
	m.tryMake()

	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.itemExpiration > 0 {
		var l = 0
		for k := range m.data {
			if m.itemTime.check(k) == true {
				l++
			}
		}
		return l
	}

	return len(m.data)
}

// NewExpiration Return a new Map with a given default expiration duration and cleanup interval
func NewExpiration(expiration, cleanupInterval time.Duration) *Map {
	var m = &Map{itemExpiration: expiration}
	m.isMake.Store(true)
	m.New()

	newJanitor(m, cleanupInterval)
	runtime.SetFinalizer(m, stopJanitor)

	return m
}
