package structures

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

type LineItem struct {
	Item struct {
		Id string `json:"id"`
	} `json:"item"`
	Printed   string `json:"printed"`
	Exchanged string `json:"exchanged"`
	Refunded  string `json:"refunded"`
	Refund    struct {
		TransactionInfo struct {
			IsTokenBasedTx string `json:"isTokenBasedTx"`
			EmergencyFlag  string `json:"emergencyFlag"`
		} `json:"transactionInfo"`
	} `json:"refund"`
	IsRevenue                                  string `json:"isRevenue"`
	Id                                         string `json:"id"`
	ColorCode                                  string `json:"colorCode"`
	Name                                       string `json:"name"`
	AlternateName                              string `json:"alternateName"`
	Price                                      int    `json:"price"`
	PriceWithModifiers                         int    `json:"priceWithModifiers"`
	PriceWithModifiersAndItemAndOrderDiscounts int    `json:"priceWithModifiersAndItemAndOrderDiscounts"`
	UnitQty                                    int    `json:"unitQty"`
	UnitName                                   string `json:"unitName"`
	ItemCode                                   string `json:"itemCode"`
	Note                                       string `json:"note"`
	BinName                                    string `json:"binName"`
	UserData                                   string `json:"userData"`
	Discounts                                  []struct {
		Id         string `json:"id"`
		Name       string `json:"name"`
		Amount     int    `json:"amount"`
		Percentage int    `json:"percentage"`
	} `json:"discounts"`
	DiscountAmount int `json:"discountAmount"`
}
