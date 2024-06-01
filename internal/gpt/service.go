package gpt

import (
	"encoding/json"
	"fmt"

	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/utils"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type GPTService struct {
	Repository interfaces.IGPTRepository
}

func (g *GPTService) GPTExtractTitle(link string) (*string, error) {
	var result TitleBody
	body, err := utils.GetResponseBody(link)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	resp, err := g.Repository.GPTChatRequest(fmt.Sprintf(GET_TITLE, *body))
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	jsonBody := utils.GetJsonString(*resp)
	err = json.Unmarshal([]byte(jsonBody), &result)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	return &result.Title, nil
}

func (g *GPTService) GeminiExtractTitle(link string) (*string, error) {
	var result TitleBody
	body, err := utils.GetResponseBody(link)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	resp, err := g.Repository.GeminiChatRequest(fmt.Sprintf(GET_TITLE, *body))
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	jsonBody := utils.GetJsonString(*resp)
	err = json.Unmarshal([]byte(jsonBody), &result)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	return &result.Title, nil
}

func (g *GPTService) GPTSummarizePost(link string) (*string, error) {
	var result ContextBody
	body, err := utils.GetResponseBody(link)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	resp, err := g.Repository.GPTChatRequest(fmt.Sprintf(POST_SUMMARY, *body))
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	jsonBody := utils.GetJsonString(*resp)
	err = json.Unmarshal([]byte(jsonBody), &result)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	return &result.Context, nil
}

func (g *GPTService) GeminiSummarizePost(link string) (*string, error) {
	var result ContextBody
	body, err := utils.GetResponseBody(link)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	resp, err := g.Repository.GeminiChatRequest(fmt.Sprintf(POST_SUMMARY, *body))
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	jsonBody := utils.GetJsonString(*resp)
	err = json.Unmarshal([]byte(jsonBody), &result)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	return &result.Context, nil
}

func InitGPTService() *GPTService {
	return &GPTService{}
}

func (g GPTService) WithRepository(repo interfaces.IGPTRepository) GPTService {
	g.Repository = repo
	return g
}
