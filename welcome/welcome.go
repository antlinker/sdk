package welcome

import (
	"github.com/antlinker/go-mqtt/client"
	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/asapi"
	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/utils"
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

// SetMQTTClient 设置MQTT客户端
func SetMQTTClient(mqcfg *utils.MQTTConfig) {
	if gHandle == nil {
		gHandle = new(Handle)
	}

	cfg := &utils.MQTTConfig{
		BrokerAddress:     mqcfg.BrokerAddress,
		ClientID:          mqcfg.ClientID,
		Timeout:           mqcfg.Timeout,
		CleanSession:      mqcfg.CleanSession,
		ReconnectInterval: mqcfg.ReconnectInterval,
		Username:          mqcfg.Username,
		PasswordHandler:   mqcfg.PasswordHandler,
		EnableTLS:         mqcfg.EnableTLS,
	}

	cli, err := utils.NewMQTTClient(cfg)
	if err != nil {
		panic(err)
	}

	gHandle.mqcli = cli
}

// SetMQTTClienter 设置MQTT客户端
func SetMQTTClienter(cli client.MqttClienter) {
	gHandle.mqcli = cli
}

// Click 任务点击记录
func Click(req *ClickRequest) (err error) {
	err = gHandle.Click(req)
	return
}

// Done 任务完成记录
func Done(req *DoneRequest) (err error) {
	err = gHandle.Done(req)
	return
}

// Cancel 任务取消记录
func Cancel(req *CancelRequest) (err error) {
	err = gHandle.Cancel(req)
	return
}
