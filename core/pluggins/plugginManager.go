package pluggins

import (
	"XiaXiaoMan/core/pluggins/helloWorld"
	"strings"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/dto/message"
)

var pluggins []pluggin = []pluggin{
	&helloWorld.HelloWorld{MatchCommand: "/test"},
}

type pluggin interface {
	ProcessEvent(data *dto.WSGroupATMessageData) error
	ReturnMetchCommand() string
}

func ProcessGroupMessage(data *dto.WSGroupATMessageData) {
	raw := strings.ToLower(message.ETLInput(data.Content))
	commandSlice := strings.SplitN(raw, " ", 2)
	cmd := commandSlice[0]

	for _, p := range pluggins {
		if cmd == p.ReturnMetchCommand() {
			p.ProcessEvent(data)
		}
	}
}
