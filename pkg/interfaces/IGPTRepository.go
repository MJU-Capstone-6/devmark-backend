package interfaces

type IGPTRepository interface {
	ChatRequest(string) (*string, error)
}
