// Package decorator 测试
package decorator

import "testing"

func TestFAW_Checklist(t *testing.T) {
	l := new(Liberation)
	t.Logf("原厂清单：%s", l.Checklist())
	faw := NewFAW(l, 1)
	faw = NewFAW(faw, 2)
	faw = NewFAW(faw, 3)
	for _, value := range faw.Checklist() {
		t.Logf("改装情况：%s", value)
	}
}
