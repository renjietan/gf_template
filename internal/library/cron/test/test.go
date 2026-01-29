package test

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"gf_template/internal/library/cron"
	"gf_template/internal/model/entity"
)

func init() {
	cron.Register(cron_t)
	cron.Start(&entity.SysCron{
		Name: "test",
		Id:   1,
	})
}

var cron_t = &cronT{name: "test"}

type cronT struct {
	name string
}

func (c *cronT) GetName() string {
	return c.name
}

func (c *cronT) Execute(ctx context.Context, parser *cron.Parser) (err error) {
	g.Log().Info(ctx, "================== 定时任务开启 =====================")
	return
}
