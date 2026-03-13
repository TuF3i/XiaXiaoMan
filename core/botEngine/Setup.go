package botEngine

import (
	"XiaXiaoMan/core/config"
	"XiaXiaoMan/core/models/onebot"
	"sync"

	"github.com/google/uuid"
)

func (c *BotEngine) Close() error {
	// 发送关闭广播
	close(c.closeChan)
	// 等待监听器关闭
	c.wg.Wait()
	// 关闭事件管道
	close(c.eventChan)
	// 关闭ws
	return c.conn.Close()
}

func (c *BotEngine) Spin() error {
	// 创建ws连接
	if err := c.connectWS(); err != nil {
		return err
	}
	// 启动LLBot事件监听器
	go c.llBotEventListener()
	// 启动本地事件监听循环
	go c.localEventListener()

	return nil
}

func SetupBotEngine(conf *config.Config) *BotEngine {
	engine := &BotEngine{
		conf: conf,
		uuid: uuid.New(),
		mu:   sync.Mutex{},
		wg:   sync.WaitGroup{},

		pendingCalls: make(map[string]chan *onebot.APIResponse),
		eventChan:    make(chan interface{}),
		closeChan:    make(chan struct{}),
		handlerFunc:  make(map[string]HandleFunc),
	}

	return engine
}
