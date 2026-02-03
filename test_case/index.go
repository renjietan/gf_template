package main

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

type cMain struct {
	g.Meta `name:"main"`
}

type cMainHttpInput struct {
	g.Meta `name:"http" brief:"start http server"`
}
type cMainHttpOutput struct{}

type cMainGrpcInput struct {
	g.Meta `name:"grpc" brief:"start grpc server"`
}
type cMainGrpcOutput struct{}

func (c *cMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainHttpOutput, err error) {
	fmt.Println("start http server")
	return
}

func (c *cMain) Grpc(ctx context.Context, in cMainGrpcInput) (out *cMainGrpcOutput, err error) {
	fmt.Println("start grpc server")
	return
}

func main() {
	cmd, err := gcmd.NewFromObject(cMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}
