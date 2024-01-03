package hello

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"io"

	"delay-notice-test/api/hello/v1"
)

type HttpNotice struct {
	Url     string            `json:"url"`
	Timeout int32             `json:"timeout"` // 回调时请求超时时间,单位: 秒
	Header  map[string]string `json:"header"`  // 回调时头的参数,没有就不填
}

type AddNoticeReq struct {
	Type       string      `json:"type"`        // 被通知时的调用,默认: http|https
	Http       *HttpNotice `json:"http"`        // http通知方法
	Delay      int32       `json:"delay"`       // 延迟多久,单位:秒
	RetryTimes int32       `json:"retry_times"` // 通知失败时重试几次
	Data       string      `json:"data"`        // 通知时原样返回的数据
}

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {

	noticeReq := AddNoticeReq{
		Type: "http",
		Http: &HttpNotice{
			Url:     "http://127.0.0.1:8099/callback",
			Timeout: 3,
			Header:  nil,
		},
		Delay:      10,
		RetryTimes: 3,
		Data:       `{"name": "test name"}`,
	}

	b, err := json.Marshal(noticeReq)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return nil, err
	}

	client := gclient.New()
	response, err := client.Post(ctx, "http://127.0.0.1:8000/api/v1/notice/add", string(b))
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return nil, err
	}
	defer response.Body.Close()

	ccb, err := io.ReadAll(response.Body)
	if err != nil {
		g.Log().Line().Error(ctx, err)
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.Writeln(string(ccb))
	return
}
