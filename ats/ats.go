package ats

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gogs.xiaoyuanjijiehao.com/antlinker/sdk/utils"
)

var (
	config *Config
)

// Config 配置参数
type Config struct {
	HTTPAddr string
}

// SetConfig 初始化
func SetConfig(cfg *Config) {
	config = cfg
}

// ConvertHTMLToPDF html转换为pdf
func ConvertHTMLToPDF(src []byte) ([]byte, error) {
	if config == nil {
		return nil, nil
	}

	addr := config.HTTPAddr
	if len(addr) == 0 {
		return nil, nil
	} else if addr[len(addr)-1] == '/' {
		addr = addr[:len(addr)-1]
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/pdf", addr), bytes.NewBuffer(src))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	var pdfData []byte
	err = utils.Request(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		zr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(zr)
		if err != nil {
			return err
		}
		pdfData = buf

		zr.Close()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return pdfData, nil
}
