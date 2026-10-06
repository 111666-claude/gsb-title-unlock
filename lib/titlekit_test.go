package titlekit

import "testing"

func rules() []Rule {
	return []Rule{
		{Title: "新兵", NeedLevel: 1, NeedAchv: 0, NeedHours: 0},
		{Title: "老兵", NeedLevel: 10, NeedAchv: 5, NeedHours: 20, Needs: []string{"新兵"}},
	}
}

func TestAllMetUnlocks(t *testing.T) {
	p := &Player{Level: 10, Achievements: 5, Hours: 20, Unlocked: []string{"新兵"}}
	got := Judge(rules(), p, "老兵", &Counter{})
	if got.Status != "unlocked" {
		t.Fatalf("条件全满足应解锁，得到 %+v", got)
	}
}

func TestLevelTooLowStaysLocked(t *testing.T) {
	p := &Player{Level: 3, Achievements: 9, Hours: 99}
	got := Judge(rules(), p, "老兵", &Counter{})
	if got.Status != "locked" {
		t.Fatalf("等级不够应锁住，得到 %+v", got)
	}
	if len(p.Unlocked) != 0 {
		t.Fatalf("锁住时不该写入已解锁列表：%v", p.Unlocked)
	}
}

func TestHoursTooLowStaysLocked(t *testing.T) {
	p := &Player{Level: 30, Achievements: 9, Hours: 1}
	got := Judge(rules(), p, "老兵", &Counter{})
	if got.Status != "locked" {
		t.Fatalf("时长不够应锁住，得到 %+v", got)
	}
}

func TestUnknownTitleStaysLocked(t *testing.T) {
	p := &Player{Level: 99, Achievements: 99, Hours: 99}
	got := Judge(rules(), p, "不存在", &Counter{})
	if got.Status != "locked" {
		t.Fatalf("未知称号应锁住，得到 %+v", got)
	}
}

func TestHasReportsUnlocked(t *testing.T) {
	p := &Player{Unlocked: []string{"新兵"}}
	if !Has(p, "新兵") || Has(p, "老兵") {
		t.Fatal("Has 判断不对")
	}
}
