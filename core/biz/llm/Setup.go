package llm

import (
	"XiaXiaoMan/core/botEngine"
	"XiaXiaoMan/core/config"
	"context"
	"sync"

	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/memory"
)

func setupChatBot(conf *config.Config) error {
	llmConnectionPool, err := openai.New(
		openai.WithModel(conf.LLMChatConfig.ChatModel),
		openai.WithBaseURL(conf.LLMChatConfig.AliyunBaseURL),
		openai.WithToken(conf.LLMChatConfig.AliyunAPIKey),
	)
	if err != nil {
		return err
	}

	S = &ChatBot{mu: sync.RWMutex{}, conf: conf, chatBot: llmConnectionPool}

	return nil
}

func setupMemory() {
	M = &MemorySet{map[int64]*memory.ConversationBuffer{}}
}

func GetReplier(c *botEngine.RequestContext, ctx context.Context) *Replier {
	if mem, ok := M.mem[c.GetGroupMessageEvent().GroupID]; ok {
		return &Replier{chatBot: S, c: c, mem: mem, ctx: ctx, msg: nil}
	}

	mem := memory.NewConversationBuffer(memory.WithReturnMessages(true))
	M.mem[c.GetGroupMessageEvent().GroupID] = mem
	return &Replier{chatBot: S, c: c, mem: mem, ctx: ctx, msg: nil}
}
