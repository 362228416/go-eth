package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	iv = "16-Bytes--String"
)

func GetAesKey() []byte {
	homeDir, _ := os.UserHomeDir()
	filename := homeDir + "/.ssh/aes_key"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read key fails", err)
	}
	return []byte(strings.Trim(string(bytes), "\n"))
}

func AesEncryptCFB(origData []byte, key []byte) string {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                         // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)           // 补全码
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv)) // 加密模式
	encrypted := make([]byte, len(origData))               // 创建数组
	blockMode.CryptBlocks(encrypted, origData)             // 加密
	return byteToHex(encrypted)
}

func AesDecryptCFB(encrypted string, key []byte) string {
	byteEncrypted := hexToByte(encrypted)
	block, _ := aes.NewCipher(key)                         // 分组秘钥
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv)) // 加密模式
	decrypted := make([]byte, len(byteEncrypted))          // 创建数组
	blockMode.CryptBlocks(decrypted, byteEncrypted)        // 解密
	decrypted = PKCS5Trimming(decrypted)
	return string(decrypted)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func byteToHex(inStr []byte) string {
	out := strings.Builder{}
	for i := 0; i < len(inStr); i++ {
		if int(inStr[i])&0xff < 0x10 {
			out.WriteString("0")
		}
		out.WriteString(fmt.Sprintf("%x", int(inStr[i])&0xff))
	}
	return out.String()
}

func hexToByte(content string) []byte {
	out := bytes.Buffer{}
	//out := make([]byte, len(content) / 2)
	for i := 0; i < len(content)/2; i++ {
		//binary[i] = (byte) Integer.parseInt(hex.substring(2 * i, 2 * i + 2), 16);
		//content[2*i:2*i+2]
		v, _ := strconv.ParseInt(content[2*i:2*i+2], 16, 32)
		b, _ := IntToBytes(v, 1)
		out.Write(b)
	}
	return out.Bytes()

}

//整形转换成字节
func IntToBytes(n int64, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

func AesDecrypt(content string) string {
	return AesDecryptCFB(content, GetAesKey())
}
