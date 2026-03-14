package biz

import (
	"XiaXiaoMan/core/botEngine"
	"context"
)

type BuiltInCommand interface {
	GetName() string
	GetDescription() string
	GetCaller() string
	GetHandlerFunc() func(ctx context.Context, c *botEngine.RequestContext)
}

type magicCommand interface {
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
	BuiltInCommand map[string]BuiltInCommand
	MagicCommand   map[string]magicCommand
	LuaPlugin      map[string]luaPlugin
}
