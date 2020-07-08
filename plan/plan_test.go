package plan_test

import (
	"testing"
	"time"

	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/asapi"
	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/plan"
)

func TestTest(t *testing.T) {
	asapi.InitAPI(&asapi.Config{
		ASURL:           "http://192.168.1.202:8090",
		ClientID:        "57d0c56597d5b1413bd01262",
		ClientSecret:    "0d1414f8743294f14c53cb46bcb8d5ea3e773026",
		ServiceIdentify: "ANT",
		IsEnabledCache:  true,
		CacheGCInterval: 60,
	})
	plan.SetConfig(&plan.Config{
		HTTPAddr: "http://127.0.0.1:8901",
	})
	err := plan.Test("", time.Now().Add(10*time.Second), 1)
	if err != nil {
		t.Error(err)
	}

}
