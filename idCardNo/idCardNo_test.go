package idCardNo

import "testing"

func TestVerification(t *testing.T) {
	testDatas := []struct {
		ID   string
		Want bool
	}{
		{ID: "123123123", Want: false},
		{ID: "sad", Want: false},
	}

	for _, v := range testDatas {
		got := Verification(v.ID)
		if v.Want != got {
			t.Fatalf("ID[%s] want %v, but got %v", v.ID, v.Want, got)
		}
	}

}
