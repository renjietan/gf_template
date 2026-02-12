package sys

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"gf_template/internal/consts"
	"gf_template/internal/library/contexts"
	"gf_template/internal/model"
	"gf_template/internal/service"
	sysconfig "gf_template/utility/config"
	"gf_template/utility/validate"
)

type sSysLog struct{}

func NewSysLog() *sSysLog {
	return &sSysLog{}
}

func init() {
	service.RegisterSysLog(NewSysLog())
}

// AutoLog 根据配置自动记录请求日志
func (s *sSysLog) AutoLog(ctx context.Context) error {
	return g.Try(ctx, func(ctx context.Context) {
		var err error
		defer func() {
			if err != nil {
				g.Log().Error(ctx, "autoLog 报错:%+v", err)
			}
		}()

		config, err := sysconfig.GetLog(ctx)
		if err != nil {
			return
		}

		if config == nil || !config.Switch {
			return
		}

		data := s.AnalysisLog(ctx)

		if ok := validate.InSliceExistStr(config.SkipCode, gconv.String(data.ErrorCode)); ok {
			return
		}
		// fmt.Println("\ndata======================", gconv.String(data))
		g.Log().Infof(ctx, "\nsys log 数据: \n%+v", gconv.String(data))
		// TODO: 这里是否有必要丢入队列
		// TODO：根据业务需求 是否需要存入数据库
	})
}

// AnalysisLog 解析日志数据
func (s *sSysLog) AnalysisLog(ctx context.Context) model.SysLog {

	var (
		mctx       = contexts.Get(ctx)
		response   = mctx.Response
		user       = mctx.User
		request    = ghttp.RequestFromCtx(ctx)
		postData   = gjson.New(consts.NilJsonToString)
		getData    = gjson.New(request.URL.Query())
		headerData = gjson.New(consts.NilJsonToString)
		errorData  = gjson.New(consts.NilJsonToString)
		data       model.SysLog
		userId     int64
		errorCode  int
		errorMsg   string
		traceID    string
		timestamp  int64
		takeUpTime int64
	)
	// 响应数据
	if response != nil {
		errorCode = response.Code
		errorMsg = response.Message
		traceID = response.TraceID
		timestamp = response.Timestamp
		if len(gconv.String(response.Error)) > 0 {
			errorData = gjson.New(response.Error)
		}
	}

	if timestamp == 0 {
		timestamp = gtime.Timestamp()
	}

	// 请求头
	if reqHeadersBytes, _ := gjson.New(request.Header).MarshalJSON(); len(reqHeadersBytes) > 0 {
		headerData = gjson.New(reqHeadersBytes)
	}

	// post参数
	if body, ok := mctx.Data["request.body"].(*gjson.Json); ok {
		postData = body
	}

	// post表单
	postForm := gjson.New(gconv.String(request.PostForm)).Map()
	if len(postForm) > 0 {
		for k, v := range postForm {
			postData.MustSet(k, v)
		}
	}

	if postData.IsNil() || len(postData.Map()) == 0 {
		postData = gjson.New(consts.NilJsonToString)
	} else {
		// 隐藏明文显示的密码
		for k := range postData.Map() {
			if gstr.ContainsI(k, "password") {
				postData.MustSet(k, "******")
			}
		}
	}

	// 当前登录用户
	if user != nil {
		userId = gconv.Int64(user.Id)
	}

	// 请求耗时
	if tt, ok := mctx.Data["request.waitTime"].(int64); ok {
		takeUpTime = tt
	}
	headerData.MustAppend("Enter-Time", request.EnterTime.String())

	data = model.SysLog{
		UserId:     userId,
		Method:     request.Method,
		Url:        request.URL.Path,
		GetData:    getData,
		PostData:   postData,
		HeaderData: s.SimplifyHeaderParams(headerData),
		ErrorCode:  errorCode,
		ErrorMsg:   errorMsg,
		ErrorData:  errorData,
		ReqId:      traceID,
		Timestamp:  timestamp,
		UserAgent:  request.Header.Get("User-Agent"),
		Status:     consts.StatusEnabled,
		TakeUpTime: takeUpTime,
		UpdatedAt:  gtime.Now(),
		CreatedAt:  request.EnterTime,
	}
	return data
}

func (s *sSysLog) SimplifyHeaderParams(data *gjson.Json) *gjson.Json {
	if data == nil || data.IsNil() {
		return data
	}
	var filters = []string{"Accept", "Authorization", "Cookie"}
	for _, filter := range filters {
		v := data.Get(filter)
		if len(v.String()) > 128 {
			data.MustRemove(filter)
		}
	}
	return data
}
