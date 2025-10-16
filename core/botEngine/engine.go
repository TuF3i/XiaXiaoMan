package botEngine

import (
	"XiaXiaoMan/core"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/interaction/webhook"
	"github.com/tencent-connect/botgo/token"
)

func initEngine() {
	tokenSource := token.NewQQBotTokenSource(core.Credentials)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := token.StartRefreshAccessToken(ctx, tokenSource)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	engine := botgo.NewOpenAPI(core.Credentials.AppID, tokenSource).WithTimeout(5 * time.Second).SetDebug(core.Conf.EnableDebug)
	core.Engine = engine

	go func() {
		r := gin.Default()
		r.Any(core.Conf.CallBackPath, func(c *gin.Context) {
			webhook.HTTPHandler(c.Writer, c.Request, core.Credentials)
		})
		r.Run(fmt.Sprintf("%s:%s", core.Conf.ListenIPAddr, core.Conf.ListenPort))
	}()
}
