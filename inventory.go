package goclover

import (
	"encoding/json"
	"errors"
	"github.com/aicam/goclover/apicall"
)

func (s *Session) GetListOfItems() (GetListOfItemsResp, error) {
	var resp GetListOfItemsResp

	body, err := apicall.MakeGetListOfItemsReq(s.AccessSecret, s.MID, s.BaseUrl)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Elements) == 0 {
		return resp, errors.New(handleFailwithStatusOk(body))
	}

	return resp, nil
}

func (s *Session) GetItemCategories(itemId string) (GetItemCategoriesResp, error) {
	var resp GetItemCategoriesResp

	body, err := apicall.MakeGetItemCategoriesReq(s.AccessSecret, s.MID, itemId, s.BaseUrl)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Elements) == 0 {
		return resp, errors.New(handleFailwithStatusOk(body))
	}

	return resp, nil
}
