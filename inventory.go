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

func (s *InventorySession) GetAllItemModifierGroups() (structures.GetAllItemModifierGroupsResp, error) {
	var resp structures.GetAllItemModifierGroupsResp

	body, err := apicall.MakeGetAllItemModifierGroupsReq(s.AccessSecret, s.MID, s.InventoryBaseUrl)
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

func (s *InventorySession) GetAllItemsOfModifierGroup(mgID string) (structures.GetAllItemsModifierGroupResp, error) {
	var resp structures.GetAllItemsModifierGroupResp

	body, err := apicall.MakeGetAllItemsModifierGroupReq(s.AccessSecret, s.MID, s.InventoryBaseUrl, mgID)
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

func (s *InventorySession) GetModifier(mgID, mdID string) (structures.ItemModifier, error) {
	var resp structures.ItemModifier

	body, err := apicall.MakeGetModifierReq(s.AccessSecret, s.MID, s.InventoryBaseUrl, mgID, mdID)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	if resp.Id == "" {
		return resp, errors.New(handleFailwithStatusOk(body))
	}

	return resp, nil
}
