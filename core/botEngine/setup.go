package botEngine

import (
	"XiaXiaoMan/core/config"
	"XiaXiaoMan/core/models/onebot"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const heartBeatInterval = 15

type BotEngine struct {
	conf         *config.Config
	conn         *websocket.Conn
	uuid         uuid.UUID
	mu           sync.Mutex
	pendingCalls map[string]chan *onebot.APIResponse
	eventChan    chan interface{}
	closeChan    chan struct{}
}

func SetupBotEngine(conf *config.Config) (*BotEngine, error) {

}
