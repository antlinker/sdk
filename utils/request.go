package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OptionHandle 自定义处理请求
type OptionHandle func(*http.Request) (*http.Request, error)

// PostJSON 发起一个json格式的POST请求
func PostJSON(ctx context.Context, url string, body interface{}, options ...OptionHandle) (data []byte, err error) {
	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return
		}
	}

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return
	}

	if len(options) > 0 {
		req, err = options[0](req)
	}

	err = Request(ctx, req, func(res *http.Response, err error) error {
		if err != nil {
			return err
		}

		defer res.Body.Close()
		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		data = buf

		if v := res.StatusCode; v != 200 {
			err = fmt.Errorf("请求发生错误，状态码：%d", v)
		}

		return nil
	})

	return
}

// Request HTTP请求处理
func Request(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan error, 1)
	go func() { c <- f(client.Do(req)) }()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		<-c
		return ctx.Err()
	case err := <-c:
		return err
	}
}
