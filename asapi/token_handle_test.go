package asapi

import (
	"fmt"
	"sync"
	"testing"
)

func TestTokenHandleGet(t *testing.T) {
	th := NewTokenHandle(gconfig)
	token, result := th.Get()
	if result != nil {
		t.Error(result.Code, result.Message)
		return
	}
	t.Log("Access Token:", token)
}

func TestTokenHandle_ForceGet(t *testing.T) {
	gconfig = &Config{
		ASURL:           "http://192.168.175.6:8090",
		ClientID:        "57a999b57a03b59ebb9b11b0",
		ClientSecret:    "9389211575bfa749b3efdfc3bcd2114e3344e025",
		ServiceIdentify: "TEST",
		MaxConns:        0,
	}
	th := NewTokenHandle(gconfig)
	var wg sync.WaitGroup
	for j := 0; j < 100; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			for i := 0; i < 100; i++ {
				t.Log("ForceGet()")
				token, result := th.ForceGet()
				if result != nil {
					t.Log(result.Code, result.Message)
				} else {
					t.Logf("Access Token:%+v\n", token)
				}
				// time.Sleep(100 * time.Millisecond)
			}
		}(j)
	}
	wg.Wait()

}
