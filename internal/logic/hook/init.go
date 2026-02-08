package hook

import (
	"gf_template/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sHook struct {
}

func init() {
	service.RegisterHook(New())
}

func New() *sHook {
	return &sHook{}
}

func (s *sHook) BeforeServe(r *ghttp.Request) {

}

func (s *sHook) AfterOutput(r *ghttp.Request) {
	s.accessLog(r)
}
