package remap

import (
	"testing"
)

type student struct {
	name  string // 姓名
	sex   int8   // 性别 1男 2女
	age   int8   // 年龄
	score int8   // 分数
}

var fuckingPrimarySchool Map // 初始化学校
var sixthGrade Map           // 初始化六班级
var fifthGrade Map           // 初始化五班级

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

var xiaoming = student{
	name:  "小明",
	sex:   1,
	age:   10,
	score: 99,
}

func TestMap(t *testing.T) {
	sixthGrade = Map{}
	fifthGrade = Map{}
	fuckingPrimarySchool = Map{}

	sixthGrade.Store("hanmeimei", hanmeimei)

	fuckingPrimarySchool.Store("SixthGrade", &sixthGrade)

	g, ok := fuckingPrimarySchool.Load("SixthGrade")
	if ok == false {
		t.Error("取出六年级数据失败")
	}

	hmm, ok := g.(*Map).Load("hanmeimei")
	if ok == false {
		t.Error("取出韩梅梅失败")
	}

	if hmm.(student).sex != 2 {
		t.Error("韩梅梅咋变性了")
	}

	_, ok = g.(*Map).Load("lilei")
	if ok == true {
		t.Error("取出了不应该存在的李雷")
	}

	hanmeimei.age = 11 // 韩梅梅长大了一岁了
	hmm2, loaded := g.(*Map).LoadOrStore("hanmeimei", hanmeimei)
	if loaded == false {
		t.Error("韩梅梅不见了")
	}

	if hmm2.(student).age == 11 {
		t.Error("韩梅梅应该是永远长不大的才对")
	}

	hanmeimei = hmm.(student)
	fuckingPrimarySchool.Delete("SixthGrade") // 删掉6年级
	_, ok = fuckingPrimarySchool.Load("SixthGrade")
	if ok == true {
		t.Error("六年级不应该还在了")
	}

	sixthGrade.Store("lilei", lilei)
	sixthGrade.Store("xiaoming", xiaoming)
	fuckingPrimarySchool.Store("SixthGrade", &sixthGrade) //重新加回六年级

	var grils = make([]string, 0, 2)
	sixthGrade.Range(func(key, value interface{}) bool {
		if value.(student).sex == 2 {
			grils = append(grils, key.(string))
		}
		return true
	})
	if len(grils) != 1 {
		t.Error("应该只有一个女的")
	}
}
