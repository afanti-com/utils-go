package vinCode

import "testing"

func TestVinCodeVerification(t *testing.T) {
	tests := []struct {
		id   string
		want bool
	}{
		{id: "LBEXDAFB77X53448A", want: false},
		{id: "LBEXDAFB77X534483", want: true},
		{id: "LFCBE22D970099074", want: false},
		{id: "LFMBE22D970099074", want: true},
		{id: "LVSFDFACX7F028380", want: false},
		{id: "LVSFDFACX7F028389", want: true},
		{id: "LSYBCAAB67K021683", want: false},
	}

	for _, test := range tests {
		if got := Verification(test.id); got != test.want {
			t.Errorf("VinCodeVerification(%s) got %v, want %v", test.id, got, test.want)
		}
	}
}
