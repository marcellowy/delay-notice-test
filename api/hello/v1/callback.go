// Package v1
// Copyright 2016-2024 chad.wang<chad.wang@icloudsky.com>. All rights reserved.
package v1

import "github.com/gogf/gf/v2/frame/g"

type CallbackReq struct {
	g.Meta `path:"/callback" tags:"Hello" method:"post" summary:""`
	Name   string `json:"name"`
}
type CallbackRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Name   string `json:"name"`
}
