package ping

import (
	"XiaXiaoMan/core/botEngine"
	"context"
)

type Command struct {
	Name        string
	Description string
	Caller      string
}

func (r *Command) GetName() string {
	return r.Name
}

func (r *Command) GetCaller() string {
	return r.Caller
}

func (r *Command) GetHandlerFunc() func(ctx context.Context, c *botEngine.RequestContext) {
	return r.Ping
}

func (r *Command) GetDescription() string {
	return r.Description
}

func Constructor(name string, description string, caller string) *Command {
	return &Command{Name: name, Description: description, Caller: caller}
}
