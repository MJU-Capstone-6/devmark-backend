package constants

import "time"

var (
	ACCESSTOKEN_EXPIRED_TIME   = time.Now().Add(24 * time.Hour)
	REFRESH_TOKEN_EXPIRED_TIME = time.Now().Add(24 * 7 * time.Hour)
	TOKEN_DATA_KEY             = "id"
	USER_CONTEXT_KEY           = "user"
	CODE_LENGTH                = 6
	USER_AGENT_HEADER          = "User-Agent"
)
