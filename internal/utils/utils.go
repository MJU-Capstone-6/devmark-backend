package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo/v4"
)

const ML_SERVER_URL = "https://ml.officialdevmark.com"

type PredictResponse struct {
	Category string `json:"category"`
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

func GetTitle(link string, domain string) (*string, error) {
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

func PredictCategory(title string) (*PredictResponse, error) {
	req, err := http.NewRequest("POST", ML_SERVER_URL, nil)
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
