package sysconfig

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// AppName 应用名称
func AppName(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.appName", "gf_system_appName").String()
}

// Debug debug
func Debug(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.debug", true).Bool()
}

// TimeZone
func GetTimeZone(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.timeZone", "Asia/Shanghai").String()
}

// mode 获取运行模式
func GetMode(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.mode", "not-set").String()
}

// GetLanguage 获取系统默认语言
func GetLanguage(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.i18n.defaultLanguage", "zh-CN").String()
}

// mode 获取运行模式
func GetLangSwitch(ctx context.Context) bool {
	return g.Cfg().MustGet(ctx, "system.i18n.switch", true).Bool()
}

//
