package botEngine

import (
	"XiaXiaoMan/core"
	"XiaXiaoMan/core/plugins"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/interaction/webhook"
	"github.com/tencent-connect/botgo/token"
)

func InitEngine() {
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

	_ = event.RegisterHandlers(
		GroupATMessageEventHandler(),
	)

	go func() {
		r := gin.Default()
		r.Any(core.Conf.CallBackPath, func(c *gin.Context) {
			webhook.HTTPHandler(c.Writer, c.Request, core.Credentials)
		})
		r.Run(fmt.Sprintf("%s:%s", core.Conf.ListenIPAddr, core.Conf.ListenPort))
	}()
}

func GroupATMessageEventHandler() event.GroupATMessageEventHandler {
	return func(event *dto.WSPayload, data *dto.WSGroupATMessageData) error {
		return plugins.ProcessGroupMessage(data)
	}
}
