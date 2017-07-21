package welcome

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/antlinker/go-mqtt/client"
	"github.com/antlinker/sdk/asapi"
)

// NewHandle 创建迎新处理
func NewHandle(auh *asapi.AuthorizeHandle, mqcli client.MqttClienter) *Handle {
	return &Handle{
		auh:   auh,
		mqcli: mqcli,
	}
}

// Handle 迎新处理
type Handle struct {
	auh   *asapi.AuthorizeHandle
	mqcli client.MqttClienter
}

// ClickRequest 任务点击记录请求参数
type ClickRequest struct {
	TaskID string
	UID    string
}

// Click 任务点击记录
func (h *Handle) Click(req *ClickRequest) (err error) {
	auids, ar := h.auh.GetAntUIDList("", req.UID)
	if ar != nil {
		err = ar
		return
	} else if len(auids) == 0 {
		err = errors.New("not found user")
		return
	}

	mreq := map[string]interface{}{
		"MT":     "CLICKTASK",
		"UID":    auids[0],
		"TaskID": req.TaskID,
	}

	err = h.publish(mreq)
	return
}

// DoneRequest 任务完成记录请求参数
type DoneRequest struct {
	DoneTime time.Time
	TaskCode string
	UID      string
}

// Done 任务完成记录
func (h *Handle) Done(req *DoneRequest) (err error) {
	auids, ar := h.auh.GetAntUIDList("", req.UID)
	if ar != nil {
		err = ar
		return
	} else if len(auids) == 0 {
		err = errors.New("not found user")
		return
	}

	mreq := map[string]interface{}{
		"MT":       "DONETASK",
		"UID":      auids[0],
		"TaskCode": req.TaskCode,
	}

	if !req.DoneTime.IsZero() {
		mreq["DoneTime"] = req.DoneTime.Format("20060102150405")
	}

	err = h.publish(mreq)
	return
}

// CancelRequest 任务取消记录请求参数
type CancelRequest struct {
	TaskCode string
	UID      string
}

// Cancel 任务取消记录
func (h *Handle) Cancel(req *CancelRequest) (err error) {
	auids, ar := h.auh.GetAntUIDList("", req.UID)
	if ar != nil {
		err = ar
		return
	} else if len(auids) == 0 {
		err = errors.New("not found user")
		return
	}

	mreq := map[string]interface{}{
		"MT":       "CANCELTASK",
		"UID":      auids[0],
		"TaskCode": req.TaskCode,
	}

	err = h.publish(mreq)
	return
}

func (h *Handle) publish(data interface{}) (err error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return
	}

	_, err = h.mqcli.Publish("S/WELCOME", client.QoS1, false, buf)
	return
}
