package goclover

type Service struct {
	ClientSecret     string
	InventoryBaseUrl string
	EcommerceBaseUrl string
}

type InventorySession struct {
	Service
	AccessSecret string
	MID          string
}
