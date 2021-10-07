package util

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
)

func Utf82Hex(r byte) byte {
	if r > '9' {
		return byte(r - '0' - 39)
	} else {
		return byte(r - '0')
	}
}

func DoubleSha256(origin []byte) []byte {
	h := sha256.New()
	mid := sha256.Sum256(origin)
	h.Write(mid[:])
	res := h.Sum(nil)
	return res[:]
}

func Reverse(origin []byte) []byte {
	i, j := 0, len(origin)-1
	for i < j {
		temp := origin[i]
		origin[i] = origin[j]
		origin[j] = temp
		i++
		j--
	}
	return origin
}

func String2ByteArray(s string) []byte {
	bytes := make([]byte, 0)
	i := 0
	length := len(s)
	for i < length {
		b := Utf82Hex(s[i])<<4 + Utf82Hex(s[i+1])
		// log.Printf("%x\n", cha)
		bytes = append(bytes, b)
		i += 2
	}
	return bytes
}

func IntToHex(num int32) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.LittleEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
