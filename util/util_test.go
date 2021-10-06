package util

import (
	"log"
	"testing"
)

func TestString2ByteArray(t *testing.T) {
	bytes := String2ByteArray("5b09bbb8d3cb2f8d4edbcf30664419fb7c9deaeeb1f62cb432e7741c80dbe5ba")
	log.Printf("%x\n", bytes)

}
