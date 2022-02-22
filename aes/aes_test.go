package aes

import (
	"fmt"
	"testing"
)

func TestAesEncryptCFB(t *testing.T) {
	content := "d47008ef74a998e20280fd59c7b68107"
	originKey := "hello"
	d := AesEncryptCFB([]byte(originKey), GetAesKey())
	fmt.Println(d)
	if d != content {
		t.Fatal("加密失败")
	}
}

func TestDencryptCFB(t *testing.T) {
	content := "d47008ef74a998e20280fd59c7b68107"
	originKey := "hello"
	d := AesDecryptCFB(content, GetAesKey())
	fmt.Println(d)
	if d != originKey {
		t.Fatal("解密失败")
	}
}
