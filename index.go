package remap

// CreateIndex create index
// 创建索引
// indexName 索引名
// f( v 值 ) 返回需要做key的值，返回索引内容
func (m *Map) CreateIndex(indexName string, f func(k, v interface{}) bool) *Map {

	index := new(Map)

	m.Range(func(key, val interface{}) bool {
		newV, ok := val.(*Map)
		if ok == true {
			newIndex := newV.CreateIndex(indexName, f)
			if newIndex != nil {
				newV.Index.DeleteIndex(indexName)
				newIndex.Range(func(thisKey, thisVal interface{}) bool {
					index.Store(thisKey, thisVal)
					return true
				})
			}
			return true
		}

		if f(key, val) == true {
			index.Store(key, val)
		}
		return true
	})

	if index.Len() > 0 {
		m.Index.mu.Lock()
		defer m.Index.mu.Unlock()
		indexF := indexF{f: f, data: index}
		m.Index.data[indexName] = &indexF
		return m.Index.data[indexName].data
	}
	return nil
}
