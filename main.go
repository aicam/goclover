package goclover

func InitApp(clientSecret string, args ...string) *Service {
	newService := Service{ClientSecret: clientSecret}
	if len(args) > 0 {
		if args[0] == "sandbox" {
			newService.BaseUrl = "https://sandbox.dev.clover.com/"
		}
	} else {
		newService.BaseUrl = "https://sandbox.dev.clover.com/"
	}
	return &newService
}

func InitSession(app Service, accessSecret, mid string) *Session {
	return &Session{
		Service:      app,
		AccessSecret: accessSecret,
		MID:          mid,
	}
}
