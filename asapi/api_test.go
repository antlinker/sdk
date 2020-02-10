package asapi

var gconfig *Config

func init() {
	gconfig = &Config{
		ASURL:           "",
		ClientID:        "",
		ClientSecret:    "",
		ServiceIdentify: "",
		IsEnabledCache:  true,
	}
}
