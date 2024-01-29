package goclover

func InitApp(clientSecret string, args ...string) *Service {
	newService := Service{ClientSecret: clientSecret}
	if len(args) > 0 {
		if args[0] == "sandbox" {
			newService.InventoryBaseUrl = "https://sandbox.dev.clover.com/"
			newService.EcommerceBaseUrl = "https://scl-sandbox.dev.clover.com/"
		}
	} else {
		newService.InventoryBaseUrl = "https://sandbox.dev.clover.com/"
		newService.EcommerceBaseUrl = "https://scl-sandbox.dev.clover.com/"
	}
	return &newService
}

func InitSession(app Service, accessSecret, mid string) *InventorySession {
	return &InventorySession{
		Service:      app,
		AccessSecret: accessSecret,
		MID:          mid,
	}
}
