package structures

type CreateAtomicOrderResp struct {
	Href     string `json:"href"`
	Id       string `json:"id"`
	Currency string `json:"currency"`
	Employee struct {
		Href   string `json:"href"`
		Id     string `json:"id"`
		Name   string `json:"name"`
		Role   string `json:"role"`
		Orders struct {
			Href string `json:"href"`
		} `json:"orders"`
	} `json:"employee"`
	Total     int `json:"total"`
	OrderType struct {
		Id                string `json:"id"`
		LabelKey          string `json:"labelKey"`
		Label             string `json:"label"`
		Taxable           bool   `json:"taxable"`
		IsDefault         bool   `json:"isDefault"`
		FilterCategories  bool   `json:"filterCategories"`
		IsHidden          bool   `json:"isHidden"`
		Fee               int    `json:"fee"`
		MinOrderAmount    int    `json:"minOrderAmount"`
		MaxOrderAmount    int    `json:"maxOrderAmount"`
		MaxRadius         int    `json:"maxRadius"`
		AvgOrderTime      int    `json:"avgOrderTime"`
		HoursAvailable    string `json:"hoursAvailable"`
		IsDeleted         bool   `json:"isDeleted"`
		SystemOrderTypeId string `json:"systemOrderTypeId"`
	} `json:"orderType"`
	TaxRemoved        bool   `json:"taxRemoved"`
	IsVat             bool   `json:"isVat"`
	State             string `json:"state"`
	ManualTransaction bool   `json:"manualTransaction"`
	GroupLineItems    bool   `json:"groupLineItems"`
	TestMode          bool   `json:"testMode"`
	CreatedTime       int64  `json:"createdTime"`
	ClientCreatedTime int64  `json:"clientCreatedTime"`
	ModifiedTime      int64  `json:"modifiedTime"`
	LineItems         struct {
		Elements []struct {
			Id       string `json:"id"`
			OrderRef struct {
				Id string `json:"id"`
			} `json:"orderRef"`
			Item struct {
				Id string `json:"id"`
			} `json:"item"`
			Name                   string `json:"name"`
			AlternateName          string `json:"alternateName"`
			Price                  int    `json:"price"`
			ItemCode               string `json:"itemCode"`
			Printed                bool   `json:"printed"`
			CreatedTime            int64  `json:"createdTime"`
			OrderClientCreatedTime int64  `json:"orderClientCreatedTime"`
			Exchanged              bool   `json:"exchanged"`
			Refunded               bool   `json:"refunded"`
			IsRevenue              bool   `json:"isRevenue"`
			TaxRates               struct {
				Elements []struct {
					Id          string `json:"id"`
					LineItemRef struct {
						Id string `json:"id"`
					} `json:"lineItemRef"`
					Name      string `json:"name"`
					Rate      int    `json:"rate"`
					IsDefault bool   `json:"isDefault"`
				} `json:"elements"`
			} `json:"taxRates"`
			IsOrderFee bool `json:"isOrderFee"`
		} `json:"elements"`
	} `json:"lineItems"`
}

type CreateAtomicOrderReq struct {
	OrderCart struct {
		OrderType struct {
			Id string `json:"id"`
		} `json:"orderType"`
		GroupLineItems string     `json:"groupLineItems"`
		LineItems      []LineItem `json:"lineItems"`
	} `json:"orderCart"`
}

type GetOrderInfoResp struct {
	Href     string `json:"href"`
	Id       string `json:"id"`
	Currency string `json:"currency"`
	Employee struct {
		Id string `json:"id"`
	} `json:"employee"`
	Total        int    `json:"total"`
	PaymentState string `json:"paymentState"`
	OrderType    struct {
		Id string `json:"id"`
	} `json:"orderType"`
	TaxRemoved        bool   `json:"taxRemoved"`
	IsVat             bool   `json:"isVat"`
	State             string `json:"state"`
	ManualTransaction bool   `json:"manualTransaction"`
	GroupLineItems    bool   `json:"groupLineItems"`
	TestMode          bool   `json:"testMode"`
	CreatedTime       int64  `json:"createdTime"`
	ClientCreatedTime int64  `json:"clientCreatedTime"`
	ModifiedTime      int64  `json:"modifiedTime"`
}

type GetOrderLineItemsLineItemStruct struct {
	Id       string `json:"id"`
	OrderRef struct {
		Id string `json:"id"`
	} `json:"orderRef"`
	Item struct {
		Id string `json:"id"`
	} `json:"item"`
	Name                   string `json:"name"`
	AlternateName          string `json:"alternateName"`
	Price                  int    `json:"price"`
	ItemCode               string `json:"itemCode"`
	Note                   string `json:"note"`
	Printed                bool   `json:"printed"`
	CreatedTime            int64  `json:"createdTime"`
	OrderClientCreatedTime int64  `json:"orderClientCreatedTime"`
	Exchanged              bool   `json:"exchanged"`
	Refunded               bool   `json:"refunded"`
	IsRevenue              bool   `json:"isRevenue"`
	IsOrderFee             bool   `json:"isOrderFee"`
}

type GetOrderLineItemsResp struct {
	Elements []GetOrderLineItemsLineItemStruct `json:"elements"`
}

type ChargeCardResp struct {
	Id                   string `json:"id"`
	Amount               int    `json:"amount"`
	PaymentMethodDetails string `json:"payment_method_details"`
	AmountRefunded       int    `json:"amount_refunded"`
	Currency             string `json:"currency"`
	Created              int64  `json:"created"`
	Captured             bool   `json:"captured"`
	RefNum               string `json:"ref_num"`
	AuthCode             string `json:"auth_code"`
	Outcome              struct {
		NetworkStatus string `json:"network_status"`
		Type          string `json:"type"`
	} `json:"outcome"`
	Paid   bool   `json:"paid"`
	Status string `json:"status"`
	Source struct {
		Id                string `json:"id"`
		AddressLine1Check string `json:"address_line1_check"`
		AddressZip        string `json:"address_zip"`
		AddressZipCheck   string `json:"address_zip_check"`
		Brand             string `json:"brand"`
		ExpMonth          string `json:"exp_month"`
		ExpYear           string `json:"exp_year"`
		First6            string `json:"first6"`
		Last4             string `json:"last4"`
	} `json:"source"`
}

type ChargeCardReq struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Source   string `json:"source"`
}

type PayOrderReq struct {
	Ecomind string `json:"ecomind"`
	Tender  struct {
		LabelKey string `json:"label_key"`
		Label    string `json:"label"`
		Id       string `json:"id"`
	} `json:"tender"`
	Customer                  string   `json:"customer"`
	Email                     string   `json:"email"`
	Amount                    int      `json:"amount"`
	Currency                  string   `json:"currency" default:"USD"`
	Expand                    []string `json:"expand"`
	ExternalReferenceId       string   `json:"external_reference_id"`
	ExternalCustomerReference string   `json:"external_customer_reference"`
	Source                    string   `json:"source"`
	TipAmount                 int      `json:"tip_amount"`
}

type PayOrderResp struct {
	Id         string `json:"id"`
	Object     string `json:"object"`
	Amount     int    `json:"amount"`
	AmountPaid int    `json:"amount_paid"`
	Currency   string `json:"currency"`
	Charge     string `json:"charge"`
	Created    int64  `json:"created"`
	RefNum     string `json:"ref_num"`
	AuthCode   string `json:"auth_code"`
	Items      []struct {
		Amount      int    `json:"amount"`
		Description string `json:"description"`
	} `json:"items"`
	Source struct {
		AddressLine1Check string `json:"address_line1_check"`
		AddressZip        string `json:"address_zip"`
		AddressZipCheck   string `json:"address_zip_check"`
		Brand             string `json:"brand"`
		ExpMonth          string `json:"exp_month"`
		ExpYear           string `json:"exp_year"`
		First6            string `json:"first6"`
		Last4             string `json:"last4"`
	} `json:"source"`
	Status            string `json:"status"`
	StatusTransitions struct {
		Paid int64 `json:"paid"`
	} `json:"status_transitions"`
}
