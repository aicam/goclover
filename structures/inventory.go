package structures

type GetListOfItemsResp struct {
	Elements []LineItem `json:"elements"`
}

type LineCategory struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sortOrder"`
	Deleted   bool   `json:"deleted"`
}

type GetItemCategoriesResp struct {
	Elements []LineCategory `json:"elements"`
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
	IsRevenue                                  bool   `json:"isRevenue"`
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

type ItemModifierGroup struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	MinRequired   int    `json:"minRequired"`
	MaxAllowed    int    `json:"maxAllowed"`
	ShowByDefault bool   `json:"showByDefault"`
	ModifierIds   string `json:"modifierIds"`
	Deleted       bool   `json:"deleted"`
}

type GetAllItemModifierGroupsResp struct {
	Elements []ItemModifierGroup `json:"elements"`
}

type ItemModifier struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Available     bool   `json:"available"`
	Price         int    `json:"price"`
	ModifiedTime  int64  `json:"modifiedTime"`
	ModifierGroup struct {
		Id string `json:"id"`
	} `json:"modifierGroup"`
	Deleted bool `json:"deleted"`
}

type GetAllItemsModifierGroupResp struct {
	Elements []LineItem `json:"elements"`
}
