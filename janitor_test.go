package remap

import (
	"testing"
	"time"
)

func TestExpirationMap(t *testing.T) {
	var m = NewExpiration(time.Millisecond*50, time.Millisecond*100)

	m.Store("test", "test")

	if s, ok := m.Load("test"); ok == true {
		if s != "test" {
			t.Error("获取错误", s)
			return
		}
	} else {
		t.Error("数据不存在")
		return
	}

	time.Sleep(time.Millisecond * 50)
	if _, ok := m.Load("test"); ok == true {
		t.Error("超时数据不应该能拿到")
		return
	}
	if m.Len() != 0 {
		t.Error("数量对不上")
		return
	}

	time.Sleep(time.Millisecond * 50)
	if len(m.itemTime.getExpirationKeys()) != 0 {
		t.Error("没有清除数据")
		return
	}

	m.Range(func(k, v interface{}) bool {
		t.Error("没有清除数据")
		return true
	})

}
