// Command titlecli 是称号解锁判定的场景入口。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"example.com/titlekit"
)

func emit(v any) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func rules() []titlekit.Rule {
	return []titlekit.Rule{
		{Title: "新兵", NeedLevel: 1},
		{Title: "探险家", NeedLevel: 5, NeedAchv: 3, NeedHours: 10},
		{Title: "老兵", NeedLevel: 10, NeedAchv: 5, NeedHours: 20, Needs: []string{"新兵"}},
		{Title: "老兵·活动", NeedLevel: 30, NeedAchv: 5, NeedHours: 20, Event: true},
	}
}

func main() {
	sample := flag.String("sample", "", "achv|prereq|event|replay|work")
	flag.Parse()

	switch *sample {
	case "achv":
		p := &titlekit.Player{Level: 9, Achievements: 0, Hours: 30}
		emit(titlekit.Judge(rules(), p, "探险家", &titlekit.Counter{}))
	case "prereq":
		p := &titlekit.Player{Level: 40, Achievements: 40, Hours: 40}
		emit(titlekit.Judge(rules(), p, "老兵", &titlekit.Counter{}))
	case "event":
		p := &titlekit.Player{Level: 9, Achievements: 9, Hours: 30}
		emit(titlekit.Judge(rules(), p, "老兵·活动", &titlekit.Counter{}))
	case "replay":
		p := &titlekit.Player{Level: 40, Achievements: 40, Hours: 40}
		c := &titlekit.Counter{}
		titlekit.Judge(rules(), p, "新兵", c)
		second := titlekit.Judge(rules(), p, "新兵", c)
		emit(map[string]any{"status": second.Status, "unlocked": len(p.Unlocked), "scanned": c.Scanned})
	case "work":
		p := &titlekit.Player{Level: 40, Achievements: 40, Hours: 40}
		c := &titlekit.Counter{}
		for i := 0; i < 4000; i++ {
			titlekit.Judge(rules(), p, "探险家", c)
		}
		emit(map[string]int{"judgements": 4000, "scanned": c.Scanned})
	case "":
		flag.Usage()
		os.Exit(2)
	default:
		fmt.Fprintln(os.Stderr, "未知场景："+*sample)
		os.Exit(2)
	}
}
