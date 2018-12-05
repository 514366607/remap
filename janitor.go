package remap

import (
	"sync"
	"time"
)

type janitor struct {
	Interval time.Duration
	stop     chan bool
}

// newJanitor 初始化
func newJanitor(m *Map, cleanupInterval time.Duration) {
	m.janitor = &janitor{
		Interval: cleanupInterval,
		stop:     make(chan bool),
	}
	go m.janitor.run(m)
}

func (j *janitor) run(m *Map) {
	ticker := time.NewTicker(j.Interval)
	for {
		select {
		case <-ticker.C:
			// 删除过期数据
			clearExpiration(m)
		case <-j.stop:
			ticker.Stop()
			return
		}
	}
}

func stopJanitor(m *Map) {
	m.janitor.stop <- true
}

func clearExpiration(m *Map) {
	for _, k := range m.itemTime.getExpirationKeys() {
		m.itemTime.del(k)
		m.Delete(k)
	}
}

type itemTime struct {
	data map[interface{}]int64
	lock sync.RWMutex
}

// 初始化
func (i *itemTime) new() {
	i.data = make(map[interface{}]int64)
}

// check 判断是否过期
func (i *itemTime) check(key interface{}) bool {
	i.lock.RLock()
	defer i.lock.RUnlock()
	var d, ok = i.data[key]
	if ok == true {
		if d > time.Now().UnixNano() {
			return true
		}
	}
	return false
}

// update 更新时间
func (i *itemTime) update(key interface{}, t time.Duration) {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.data[key] = time.Now().Add(t).UnixNano()
}

// del 删除
func (i *itemTime) del(key interface{}) {
	i.lock.Lock()
	delete(i.data, key)
	i.lock.Unlock()
	return
}

// getExpirationKeys 取得过期键值
func (i *itemTime) getExpirationKeys() []interface{} {
	i.lock.RLock()
	defer i.lock.RUnlock()
	var exKeys = make([]interface{}, 0)
	for key, t := range i.data {
		if t <= time.Now().UnixNano() {
			exKeys = append(exKeys, key)
		}
	}
	return exKeys
}
