package biz

import (
	"XiaXiaoMan/core/botEngine"
	"context"
)

type builtInCommand interface {
	GetName() string
	GetDescription() string
	GetCaller() string
	GetHandlerFunc() func(ctx context.Context, c *botEngine.RequestContext)
}

type luaPlugin interface {
	GetName() string
	GetCaller() string
	GetDescription() string
	GetHandlerFunc() func(ctx context.Context, c *botEngine.RequestContext)
}

type Biz struct {
}
