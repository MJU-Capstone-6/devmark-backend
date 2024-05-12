package constants

import "time"

var (
	ACCESSTOKEN_EXPIRED_TIME   = time.Now().Add(24 * time.Hour)
	REFRESH_TOKEN_EXPIRED_TIME = time.Now().Add(24 * 7 * time.Hour)
	TOKEN_DATA_KEY             = "id"
	CODE_LENGTH                = 6
)
