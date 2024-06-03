package constants

import "time"

var (
	ACCESSTOKEN_EXPIRED_TIME           = time.Now().Add(24 * time.Hour)
	REFRESH_TOKEN_EXPIRED_TIME         = time.Now().Add(24 * 7 * time.Hour)
	TOKEN_DATA_KEY                     = "id"
	USER_CONTEXT_KEY                   = "user"
	CODE_LENGTH                        = 6
	USER_AGENT_HEADER                  = "User-Agent"
	GEMINI_FLASH_MODEL                 = "gemini-1.5-flash"
	UNREAD_BOOKMARK_NOTIFIACTION_TITLE = "%s에서 읽지 않은 북마크가 %d개 있습니다."
	UNREAD_BOOKMARK_NOTIFIACTION_BODY  = "북마크 제목 : %s"
)
