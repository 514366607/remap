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

	m.Delete("test")      //删除
	d, _ = m.Load("test") // 读取
	log.Println(2, d)

	// 创建索引
	m.Store("hanmeimei", hanmeimei) // 修改操作
	m.Store("lilei", lilei)         // 修改操作
	m.CreateIndex("boys", func(k, v interface{}) interface{} {
		if v.(student).sex == 1 {
			return v.(student).name
		}
		return nil
	})
	i, _ := m.Index.GetIndex("boys") // 取出索引内容
	log.Println(3, i)

	m.Index.DeleteIndex("boys") // 删除索引

	// 以学员号创建索引
	m.Delete("lilei") // 李雷退学了
	m.CreateIndex("StudentId", func(k, v interface{}) interface{} {
		return v.(student).studentID
	})
	i, _ = m.Index.GetIndex("StudentId") // 取出索引内容
	log.Println(4, i)

	m.Index.DeleteIndex("StudentId") // 删除索引
}
