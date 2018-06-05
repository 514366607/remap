package main

import (
	"log"
	"remap"
)

type student struct {
	studentID int    // 学员号
	name      string // 姓名
	sex       int8   // 性别 1男 2女
	age       int8   // 年龄
	score     int8   // 分数
}

func main() {

	var hanmeimei = student{
		studentID: 1111,
		name:      "韩梅梅",
		sex:       2,
		age:       10,
		score:     90,
	}

	var lilei = student{
		studentID: 1112,
		name:      "李雷",
		sex:       1,
		age:       10,
		score:     80,
	}

	var m remap.Map
	m.Store("test", hanmeimei) // 修改操作
	d, _ := m.Load("test")     // 读取
	log.Println(1, d)
	// 1 {1111 韩梅梅 2 10 90}

	m.Delete("test")      //删除
	d, _ = m.Load("test") // 读取
	log.Println(2, d)
	// <nil>

	// 创建索引
	m.Store("hanmeimei", hanmeimei) // 修改操作
	m.Store("lilei", lilei)         // 修改操作
	m.CreateIndex("boys", func(k, v interface{}) bool {
		if v.(student).sex == 1 {
			return true
		}
		return false
	})
	i, _ := m.Index.GetIndex("boys") // 取出索引内容
	log.Println(3, i)
	// 3 &{{true} <nil> {{0 0} 0 0 0 0} {{{0 0} 0 0 0 0} map[]} map[lilei:{1112 李雷 1 10 80}]}

	m.Index.DeleteIndex("boys") // 删除索引
}
