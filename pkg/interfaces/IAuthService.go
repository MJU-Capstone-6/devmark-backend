package interfaces

import (
	"github.com/MJU-Capstone-6/devmark-backend/internal/responses"
)

//go:generate mockery --name IAuthService
type IAuthService interface {
	KakaoSignUp(string, string) (*responses.GetKakaoInfoResponse, error)
}
