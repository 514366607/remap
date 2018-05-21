package remap

// CreateIndex create index
// 创建索引
// indexName 索引名
// f( v 值 ) 返回需要做key的值，返回索引内容
func (m *Map) CreateIndex(indexName string, f func(k, v interface{}) interface{}) map[interface{}]interface{} {

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

		thisKey := f(key, v)
		if thisKey != nil {
			index[thisKey] = v
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
