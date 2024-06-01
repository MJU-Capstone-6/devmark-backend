package interfaces

type IGPTService interface {
	GPTExtractTitle(string) (*string, error)
	GeminiExtractTitle(string) (*string, error)
	GPTSummarizePost(string) (*string, error)
	GeminiSummarizePost(string) (*string, error)
}
