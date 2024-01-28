package goclover

type Service struct {
	ClientSecret string
	BaseUrl      string
}

type Session struct {
	Service
	AccessSecret string
	MID          string
}

type GetAccessTokenResp struct {
	AccessToken string `json:"access_token"`
}

type RespFailStruct struct {
	Message string `json:"message"`
}

type GetItemsRespElements struct {
	Id              string `json:"id"`
	Hidden          bool   `json:"hidden"`
	Available       bool   `json:"available"`
	AutoManage      bool   `json:"autoManage"`
	Name            string `json:"name"`
	AlternateName   string `json:"alternateName"`
	Code            string `json:"code"`
	Sku             string `json:"sku"`
	Price           int    `json:"price"`
	PriceType       string `json:"priceType"`
	DefaultTaxRates bool   `json:"defaultTaxRates"`
	UnitName        string `json:"unitName"`
	Cost            int    `json:"cost"`
	IsRevenue       bool   `json:"isRevenue"`
	StockCount      int    `json:"stockCount"`
	ModifiedTime    int64  `json:"modifiedTime"`
	Deleted         bool   `json:"deleted"`
}

type GetListOfItemsResp struct {
	Elements []GetItemsRespElements `json:"elements"`
}

type GetCategoriesRespElements struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
	Deleted   bool   `json:"deleted"`
}

type GetItemCategoriesResp struct {
	Elements []GetCategoriesRespElements `json:"elements"`
}
