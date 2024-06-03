package customerror

import "net/http"

var RECOMMEND_LINK_CREATION_FAIL = NewError(http.StatusUnprocessableEntity, "Recommend Link creation fail.", "1201")

func RecommendLinkCreationFail(err error) error {
	recommendLinkCreationFail := RECOMMEND_LINK_CREATION_FAIL
	recommendLinkCreationFail.Data = err
	return recommendLinkCreationFail
}
