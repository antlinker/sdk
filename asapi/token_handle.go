package asapi

import (
	"net/http"
	"sync"
	"time"

	"io/ioutil"

	"github.com/astaxie/beego/httplib"
)

// Token 令牌信息
type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	CreateTime  time.Time `json:"-"`
}

// NewTokenHandle 创建令牌验证
func NewTokenHandle(cfg *Config) *TokenHandle {
	t := &TokenHandle{
		cfg: cfg,
	}
	if cfg.MaxConns == 0 {
		t.transport = &http.Transport{
			DisableKeepAlives: true,
		}
		return t
	}
	if cfg.MaxConns < 0 {
		cfg.MaxConns = 10
	}
	t.transport = &http.Transport{
		MaxConnsPerHost: cfg.MaxConns,
	}
	return t
}

// TokenHandle 令牌验证处理
type TokenHandle struct {
	cfg       *Config
	lock      sync.Mutex
	token     *Token
	transport http.RoundTripper
}

// ForceGet 强制获取最新的令牌数据
func (th *TokenHandle) ForceGet() (token *Token, result *ErrorResult) {
	req := httplib.Post(th.cfg.GetURL("/oauth2/token"))
	req = req.SetBasicAuth(th.cfg.ClientID, th.cfg.ClientSecret)
	req.SetTransport(th.transport)
	req = req.Param("grant_type", "client_credentials")
	res, err := req.Response()
	if err != nil {
		result = NewErrorResult(err.Error())
		return
	} else if res.StatusCode != 200 {
		defer res.Body.Close()
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			result = NewErrorResult(err.Error())
			return
		}
		result = NewErrorResult(string(buf), res.StatusCode)
		return
	}
	var t Token
	err = req.ToJSON(&t)
	if err != nil {
		result = NewErrorResult(err.Error())
		return
	}
	t.CreateTime = time.Now()
	token = &t
	return
}

// Get 获取令牌
func (th *TokenHandle) Get() (tokenString string, result *ErrorResult) {
	th.lock.Lock()
	defer th.lock.Unlock()
	if th.token == nil ||
		th.token.CreateTime.Add(time.Duration(th.token.ExpiresIn-10)*time.Second).Before(time.Now()) {
		token, vresult := th.ForceGet()
		if vresult != nil {
			result = vresult
			return
		}
		th.token = token
	}
	tokenString = th.token.AccessToken
	return
}
