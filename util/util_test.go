package util

import (
	"log"
	"testing"
)

func TestString2ByteArray(t *testing.T) {
	bytes := String2ByteArray("5b09bbb8d3cb2f8d4edbcf30664419fb7c9deaeeb1f62cb432e7741c80dbe5ba")
	log.Printf("%x\n", bytes)

}

func TestDoubleSha256(t *testing.T) {
	node1 := "f5dd84f36b869e062a473b40a2ce283af6d40b235a461f6ca239535e236ce0b2"
	node2 := "f5dd84f36b869e062a473b40a2ce283af6d40b235a461f6ca239535e236ce0b2"
	newHash := append(Reverse(String2ByteArray(node1)), Reverse(String2ByteArray(node2))...)
	log.Printf("%x\n", newHash)
	log.Printf("%x", Reverse(DoubleSha256(newHash))) //shoule be 827f header
}
