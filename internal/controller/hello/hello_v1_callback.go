package hello

import (
	"context"
	"delay-notice-test/api/hello/v1"
)

func (c *ControllerV1) Callback(ctx context.Context, req *v1.CallbackReq) (res *v1.CallbackRes, err error) {
	res = &v1.CallbackRes{}

	res.Name = " your say name: " + req.Name

	return res, nil
}
