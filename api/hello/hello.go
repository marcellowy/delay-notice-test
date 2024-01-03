// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import (
	"context"

	"delay-notice-test/api/hello/v1"
)

type IHelloV1 interface {
	Callback(ctx context.Context, req *v1.CallbackReq) (res *v1.CallbackRes, err error)
	Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error)
}
