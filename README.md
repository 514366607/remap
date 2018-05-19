# 关系型Map

## 在写map的时候总会遇到要拿下级数据的情况，代码会写得很冗杂，而且lock来lock去的很烦，就直接写一个，虽然interface的效率不太好。

### 其实操作在map_test.go和index_test.go里已经有详细的代码看里面就行
```
package main

import (
	"log"
	"remap"
)

func main() {

	var m remap.Map
	m.Store("test", "test") // 修改操作
	d, _ := m.Load("test")  // 读取
	log.Println(d)

	m.Delete("test")      //删除
	d, _ = m.Load("test") // 读取
	log.Println(d)

	// 创建索引
	m.Store("test1", "test") // 修改操作
	m.Store("test2", "test") // 修改操作
	m.Store("test3", "test") // 修改操作
	m.Store("test4", "test") // 修改操作
	m.Store("test5", "test") // 修改操作
	m.CreateIndex("索引名", func(v interface{}) bool {
		if v.(string) == "test" {
			return true
		}
		return false
	})
	i, _ := m.Index.GetIndex("索引名") // 取出索引内容
	log.Println(i)

	m.Index.DeleteIndex("索引名") // 删除索引
}
```