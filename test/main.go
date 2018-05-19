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
