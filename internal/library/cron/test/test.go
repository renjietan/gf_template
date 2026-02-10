package test

import (
	"context"
	"fmt"
	"time"

	"gf_template/internal/consts"
	"gf_template/internal/library/cron"
	"gf_template/internal/model/entity"
	"gf_template/utility/simple"
)

var count = 0
var entity_syscron = &entity.SysCron{
	Name:    Test.name,
	Pattern: "*/3 * * * * *",
	Params:  "123456,2226",
	Policy:  int64(consts.CronPolicyTimes),
	Count:   3,
}

func TestCron() {
	ctx := context.TODO()
	cron.Register(Test)
	simple.SafeGo(ctx, func(ctx context.Context) {
		cron.Start(entity_syscron)
	})
}

// Test 测试任务（无参数）
var Test = &cTest{name: "test11111111111+++++"}

type cTest struct {
	name string
}

func (c *cTest) GetName() string {
	return c.name
}

// Execute 执行任务
func (c *cTest) Execute(ctx context.Context, parser *cron.Parser) (err error) {
	count++
	fmt.Printf("parser.Arg%vs=========================%v", count, parser.Args)
	parser.Logger.Infof(ctx, "cron test Execute:%v", time.Now())
	if count > 5 {
		cron.Stop(entity_syscron)
	}
	return
}
