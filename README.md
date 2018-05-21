# 关系型Map

## 在写map的时候总会遇到要拿下级数据的情况，代码会写得很冗杂，而且lock来lock去的很烦，就直接写一个，虽然interface的效率不太好。

### 其实操作在map_test.go和index_test.go里已经有详细的代码看里面就行
```
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
	//  print	2018/05/21 15:30:59 1 {1111 韩梅梅 2 10 90}

	m.Delete("test")      //删除
	d, _ = m.Load("test") // 读取
	log.Println(2, d)
	//  print	2018/05/21 15:30:59 2 <nil>

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
	//  print	2018/05/21 15:30:59 3 map[李雷:{1112 李雷 1 10 80}]

	m.Index.DeleteIndex("boys") // 删除索引

	// 以学员号创建索引
	m.Delete("lilei") // 李雷退学了
	m.CreateIndex("StudentId", func(k, v interface{}) interface{} {
		return v.(student).studentID
	})
	i, _ = m.Index.GetIndex("StudentId") // 取出索引内容
	log.Println(4, i)
	// print	2018/05/21 15:30:59 4 map[1111:{1111 韩梅梅 2 10 90}]

	m.Index.DeleteIndex("StudentId") // 删除索引
}

```