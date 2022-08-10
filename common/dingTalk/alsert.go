package dingTalk

import (
	"fmt"
	"github.com/blinkbean/dingtalk"
)

func WarnDayTaskToDing(TaskList []map[string]string) {

	// 单个机器人有单位时间内消息条数的限制，如果有需要可以初始化多个token，发消息时随机发给其中一个机器人。
	var (
		dingToken = "aea0b975929e75f86b87ba62145f1e7594188bf4984fcad5f3580426671e8a48"
		secret    = "SEC180120e837b23042613815f941b985b22edbc8e2a898b7ebc4996694a7ab421c"
	)
	cli := dingtalk.InitDingTalkWithSecret(dingToken, secret)

	dm := dingtalk.DingMap()
	dm.Set("今日任务", dingtalk.H2)

	for _, task := range TaskList {
		dm.Set(task["task"], dingtalk.N)
	}
	fmt.Print(TaskList)
	cli.SendMarkDownMessageBySlice("color test", dm.Slice())
}
