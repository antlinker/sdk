package asapi

var gconfig *Config

func init() {
	gconfig = &Config{
		ASURL:           "https://127.0.0.1:8090",
		ClientID:        "5e3e73941d41c83e2f0321b2",
		ClientSecret:    "bcd75fbea088359421ab38f790a7143e5e3243cd",
		ServiceIdentify: "ANT",
		IsEnabledCache:  true,
	}
}
