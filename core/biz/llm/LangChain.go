package llm

import (
	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/tools"
)

func (r *Replier) Do() {
	toolList := []tools.Tool{
		tools.Calculator{},
	}
	agent := agents.NewConversationalAgent(
		r.chatBot.chatBot,
		toolList,
		agents.WithMaxIterations(5),
		agents.WithReturnIntermediateSteps(),
	)
	executor := agents.NewExecutor(agent, agents.WithMemory(r.mem))

	messages := []llms.MessageContent{
		{
			Role:  llms.ChatMessageTypeSystem,
			Parts: []llms.ContentPart{llms.TextPart(systemPrompt)},
		},
		{
			Role:  llms.ChatMessageTypeHuman,
			Parts: []llms.ContentPart{llms.TextPart()},
		},
	}

	_, err := chains.Run(r.ctx, executor, messages)
}
