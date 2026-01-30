package test

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"

	"gf_template/internal/consts"
	"gf_template/internal/library/cron"
	"gf_template/internal/model/entity"
)

func init() {
	cron.Register(cron_t)
	cron.Start(&entity.SysCron{
		Name:    "test",
		Id:      1,
		Pattern: "*/3 * * * * *",
		Policy:  int64(consts.CronPolicySame),
	})
	time.Sleep(15 * time.Second)
	cron.Stop(&entity.SysCron{
		Name:    "test",
		Id:      1,
		Pattern: "*/3 * * * * *",
		Policy:  int64(consts.CronPolicySame),
	})
}

var cron_t = &cronT{Name: "test", Id: 1}

type cronT struct {
	Id      int64
	Name    string
	Pattern string
}

func (c *cronT) GetName() string {
	return c.Name
}

func (c *cronT) Execute(ctx context.Context, parser *cron.Parser) (err error) {
	g.Log().Info(ctx, "\n================== 定时任务开启 =====================")
	return
}
