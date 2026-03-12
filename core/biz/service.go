package biz

import (
	"XiaXiaoMan/core/botEngine"
	"context"
)

type HandlerFunc func(ctx context.Context, c *botEngine.BotEngine)

type builtInCommand interface {
	GetName() string
	GetCaller() string
	GetHandlerFunc() HandlerFunc
}

type luaPlugin interface {
	GetName() string
	GetCaller() string
	GetHandlerFunc() HandlerFunc
}

var ()
