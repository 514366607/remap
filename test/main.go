package main

import (
	"log"
	"remap"
)

type student struct {
	name  string // 姓名
	sex   int8   // 性别 1男 2女
	age   int8   // 年龄
	score int8   // 分数
}

func main() {

	var hanmeimei = student{
		name:  "韩梅梅",
		sex:   2,
		age:   10,
		score: 90,
	}

	var lilei = student{
		name:  "李雷",
		sex:   1,
		age:   10,
		score: 80,
	}

	var m remap.Map
	m.Store("test", hanmeimei) // 修改操作
	d, _ := m.Load("test")     // 读取
	log.Println(d)

	m.Delete("test")      //删除
	d, _ = m.Load("test") // 读取
	log.Println(d)

	// 创建索引
	m.Store("hanmeimei", hanmeimei) // 修改操作
	m.Store("lilei", lilei)         // 修改操作
	m.CreateIndex("索引名", func(v interface{}) bool {
		if v.(student).sex == 1 {
			return true
		}
		return false
	})
	i, _ := m.Index.GetIndex("索引名") // 取出索引内容
	log.Println(i)

	m.Index.DeleteIndex("索引名") // 删除索引
}
