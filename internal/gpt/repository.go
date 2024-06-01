package gpt

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/sashabaranov/go-openai"
)

type GPTRepository struct {
	GPTClient    *openai.Client
	GeminiClient *genai.GenerativeModel
}

func (g *GPTRepository) GPTChatRequest(content string) (*string, error) {
	resp, err := g.GPTClient.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
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

func (g *GPTRepository) GeminiChatRequest(content string) (*string, error) {
	resp, err := g.GeminiClient.GenerateContent(context.Background(), genai.Text(content))
	if err != nil {
		return nil, err
	}
	result := fmt.Sprintf("%+v", resp.Candidates[0].Content.Parts[0])
	return &result, nil
}

func InitGPTRepository() *GPTRepository {
	return &GPTRepository{}
}

func (g GPTRepository) WithClient(client *openai.Client) GPTRepository {
	g.GPTClient = client
	return g
}

func (g GPTRepository) WithGeminiModel(model *genai.GenerativeModel) GPTRepository {
	g.GeminiClient = model
	return g
}
