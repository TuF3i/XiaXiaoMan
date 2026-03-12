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

	handlerFunc map[string]HandleFunc
}

func SetupBotEngine(conf *config.Config) *BotEngine {
	engine := &BotEngine{
		conf:         conf,
		uuid:         uuid.New(),
		mu:           sync.Mutex{},
		pendingCalls: make(map[string]chan *onebot.APIResponse),
		eventChan:    make(chan interface{}),
		closeChan:    make(chan struct{}),
		handlerFunc:  make(map[string]HandleFunc),
	}

	return engine
}

func (c *BotEngine) Close() error {
	close(c.closeChan)
	return c.conn.Close()
}

func (c *BotEngine) Spin() error {
	// 创建ws连接
	if err := c.connectWS(); err != nil {
		return err
	}
	// 启动事件监听循环
	c.readLoop()
	// 阻塞
	return nil
}
