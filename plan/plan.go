package plan

import (
	"time"

	"github.com/antlinker/sdk/asapi"
)

var (
	gHandle *Handle
)

// SetAuthorizeHandle 设置授权处理
func SetAuthorizeHandle(auh *asapi.AuthorizeHandle) {
	if gHandle == nil {
		gHandle = new(Handle)
	}
	gHandle.auh = auh
}

// SetConfig 设置配置参数
func SetConfig(cfg *Config) {
	if gHandle == nil {
		gHandle = new(Handle)
	}
	gHandle.cfg = cfg
}

// Test 测试任务
func Test(spec string, startTime time.Time, repeat int) (err error) {
	err = gHandle.Test(spec, startTime, repeat)
	return
}
