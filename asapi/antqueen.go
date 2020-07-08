package asapi

import (
	"gogs.xiaoyuanjijiehao.com/aag/ant-queen/pkg/conf"
)

// AntQueenInit 使用antQueen框架时可以使用该方法初始化
// key 文件名 不指定默认使用antlinkerauth.toml
func AntQueenInit(key ...string) {
	file := "antlinkerauth.toml"
	if len(key) == 1 {
		file = key[0]
	}
	cfg := loadCfg(file)
	initAuth(&cfg)
}

func loadCfg(key string) (cfg authorizeConfig) {
	if err := conf.Get(key).UnmarshalTOML(&cfg); err != nil {
		if err != conf.ErrNotExist {
			panic(err)
		}
	}
	return
}

type authorizeConfig struct {
	Enable          bool   `toml:"enable,omitempty" json:"enable,omitempty"`
	URL             string `toml:"url,omitempty" json:"url,omitempty" yaml:"url"`
	ClientID        string `toml:"client_id,omitempty" json:"client_id,omitempty" yaml:"client_id"`
	ClientSecret    string `toml:"client_secret,omitempty" json:"client_secret,omitempty" yaml:"client_secret"`
	Identify        string `toml:"identify,omitempty" json:"identify,omitempty" yaml:"identify"`
	IsEnabledCache  bool   `toml:"is_enabled_cache,omitempty" json:"is_enabled_cache,omitempty"`   // 是否启用缓存
	CacheGCInterval int    `toml:"cache_gc_interval,omitempty" json:"cache_gc_interval,omitempty"` // 缓存gc间隔(单位秒)
	MaxConns        int    `toml:"max_conns,omitempty" json:"max_conns,omitempty"`
}

func initAuth(config *authorizeConfig) {
	cache := true
	if !config.IsEnabledCache {
		cache = false
	}
	interval := 60
	if config.CacheGCInterval > 0 {
		interval = config.CacheGCInterval
	}
	InitAPI(&Config{
		ASURL:           config.URL,
		ClientID:        config.ClientID,
		ClientSecret:    config.ClientSecret,
		ServiceIdentify: config.Identify,
		IsEnabledCache:  cache,
		CacheGCInterval: interval,
		MaxConns:        config.MaxConns,
	})
}
