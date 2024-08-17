package apicall

import (
	"io"
	"log"
	"net/http"
)

func MakeGetListOfItemsReq(accessKey, mid, baseUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"v3/merchants/"+mid+"/items", nil)
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

func MakeGetItemCategoriesReq(accessKey, mid, itemId, baseUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"v3/merchants/"+mid+"/items/"+itemId+"/categories", nil)
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

func MakeGetAllItemModifierGroupsReq(accessKey, mid, baseUrl string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"v3/merchants/"+mid+"/modifier_groups/", nil)
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

func MakeGetAllItemsModifierGroupReq(accessKey, mid, baseUrl, mgID string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"v3/merchants/"+mid+"/modifier_groups/"+mgID+"/items", nil)
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

func MakeGetModifierReq(accessKey, mid, baseUrl, mgID, mdID string) ([]byte, error) {
	req, err := http.NewRequest("GET", baseUrl+"v3/merchants/"+mid+"/modifier_groups/"+mgID+"/modifiers/"+mdID, nil)
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
