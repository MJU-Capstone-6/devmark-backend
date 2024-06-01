package gpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type GPTRepository struct {
	Client *openai.Client
}

func (g *GPTRepository) ChatRequest(content string) (*string, error) {
	resp, err := g.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: content,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return &resp.Choices[0].Message.Content, nil
}

func InitGPTRepository() *GPTRepository {
	return &GPTRepository{}
}

func (g GPTRepository) WithClient(client *openai.Client) GPTRepository {
	g.Client = client
	return g
}
