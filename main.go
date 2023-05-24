package main

import (
	_ "goframe-shop-v2/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "goframe-shop-v2/internal/logic"

	"goframe-shop-v2/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
