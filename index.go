package remap

// CreateIndex create index
// 创建索引
func (m *Map) CreateIndex(indexName string, f func(v interface{}) bool) map[interface{}]interface{} {

	index := make(map[interface{}]interface{})

	m.Range(func(key, v interface{}) bool {
		newV, ok := v.(*Map)
		if ok == true {
			newIndex := newV.CreateIndex(indexName, f)
			newV.Index.DeleteIndex(indexName)
			for newK, newV := range newIndex {
				index[newK] = newV
			}
			return true
		}

		flag := f(v)
		if flag == true {
			index[key] = v
		}
		return true
	})

	if len(index) > 0 {
		m.Index.mu.Lock()
		defer m.Index.mu.Unlock()
		var indexF = indexF{f: f, data: index}
		m.Index.data[indexName] = indexF
		return m.Index.data[indexName].data
	}
	return nil
}
