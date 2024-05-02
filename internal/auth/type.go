package auth

import "time"

type KakaoAccount struct {
	Profile                       Profile `json:"profile"`
	ProfileNicknameNeedsAgreement bool    `json:"profile_nickname_needs_agreement"`
	ProfileImageNeedsAgreement    bool    `json:"profile_image_needs_agreement"`
}

type Profile struct {
	Nickname          string `json:"nickname"`
	IsDefaultNickname bool   `json:"is_default_nickname"`
}

type KakaoUser struct {
	ConnectedAt  time.Time    `json:"connected_at"`
	Properties   Properties   `json:"properties"`
	KakaoAccount KakaoAccount `json:"kakao_account"`
	ID           int          `json:"id"`
}

type Properties struct {
	Nickname string `json:"nickname"`
}

type GetKakaoInfoResponse struct {
	AccessToken  string `json:"access_key"`
	RefreshToken string `json:"refresh_token"`
}
