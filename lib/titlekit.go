// Package titlekit 判定称号解锁。
package titlekit

// Rule 是一条称号规则。
type Rule struct {
	Title     string   `json:"title"`
	NeedLevel int      `json:"need_level"`
	NeedAchv  int      `json:"need_achv"`
	NeedHours int      `json:"need_hours"`
	Needs     []string `json:"needs"`
	Event     bool     `json:"event"`
}

// Player 是判定用到的玩家快照。
type Player struct {
	Level        int      `json:"level"`
	Achievements int      `json:"achievements"`
	Hours        int      `json:"hours"`
	Unlocked     []string `json:"unlocked"`
}

// Counter 统计本轮检查过的称号条数。
type Counter struct {
	Scanned int
}

// Decision 是一次判定的输出。
type Decision struct {
	Title    string `json:"title"`
	Status   string `json:"status"`
	Unlocked int    `json:"unlocked"`
	Scanned  int    `json:"scanned"`
}

func has(list []string, want string) bool {
	for _, item := range list {
		if item == want {
			return true
		}
	}
	return false
}

// Judge 判定一个称号能不能解锁，并在成功时记进玩家的已解锁列表。
func Judge(rules []Rule, p *Player, title string, c *Counter) Decision {
	if has(p.Unlocked, title) {
		return Decision{Title: title, Status: "unlocked", Unlocked: len(p.Unlocked), Scanned: c.Scanned}
	}
	var rule *Rule
	for i := range rules {
		c.Scanned++
		if rules[i].Title == title {
			rule = &rules[i]
			break
		}
	}
	if rule == nil {
		return Decision{Title: title, Status: "locked", Unlocked: len(p.Unlocked), Scanned: c.Scanned}
	}
	if !rule.Event && p.Level < rule.NeedLevel {
		return Decision{Title: title, Status: "locked", Unlocked: len(p.Unlocked), Scanned: c.Scanned}
	}
	if p.Achievements < rule.NeedAchv || p.Hours < rule.NeedHours {
		return Decision{Title: title, Status: "locked", Unlocked: len(p.Unlocked), Scanned: c.Scanned}
	}
	for _, need := range rule.Needs {
		if !has(p.Unlocked, need) {
			return Decision{Title: title, Status: "locked", Unlocked: len(p.Unlocked), Scanned: c.Scanned}
		}
	}
	p.Unlocked = append(p.Unlocked, title)
	return Decision{Title: title, Status: "unlocked", Unlocked: len(p.Unlocked), Scanned: c.Scanned}
}

// Has 报告玩家是否已解锁某个称号。
func Has(p *Player, title string) bool { return has(p.Unlocked, title) }
