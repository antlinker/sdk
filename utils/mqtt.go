package utils

import (
	"crypto/tls"

	"github.com/antlinker/go-mqtt/client"
)

// MQTTConfig mqtt配置参数
type MQTTConfig struct {
	BrokerAddress     string
	ClientID          string
	Timeout           uint16
	CleanSession      bool
	ReconnectInterval int
	Username          string
	PasswordHandler   func() string
	EnableTLS         bool
}

// NewMQTTClient 创建默认的MQTT客户端
func NewMQTTClient(cfg *MQTTConfig) (cli client.MqttClienter, err error) {
	opts := client.MqttOption{
		Addr:               cfg.BrokerAddress,
		Clientid:           cfg.ClientID,
		CleanSession:       cfg.CleanSession,
		ReconnTimeInterval: cfg.ReconnectInterval,
		UserName:           cfg.Username,
		PasswordHandler:    cfg.PasswordHandler,
		KeepAlive:          cfg.Timeout,
	}

	if cfg.EnableTLS {
		opts.TLS = &tls.Config{InsecureSkipVerify: true}
	}

	cli, err = client.CreateClient(opts)
	if err != nil {
		return
	}

	l := &client.DefaultPrintListener{}
	cli.AddConnListener(l)
	cli.AddDisConnListener(l)

	err = cli.Connect()
	return
}
