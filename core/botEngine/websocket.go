package botEngine

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func (c *BotEngine) connectWS() error {
	// 生成URL
	wsURL, err := url.Parse(c.conf.LLBotConfig.Url)
	if err != nil {
		return fmt.Errorf("parse url error: %v", err.Error())
	}
	// 设置头部
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprintf("Bearer %s", c.conf.LLBotConfig.Token))
	// 创建连接
	conn, _, err := websocket.DefaultDialer.Dial(wsURL.String(), headers)
	if err != nil {
		return fmt.Errorf("dialer connection error: %v", err.Error())
	}

	c.conn = conn
	return nil
}

func (c *BotEngine) readLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				select {
				case <-ctx.Done():
				default:
					fmt.Printf("Read Message Error: %v\n", err)
				}
				return
			}
			c.handleMessage(message)
		}
	}
}

func (c *BotEngine) closeWS() error {
	return c.conn.Close()
}
