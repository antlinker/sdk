package asapi

import (
	"crypto/md5"
	"fmt"
	"sync"
)

// 部分接口缓存时间，默认60秒，可以使用SetRouterExpires函数重置特定的接口缓存时间
var routerCached = map[string]int64{
	"/api/authorize/getstaffparam":      60,
	"/api/authorize/antuidbyuniversity": 60,
	"/api/authorize/usercode":           60,
}

var routerLock sync.Mutex

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
	s := fmt.Sprintf("/api/authorize/getstaffparam:%s:%s", r.ServiceIdentify, r.UID)
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
	s := fmt.Sprintf("/api/authorize/antuidbyuniversity:%s:%s:%s",
		r.ServiceIdentify, r.UserID, r.University)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Expires 返回获取学工账号绑定的集结号UID的请求响应结果缓存时间
func (r *GetAntUIDByUniversityRequest) Expires(router string) int64 {
	return routerCached[router]
}

// GetUserCodeRequest 查询用户学号的接口
type GetUserCodeRequest struct {
	UID string `json:"UID"`
}

// Hash 返回查询用户学(工)号时请求的哈希值
func (r *GetUserCodeRequest) Hash() string {
	s := fmt.Sprintf("/api/authorize/usercode:%s", r.UID)
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Expires 返回获取学工号的请求响应结果缓存时间
func (r *GetUserCodeRequest) Expires(router string) int64 {
	return routerCached[router]
}

// SetRouterExpires 设置启用了缓存的接口的过期时间
func SetRouterExpires(m map[string]int64) {
	routerLock.Lock()
	for k, v := range m {
		if _, ok := routerCached[k]; ok {
			routerCached[k] = v
		}
	}
	routerLock.Unlock()
}
