package util

import "testing"

func TestGenerateToken(t *testing.T) {
	token , err := GenerateToken("vj", 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	info , err := ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
