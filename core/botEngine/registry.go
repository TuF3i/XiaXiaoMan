package botEngine

import "context"

type HandleFunc func(ctx context.Context, c *BotEngine)
