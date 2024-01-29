package goclover

import (
	"encoding/json"
	"errors"
	"github.com/aicam/goclover/apicall"
	"github.com/aicam/goclover/structures"
)

func (s *InventorySession) CreateAtomicOrder(order structures.CreateAtomicOrderReq) (structures.CreateAtomicOrderResp, error) {
	var resp structures.CreateAtomicOrderResp

	body, err := apicall.MakeCreateAtomicOrderReq(order, s.AccessSecret, s.MID, s.InventoryBaseUrl)
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

func (s *InventorySession) GetOrderInfo(id string) (structures.GetOrderInfoResp, error) {
	var resp structures.GetOrderInfoResp

	body, err := apicall.MakeGetOrderInfoReq(id, s.AccessSecret, s.MID, s.InventoryBaseUrl)
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

func (s *InventorySession) GetOrderLineItems(id string) (structures.GetOrderLineItemsResp, error) {
	var resp structures.GetOrderLineItemsResp

	body, err := apicall.MakeGetOrderLineItemsReq(id, s.AccessSecret, s.MID, s.InventoryBaseUrl)
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

func (s *InventorySession) ChargeCard(amount int, currency, token string) (structures.ChargeCardResp, error) {
	var resp structures.ChargeCardResp

	body, err := apicall.MakeChargeCardReq(amount, currency, token, s.AccessSecret, s.EcommerceBaseUrl)
	if err != nil {
		return resp, err
	}

	json.Unmarshal(body, &resp)
	if resp.Id == "" {
		return resp, errors.New(handleFailwithStatusOk(body))
	}

	return resp, nil
}

func (s *InventorySession) PayOrder(payObj structures.PayOrderReq, orderId string) (structures.PayOrderResp, error) {
	var resp structures.PayOrderResp

	body, err := apicall.MakePayOrderReq(payObj, orderId, s.AccessSecret, s.EcommerceBaseUrl)
	if err != nil {
		return resp, err
	}

	json.Unmarshal(body, &resp)
	if resp.Id == "" {
		return resp, errors.New(handleFailwithStatusOk(body))
	}

	return resp, nil
}
