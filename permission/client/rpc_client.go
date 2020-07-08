package client

import (
	"fmt"

	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/permission/proto/permission"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// DefaultRPCPort 定义默认的RPC端口
	DefaultRPCPort = 50051
)

var (
	rpcClient *RPCClient
)

// RPCConfig RPC配置参数
type RPCConfig struct {
	Addr   string
	AppID  string
	AppKey string
}

// InitRPCClient 初始化全局的RPC客户端
func InitRPCClient(cfg *RPCConfig) {
	cli, err := NewRPCClient(cfg)
	if err != nil {
		panic(err)
	}
	rpcClient = cli
}

// NewRPCClient 创建RPC客户端
func NewRPCClient(cfg *RPCConfig) (*RPCClient, error) {
	if cfg == nil {
		cfg = new(RPCConfig)
	}

	if cfg.Addr == "" {
		cfg.Addr = fmt.Sprintf(":%d", DefaultRPCPort)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(new(tokenCredential).Init(cfg.AppID, cfg.AppKey)))

	conn, err := grpc.Dial(cfg.Addr, opts...)
	if err != nil {
		return nil, err
	}

	return &RPCClient{
		PermissionClient: permission.NewPermissionClient(conn),
	}, nil
}

// RPCClient RPC客户端
type RPCClient struct {
	PermissionClient permission.PermissionClient
}

// PermissionClient 权限客户端
func PermissionClient() permission.PermissionClient {
	if rpcClient == nil {
		panic("未初始化RPC客户端")
	}
	return rpcClient.PermissionClient
}

// tokenCredential 令牌认证
type tokenCredential struct {
	appID  string
	appKey string
}

func (t *tokenCredential) Init(appID, appKey string) *tokenCredential {
	t.appID = appID
	t.appKey = appKey
	return t
}

func (t *tokenCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"app_id":  t.appID,
		"app_key": t.appKey,
	}, nil
}

func (t *tokenCredential) RequireTransportSecurity() bool {
	return false
}
