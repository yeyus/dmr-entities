package utils

import (
	"testing"
)

func TestGetCountryCodeForId(t *testing.T) {
	var testCases = map[int]string{
		21441:   "es",
		3106:    "us",
		20811:   "fr",
		208647:  "fr",
		647647:  "re",
		1106167: "us",
		2147197: "es",
		7229999: "ar",
		9999:    "",
		0000:    "",
	}

	for k, v := range testCases {
		o := GetCountryCodeForID(k)
		if o != v {
			t.Errorf("Expected \"%s\" for id %d and found \"%v\"", v, k, o)
		}
	}
}
