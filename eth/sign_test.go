package core

import "testing"

func TestValidSign(t *testing.T) {
	address := "0xb5686faa50e064917593ad753c0c26b7764fdd85"
	msg := "hello"
	sign := "0xbfb56cb5e4426bf1a54eeb09d0d78eab84948afb5f8d950f9e7ede3a911ce5f839a05a07e097f56a7235d6149aceca55809465cb00f36f9c63b480cad104a6c51b"
	ret := ValidSign(address, sign, msg)
	if !ret {
		t.Error("验签失败")
	}
}
