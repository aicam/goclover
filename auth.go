package goclover

import (
	"encoding/json"
	"errors"
	"github.com/aicam/goclover/apicall"
	"github.com/aicam/goclover/structures"
)

func (s *Service) GetAccessToken(clientId, code string) (string, error) {

	body, err := apicall.MakeGetAccessTokenReq(clientId, code, s.ClientSecret, s.InventoryBaseUrl)
	if err != nil {
		return "", err
	}

	var respJson structures.GetAccessTokenResp

	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return "", err
	}

	if respJson.AccessToken == "" {
		return "", errors.New(handleFailwithStatusOk(body))
	}

	return respJson.AccessToken, nil
}
