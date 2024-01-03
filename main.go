package main

import (
	_ "delay-notice-test/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"delay-notice-test/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
