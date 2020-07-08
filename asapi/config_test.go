package asapi

import (
	"reflect"
	"testing"
)

func TestParseErrorResult(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name       string
		args       args
		wantResult *ErrorResult
	}{
		{"true", args{msg: `{"code":0,"message":"无效令牌"}`}, &ErrorResult{Code: 0, Message: "无效令牌"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ParseErrorResult(tt.args.msg); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ParseErrorResult() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
