package interfaces

type IGPTRepository interface {
	GPTChatRequest(string) (*string, error)
	GeminiChatRequest(string) (*string, error)
}
