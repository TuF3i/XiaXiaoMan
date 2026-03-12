package botEngine

import (
	"XiaXiaoMan/core/config"
	"XiaXiaoMan/core/models/onebot"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type BotEngine struct {
	conf         *config.Config
	conn         *websocket.Conn
	uuid         uuid.UUID
	mu           sync.Mutex
	pendingCalls map[string]chan *onebot.APIResponse
	eventChan    chan interface{}
	closeChan    chan struct{}
}

func SetupBotEngine(conf *config.Config) *BotEngine {
	return &BotEngine{
		conf:         conf,
		uuid:         uuid.New(),
		mu:           sync.Mutex{},
		pendingCalls: make(map[string]chan *onebot.APIResponse),
		eventChan:    make(chan interface{}),
		closeChan:    make(chan struct{}),
	}
}
