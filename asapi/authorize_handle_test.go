package asapi

import (
	"testing"
)

var _ah *AuthorizeHandle

func TestMain(m *testing.M) {
	_ah = NewAuthorizeHandle(gconfig)
	m.Run()
}

func TestGetUser(t *testing.T) {
	info, ar := _ah.GetUser("AA0000125923")
	t.Log(info, ar)
}

func TestGetAntUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		uid, err := _ah.GetAntUIDByUniversity("fa6d77be-11a6-4c29-bdb1-a86efa450f29", "11906")
		if err != nil {
			t.Errorf("GetAntUIDByUniversity error: %s", err)
			return
		}
		t.Logf("GetAntUIDByUniversity uid: %s", uid)
	}
}

func TestGetStaffParam(t *testing.T) {
	for i := 0; i < 10; i++ {
		info, res := _ah.GetAntStaffParam("AA0000125923")
		if res != nil {
			t.Errorf("GetAntStaffParam error: %s", res)
			return
		}
		t.Logf("GetAntStaffParam info: %s", info)
	}
}
