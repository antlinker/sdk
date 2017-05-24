package todo

import (
	"github.com/antlinker/go-mqtt/client"
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

// SetMQTTClient 设置MQTT客户端
func SetMQTTClient(mqcfg *MQTTConfig) {
	if gHandle == nil {
		gHandle = new(Handle)
	}

	cli, err := NewMQTTClient(mqcfg)
	if err != nil {
		panic(err)
	}

	gHandle.mqcli = cli
}

// SetMQTTClienter 设置MQTT客户端
func SetMQTTClienter(cli client.MqttClienter) {
	gHandle.mqcli = cli
}

// Add 增加待办事项
func Add(req *AddRequest) (err error) {
	err = gHandle.Add(req)
	return
}

// Done 完成待办事项
func Done(req *DoneRequest) (err error) {
	err = gHandle.Done(req)
	return
}

// Del 删除待办事项
func Del(req *DelRequest) (err error) {
	err = gHandle.Del(req)
	return
}
