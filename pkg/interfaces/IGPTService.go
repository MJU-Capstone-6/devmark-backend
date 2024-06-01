package interfaces

type IGPTService interface {
	ExtractTitle(string) (*string, error)
	SummarizePost(string) (*string, error)
}
