package goclover

import (
	"encoding/json"
	"errors"
	"github.com/aicam/goclover/apicall"
	"github.com/aicam/goclover/structures"
)

func (s *InventorySession) GetListOfItems() (structures.GetListOfItemsResp, error) {
	var resp structures.GetListOfItemsResp

	body, err := apicall.MakeGetListOfItemsReq(s.AccessSecret, s.MID, s.InventoryBaseUrl)
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

func (s *InventorySession) GetItemCategories(itemId string) (structures.GetItemCategoriesResp, error) {
	var resp structures.GetItemCategoriesResp

	body, err := apicall.MakeGetItemCategoriesReq(s.AccessSecret, s.MID, itemId, s.InventoryBaseUrl)
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
