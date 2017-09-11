package job

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/antlinker/sdk/asapi"
	"github.com/antlinker/sdk/utils"
)

const (
	jobRouter = "/job/exec"
)

// Config 配置参数
type Config struct {
	HTTPAddr string
}

// NewHandle 创建作业任务
func NewHandle(auh *asapi.AuthorizeHandle, config *Config) *Handle {
	return &Handle{
		auh: auh,
		cfg: config,
	}
}

// Handle 作业任务
type Handle struct {
	auh *asapi.AuthorizeHandle
	cfg *Config
}

func (h *Handle) getURL(router string) string {
	addr := h.cfg.HTTPAddr
	if len(addr) == 0 {
		return ""
	}

	if addr[len(addr)-1] == '/' {
		addr = addr[:len(addr)-1]
	}

	return addr + router
}

func (h *Handle) getANTUserID(intelUserCode string) (userID string, err error) {
	auids, ar := h.auh.GetAntUIDList("", intelUserCode)
	if ar != nil {
		err = ar
		return
	} else if len(auids) == 0 {
		err = errors.New("not found user")
		return
	}
	userID = auids[0]
	return
}

func (h *Handle) request(body interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	data, err := utils.PostJSON(ctx, h.getURL(jobRouter), body, func(req *http.Request) (*http.Request, error) {
		token, err := asapi.GetToken()
		if err != nil {
			return nil, err
		}
		req.Header.Set("AccessToken", token)
		return req, nil
	})
	if err != nil {
		if len(data) > 0 {
			log.Println("请求发生错误：", string(data))
		}
		return
	}

	var str string
	json.Unmarshal(data, &str)
	if str != "ok" {
		err = fmt.Errorf(string(data))
	}

	return
}

// ModifyStaffName 学工更改用户姓名
func (h *Handle) ModifyStaffName(intelUserCode, name string) (err error) {
	userID, err := h.getANTUserID(intelUserCode)
	if err != nil {
		return
	}

	body := map[string]string{
		"UID":  userID,
		"Name": name,
	}

	err = h.request(body)

	return
}

// ModifyStaffDept 学工更改学工用户所在的部门
func (h *Handle) ModifyStaffDept(intelUserCode, deptID, deptName string) (err error) {
	userID, err := h.getANTUserID(intelUserCode)
	if err != nil {
		return
	}

	body := map[string]string{
		"UID":      userID,
		"DeptID":   deptID,
		"DeptName": deptName,
	}

	err = h.request(body)

	return
}

// ModifyStudentClassRequest 学工更改学生用户所在的班级请求参数
type ModifyStudentClassRequest struct {
	AcademyID   string
	AcademyName string
	MajorID     string
	MajorName   string
	GradeName   string
	ClassID     string
	ClassName   string
}

// ModifyStudentClass 学工更改学生用户所在的班级
func (h *Handle) ModifyStudentClass(intelUserCode string, req *ModifyStudentClassRequest) (err error) {
	userID, err := h.getANTUserID(intelUserCode)
	if err != nil {
		return
	}

	body := map[string]string{
		"UID":         userID,
		"AcademyID":   req.AcademyID,
		"AcademyName": req.AcademyName,
		"MajorID":     req.MajorID,
		"MajorName":   req.MajorName,
		"GradeName":   req.GradeName,
		"ClassID":     req.ClassID,
		"ClassName":   req.ClassName,
	}

	err = h.request(body)

	return
}

// ModifyStaffClass 更改辅导员用户所管理的班级
func (h *Handle) ModifyStaffClass(intelUserCode string) (err error) {
	userID, err := h.getANTUserID(intelUserCode)
	if err != nil {
		return
	}

	body := map[string]string{
		"UID": userID,
	}

	err = h.request(body)

	return
}

// ModifyStaffDeptName 更改学工用户所在的部门名称
func (h *Handle) ModifyStaffDeptName(university, deptID, deptName string) (err error) {
	body := map[string]string{
		"University": university,
		"DeptID":     deptID,
		"DeptName":   deptName,
	}

	err = h.request(body)

	return
}

// ModifyStaffClassName 更改学工用户所在的班级名称
func (h *Handle) ModifyStaffClassName(university, deptID, deptName string) (err error) {
	body := map[string]string{
		"University": university,
		"DeptID":     deptID,
		"DeptName":   deptName,
	}

	err = h.request(body)

	return
}
