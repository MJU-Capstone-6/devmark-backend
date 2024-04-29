package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func requestKakaoUserInfo(key string) (*KakaoUser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", BASE_URL, nil)
	if err != nil {
		return nil, err
	}
	bearerToken := fmt.Sprintf("Bearer %s", key)
	req.Header.Set("Authorization", bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var user KakaoUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
