package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
)

const ML_SERVER_URL = "https://ml.officialdevmark.com"

type PredictRequest struct {
	Title string `json:"title"`
}

type PredictResponse struct {
	Category string `json:"category"`
}

type PredictCategoryDTO struct {
	Title    string
	Category string
}

func ParseURLParam(ctx echo.Context, paramName string) (*int, error) {
	param, err := strconv.Atoi(ctx.Param(paramName))
	if err != nil {
		return nil, customerror.InvalidParamError(err)
	}
	return &param, nil
}

func GetAuthUser(ctx echo.Context) (*repository.FindUserByIdRow, error) {
	if user, ok := ctx.Get(constants.USER_CONTEXT_KEY).(*repository.FindUserByIdRow); ok {
		return user, nil
	}
	return nil, customerror.UserNotFound(errors.New(""))
}

func getTitle(link string, domain string) (*string, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	docs, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	switch domain {
	case "tistory":
		title := docs.Find(".tit_blogview").Text()
		if title == "" {
			title = docs.Find(".tit_post").Text()
		}
		if title == "" {
			title = docs.Find(".post-cover > .inner > h1").Text()
		}
		if title == "" {
			title = docs.Find(".hgroup > h1").Text()
		}
		if title == "" {
			title = docs.Find(".description > h1").Text()
		}
		if title == "" {
			title = docs.Find(".title_view").Text()
		}
		if title == "" {
			title = docs.Find(".title-article").Text()
		}
		return &title, nil
	case "medium":
		title := docs.Find("h1").Text()
		return &title, nil
	case "github":
		title := docs.Find("[data-toc-skip]").Text()
		return &title, nil
	}
	return nil, nil
}

func GetJsonString(str string) string {
	re := regexp.MustCompile(`\{[\s\S]*\}`)
	jsonString := re.FindString(str)
	return jsonString
}

func GetResponseBody(link string) (*string, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	htmlString, err := doc.Find("body").Html()
	if err != nil {
		return nil, err
	}
	return &htmlString, nil
}

func predictCategory(title string) (*PredictResponse, error) {
	predictReq := PredictRequest{
		Title: title,
	}
	jsonBody, err := json.Marshal(predictReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", ML_SERVER_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var predictResponse PredictResponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &predictResponse)
	if err != nil {
		return nil, err
	}
	return &predictResponse, nil
}

func PredictCategoryRequest(title string) (*PredictCategoryDTO, error) {
	category, err := predictCategory(title)
	if err != nil {
		return nil, err
	}
	dto := PredictCategoryDTO{
		Category: category.Category,
	}
	return &dto, nil
}

func SliceValueIntoNum(arr []string) (*[]int64, error) {
	intSlice := make([]int64, len(arr))
	for index, value := range arr {
		num, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		intSlice[index] = int64(num)
	}
	return &intSlice, nil
}
