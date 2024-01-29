package apicall

import (
	"bytes"
	"encoding/json"
	"github.com/aicam/goclover/structures"
	"io"
	"log"
	"net/http"
)

func MakeCreateAtomicOrderReq(order structures.CreateAtomicOrderReq, accessKey, mid, baseUrl string) ([]byte, error) {

	orderJson, err := json.Marshal(order)
	if err != nil {
		return []byte(""), err
	}

	req, err := http.NewRequest("POST", baseUrl+"v3/merchants/"+mid+"/atomic_order/orders", bytes.NewBuffer(orderJson))

	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Authorization", "Bearer "+accessKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}

func MakeGetOrderInfoReq(orderId, accessKey, mid, baseUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"v3/merchants/"+mid+"/orders/"+orderId, nil)
	if err != nil {
		log.Print(err)
		return []byte(""), err
	}
	req.Header.Add("Authorization", "Bearer "+accessKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}

func MakeGetOrderLineItemsReq(orderId, accessKey, mid, baseUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"v3/merchants/"+mid+"/orders/"+orderId+"/line_items", nil)
	if err != nil {
		log.Print(err)
		return []byte(""), err
	}
	req.Header.Add("Authorization", "Bearer "+accessKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}

func MakeChargeCardReq(amount int, currency, token string, accessKey, baseUrl string) ([]byte, error) {
	reqJson, _ := json.Marshal(structures.ChargeCardReq{
		Amount:   amount,
		Currency: currency,
		Source:   token,
	})

	req, err := http.NewRequest("POST", baseUrl+"v1/charges", bytes.NewBuffer(reqJson))

	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Authorization", "Bearer "+accessKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}

func MakePayOrderReq(reqStruct structures.PayOrderReq, orderId string, accessKey, baseUrl string) ([]byte, error) {
	reqJson, _ := MarshalJSONFiltered(reqStruct)
	req, err := http.NewRequest("POST", baseUrl+"v1/orders/"+orderId+"/pay", bytes.NewBuffer(reqJson))

	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Authorization", "Bearer "+accessKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []byte(""), err
	}

	return body, nil
}
