package todo

import (
	"encoding/json"
	"time"

	"github.com/antlinker/go-mqtt/client"
	"github.com/antlinker/sdk/asapi"
	"github.com/satori/go.uuid"
)

// NewHandle 创建待办事项处理
func NewHandle(auh *asapi.AuthorizeHandle, mqcli client.MqttClienter) *Handle {
	return &Handle{
		auh:   auh,
		mqcli: mqcli,
	}
}

// Handle 待办事项处理
type Handle struct {
	auh   *asapi.AuthorizeHandle
	mqcli client.MqttClienter
}

// AddRequest 增加待办事项请求参数
type AddRequest struct {
	UIDs         []string
	BuID         string
	TodoType     string
	ContentValue map[string]string
	URIValue     map[string]string
	PubTime      time.Time
	EndTime      time.Time
	Status       int
}

// Add 增加待办事项
func (h *Handle) Add(req *AddRequest) (err error) {
	auids, ar := h.auh.GetAntUIDList("", req.UIDs...)
	if ar != nil {
		err = ar
		return
	}

	mreq := map[string]interface{}{
		"MT":           "ADDTODO",
		"AID":          uuid.NewV4().String(),
		"UIDs":         auids,
		"TodoType":     req.TodoType,
		"ContentValue": req.ContentValue,
		"URIValue":     req.URIValue,
		"Status":       req.Status,
		"BuID":         req.BuID,
	}

	if !req.PubTime.IsZero() {
		mreq["PubTime"] = req.PubTime.Format("20060102150405")
	}

	if !req.EndTime.IsZero() {
		mreq["EndTime"] = req.EndTime.Format("20060102150405")
	}

	err = h.publish(mreq)
	return
}

// DoneRequest 完成待办事项请求参数
type DoneRequest struct {
	UID      string
	BuID     string
	TodoType string
}

// Done 完成待办事项
func (h *Handle) Done(req *DoneRequest) (err error) {
	auids, ar := h.auh.GetAntUIDList("", req.UID)
	if ar != nil {
		err = ar
		return
	} else if len(auids) == 0 {
		return
	}

	mreq := map[string]interface{}{
		"MT":       "COMPLETETODO",
		"UID":      auids[0],
		"BuID":     req.BuID,
		"TodoType": req.TodoType,
	}
	err = h.publish(mreq)
	return
}

// DelRequest 删除待办事项请求参数
type DelRequest struct {
	UID      string
	BuID     string
	TodoType string
}

// Del 删除待办事项
func (h *Handle) Del(req *DelRequest) (err error) {
	auids, ar := h.auh.GetAntUIDList("", req.UID)
	if ar != nil {
		err = ar
		return
	} else if len(auids) == 0 {
		return
	}

	mreq := map[string]interface{}{
		"MT":       "DELTODO",
		"UID":      auids[0],
		"BuID":     req.BuID,
		"TodoType": req.TodoType,
	}
	err = h.publish(mreq)
	return
}

func (h *Handle) publish(data interface{}) (err error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = h.mqcli.Publish("S/TODO", client.QoS1, false, buf)
	return
}
