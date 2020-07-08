package plan

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/asapi"
	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/utils"
)

const (
	jobRouter = "/job/plan"
)

// Config 配置参数
type Config struct {
	HTTPAddr string
}

// NewHandle 创建作业任务
func NewHandle(auh *asapi.AuthorizeHandle, config *Config) *Handle {
	return &Handle{
		auh: auh,
		cfg: config,
	}
}

// Handle 作业任务
type Handle struct {
	auh *asapi.AuthorizeHandle
	cfg *Config
}

func (h *Handle) getURL(router string) string {
	addr := h.cfg.HTTPAddr
	if len(addr) == 0 {
		return ""
	}

	if addr[len(addr)-1] == '/' {
		addr = addr[:len(addr)-1]
	}

	return addr + router
}

func (h *Handle) getANTUserID(intelUserCode string) (userID string, err error) {
	auids, ar := h.auh.GetAntUIDList("", intelUserCode)
	if ar != nil {
		err = ar
		return
	} else if len(auids) == 0 {
		err = errors.New("not found user")
		return
	}
	userID = auids[0]
	return
}

// Request 计划请求
type Request struct {
	Type      string    `json:"typ"`
	Data      string    `json:"data"`
	StartTime time.Time `json:"startTime"`
	Spec      string    `json:"spec"`
	Repeat    int       `json:"repeat"`
}

func (h *Handle) request(req Request) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	data, err := utils.PostJSON(ctx, h.getURL(jobRouter), req, func(req *http.Request) (*http.Request, error) {
		token, err := asapi.GetToken()
		if err != nil {
			return nil, err
		}
		req.Header.Set("AccessToken", token)
		return req, nil
	})
	if err != nil {
		if len(data) > 0 {
			log.Println("请求发生错误：", string(data))
		}
		return
	}

	var str string
	json.Unmarshal(data, &str)
	if str != "ok" {
		err = fmt.Errorf(string(data))
	}

	return
}

// Test 测试任务
func (h *Handle) Test(spec string, startTime time.Time, repeat int) (err error) {

	err = h.AddPlan(Request{
		Type:      "test",
		Spec:      spec,
		StartTime: startTime,
		Repeat:    repeat,
	})

	return
}

// AddPlan 添加任务
func (h *Handle) AddPlan(req Request) error {
	if req.StartTime.IsZero() {
		req.StartTime = time.Now()
	}
	return h.request(req)
}
