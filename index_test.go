package remap

import (
	"testing"
)

func TestIndex(t *testing.T) {
	sixthGrade = Map{}
	fifthGrade = Map{}
	fuckingPrimarySchool = Map{}

	sixthGrade.Store("hanmeimei", &hanmeimei)
	sixthGrade.Store("lilei", &lilei)

	// 班级男生有那些
	sixthGrade.CreateIndex("GradeBoys", func(k, v interface{}) bool {
		if v.(*student).sex == 1 {
			return true
		}
		return false
	})

	gradeBoys, ok := sixthGrade.Index.GetIndex("GradeBoys")
	if ok == false {
		t.Error("获取全班男生索引失败")
	}
	if gradeBoys.Len() != 1 {
		t.Error("索引拿出来的班级男生数量不对")
	}

	// 李雷克隆人杀了李雷顶替了李雷
	var lileiCopy = student{
		name:  "李雷",
		sex:   1,
		age:   1,
		score: 100,
	}
	sixthGrade.Store("lilei", &lileiCopy)
	ll, _ := gradeBoys.Load("lilei")
	if ll.(*student).age != 1 || ll.(*student).score != 100 {
		t.Error("李雷克隆人没有顶替成功")
	}

	// 李雷克隆人 变性为  李雷克隆人（女）
	lileiCopy.sex = 2
	sixthGrade.Store("lilei", &lileiCopy)
	if gradeBoys.Len() > 0 {
		t.Error("索引修复失败", gradeBoys)
	}
}

// TestMainIndexByPt 传入指针测试
func TestMainIndexByPt(t *testing.T) {
	sixthGrade = Map{}
	fifthGrade = Map{}
	fuckingPrimarySchool = Map{}

	sixthGrade.Store("hanmeimei", &hanmeimei)
	sixthGrade.Store("lilei", &lilei)

	fifthGrade.Store("xiaoming", &xiaoming)

	// 学校男生有那些
	fuckingPrimarySchool.Store("FifthGrade", &fifthGrade)
	fuckingPrimarySchool.Store("SixthGrade", &sixthGrade)
	fuckingPrimarySchool.CreateIndex("SchoolBoys", func(k, v interface{}) bool {
		if v.(*student).sex == 1 {
			return true
		}
		return false
	})
	schoolBoys, ok := fuckingPrimarySchool.Index.GetIndex("SchoolBoys")
	if ok == false {
		t.Error("获取全校男生索引失败")
	}
	if schoolBoys.Len() != 2 {
		t.Error("索引拿出来的学校男生数量不对")
	}

	fifthGrade.Delete("xiaoming") // 小明退学了
	lilei.age = 11                // 李雷刚过了生日
	schoolBoys, ok = fuckingPrimarySchool.Index.GetIndex("SchoolBoys")
	if ok == false {
		t.Error("第二次获取全校男生索引失败")
	}
	if schoolBoys.Len() != 1 {
		t.Error("小明已经给退学了，索引没有更新到")
	}
	ll, _ := schoolBoys.Load("lilei")
	if ll.(*student).age != 11 {
		t.Error("李雷过了生日应该大一岁了")
	}

	// 再次删除下看下会不会报错
	sixthGrade.Delete("xiaoming") // 小明退学了

}

func TestMainIndex(t *testing.T) {
	sixthGrade = Map{}
	fifthGrade = Map{}
	fuckingPrimarySchool = Map{}

	sixthGrade.Store("hanmeimei", hanmeimei)
	sixthGrade.Store("lilei", lilei)

	fifthGrade.Store("xiaoming", xiaoming)

	// 学校男生有那些
	fuckingPrimarySchool.Store("FifthGrade", &fifthGrade)
	fuckingPrimarySchool.Store("SixthGrade", &sixthGrade)
	fuckingPrimarySchool.CreateIndex("SchoolBoys", func(k, v interface{}) bool {
		if v.(student).sex == 1 {
			return true
		}
		return false
	})
	schoolBoys, ok := fuckingPrimarySchool.Index.GetIndex("SchoolBoys")
	if ok == false {
		t.Error("获取全校男生索引失败")
	}
	if schoolBoys.Len() != 2 {
		t.Error("索引拿出来的学校男生数量不对")
	}

	fifthGrade.Delete("xiaoming")    // 小明退学了
	lilei.age = 11                   // 李雷刚过了生日
	sixthGrade.Store("lilei", lilei) // 不是指针要重新store下才行

	schoolBoys, ok = fuckingPrimarySchool.Index.GetIndex("SchoolBoys")
	if ok == false {
		t.Error("第二次获取全校男生索引失败")
	}
	if schoolBoys.Len() != 1 {
		t.Error("小明已经给退学了，索引没有更新到", schoolBoys)
	}
	ll, _ := schoolBoys.Load("lilei")
	if ll.(student).age != 11 {
		t.Error("李雷过了生日应该大一岁了")
	}

	// 再次删除下看下会不会报错
	sixthGrade.Delete("xiaoming") // 小明退学了

}

func TestStoreIndex(t *testing.T) {
	sixthGrade = Map{}
	fifthGrade = Map{}

	sixthGrade.Store("hanmeimei", hanmeimei)
	sixthGrade.Store("lilei", lilei)

	// 学校男生有那些
	fuckingPrimarySchool = Map{}
	fuckingPrimarySchool.Store("FifthGrade", &fifthGrade)
	fuckingPrimarySchool.Store("SixthGrade", &sixthGrade)
	fuckingPrimarySchool.CreateIndex("SchoolBoys", func(k, v interface{}) bool {
		if v.(student).sex == 1 {
			return true
		}
		return false
	})

	fifthGrade.Store("xiaoming", xiaoming) //创建索引后加入小时，看索引是否加入
	schoolBoys, ok := fuckingPrimarySchool.Index.GetIndex("SchoolBoys")
	if ok == false {
		t.Error("获取全校男生索引失败")
	}
	if schoolBoys.Len() != 2 {
		t.Error("索引拿出来的学校男生数量不对")
	}
}

func TestSecIndex(t *testing.T) {
	sixthGrade = Map{}
	fifthGrade = Map{}

	sixthGrade.Store("hanmeimei", &hanmeimei)
	sixthGrade.Store("lilei", &lilei)

	xiaoming.age = 11
	fifthGrade.Store("xiaoming", &xiaoming)

	// 学校男生有那些
	fuckingPrimarySchool = Map{}
	fuckingPrimarySchool.Store("FifthGrade", &fifthGrade)
	fuckingPrimarySchool.Store("SixthGrade", &sixthGrade)
	schoolBoy := fuckingPrimarySchool.CreateIndex("SchoolBoys", func(k, v interface{}) bool {
		if v.(*student).sex == 1 {
			return true
		}
		return false
	})

	age10 := schoolBoy.CreateIndex("Age10", func(k, v interface{}) bool {
		if v.(*student).age == 10 {
			return true
		}
		return false
	})

	if age10.Len() != 1 {
		t.Error("年龄索引创建失败")
	}

	lilei.age = 11
	sixthGrade.Store("lilei", &lilei)
	if age10.Len() != 0 {
		t.Error("修改基础数据，年龄索引没有修改到")
	}

}
