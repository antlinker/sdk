package job

import (
	"github.com/antlinker/sdk/asapi"
)

var (
	gHandle *Handle
)

// SetAuthorizeHandle 设置授权处理
func SetAuthorizeHandle(auh *asapi.AuthorizeHandle) {
	if gHandle == nil {
		gHandle = new(Handle)
	}
	gHandle.auh = auh
}

// SetConfig 设置配置参数
func SetConfig(cfg *Config) {
	if gHandle == nil {
		gHandle = new(Handle)
	}
	gHandle.cfg = cfg
}

// ModifyStaffName 学工更改用户姓名
func ModifyStaffName(intelUserCode, name string) (err error) {
	err = gHandle.ModifyStaffName(intelUserCode, name)
	return
}

// ModifyStaffDept 学工更改学工用户所在的部门
func ModifyStaffDept(intelUserCode, deptID, deptName string) (err error) {
	err = gHandle.ModifyStaffDept(intelUserCode, deptID, deptName)
	return
}

// ModifyStudentClass 学工更改学生用户所在的班级
func ModifyStudentClass(intelUserCode string, req *ModifyStudentClassRequest) (err error) {
	err = gHandle.ModifyStudentClass(intelUserCode, req)
	return
}

// ModifyStaffClass 更改辅导员用户所管理的班级
func ModifyStaffClass(intelUserCode string) (err error) {
	err = gHandle.ModifyStaffClass(intelUserCode)
	return
}

// ModifyStaffDeptName 更改学工用户所在的部门名称
func ModifyStaffDeptName(university, deptID, deptName string) (err error) {
	err = gHandle.ModifyStaffDeptName(university, deptID, deptName)
	return
}

// ModifyStaffClassName 更改学工用户所在的班级名称
func ModifyStaffClassName(university, deptID, deptName string) (err error) {
	err = gHandle.ModifyStaffClassName(university, deptID, deptName)
	return
}

// ModifyStudentGraduate 更改已毕业学生离校状态
func ModifyStudentGraduate(university, usercode []string) (err error) {
	err = gHandle.ModifyStudentGraduate(university, usercode)
	return
}
