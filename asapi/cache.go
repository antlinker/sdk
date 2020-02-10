package asapi

import (
	"crypto/md5"
	"fmt"
)

var routerCached = map[string]int64{
	"/api/authorize/getstaffparam":      60,
	"/api/authorize/antuidbyuniversity": 60,
}

// RequestReader 请求
type RequestReader interface {
	// Hash 返回请求的哈希值
	Hash() string
	// Expires 返回请求缓存的过期时间
	// 如果返回的值小于0，则不会使用缓存
	Expires(router string) int64
}

// GetStaffParamRequest 获取学工参数的请求
type GetStaffParamRequest struct {
	ServiceIdentify string `json:"ServiceIdentify"`
	UID             string `json:"UID"`
}

// Hash 根据请求中的ServiceIdentify和UID拼接后返回md5值
func (r *GetStaffParamRequest) Hash() string {
	s := fmt.Sprintf("getstaffparam:%s:%s", r.ServiceIdentify, r.UID)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Expires 查询学工参数的请求，其响应结果需要缓存的时间
func (r *GetStaffParamRequest) Expires(router string) int64 {
	return routerCached[router]
}

// GetAntUIDByUniversityRequest 获取学工账号绑定的集结号UID的请求
type GetAntUIDByUniversityRequest struct {
	ServiceIdentify string `json:"ServiceIdentify"`
	UserID          string `json:"UserID"`
	University      string `json:"University"`
}

// Hash 获取学工绑定集结号账号的UID请求参数的哈希值
func (r *GetAntUIDByUniversityRequest) Hash() string {
	s := fmt.Sprintf("antuidbyuniversity:%s:%s:%s",
		r.ServiceIdentify, r.UserID, r.University)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Expires 返回获取学工账号绑定的集结号UID的请求响应结果缓存时间
func (r *GetAntUIDByUniversityRequest) Expires(router string) int64 {
	return routerCached[router]
}

// SetRouterExpires 设置启用了缓存的接口的过期时间
func SetRouterExpires(m map[string]int64) {
	for k, v := range m {
		if _, ok := routerCached[k]; ok {
			routerCached[k] = v
		}
	}
}
