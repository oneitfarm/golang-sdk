package token

import (
	"testing"
)

func Test_GetBearerFromAuth(t *testing.T) {

	a, b := GetBearerFromAuth("Bearer mytoken")
	if !b || a != "mytoken" {
		t.FailNow()
	}

}
