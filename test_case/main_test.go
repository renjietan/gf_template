package main

import (
	"context"
	"testing"

	"github.com/gogf/gf/v2/os/glog"
)

func TestContext(t *testing.T) {
	ctx := context.TODO()
	l := glog.New()
	l.SetFlags(glog.F_TIME_TIME | glog.F_FILE_SHORT)
	l.Print(ctx, "time and short line number")
	l.SetFlags(glog.F_TIME_MILLI | glog.F_FILE_LONG)
	l.Print(ctx, "time with millisecond and long line number")
	l.SetFlags(glog.F_TIME_STD | glog.F_FILE_LONG)
	l.Print(ctx, "standard time format and long line number")
}
