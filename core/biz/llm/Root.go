package llm

import (
	"XiaXiaoMan/core/botEngine"
	"XiaXiaoMan/core/config"
	"context"
	"sync"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
)

var S *ChatBot
var M *MemorySet

type ChatBot struct {
	mu      sync.RWMutex
	conf    *config.Config
	chatBot *openai.LLM
}

type Replier struct {
	chatBot *ChatBot
	c       *botEngine.RequestContext
	msg     *botEngine.MessageSegmentSet
	mem     *memory.ConversationBuffer
	ctx     context.Context
}

type MemorySet struct {
	mem map[int64]*memory.ConversationBuffer
}
