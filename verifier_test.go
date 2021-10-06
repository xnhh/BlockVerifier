package blockverifier

import (
	"log"
	"shawnH/blockverifier/request"
	. "shawnH/blockverifier/util"
	"testing"
)

func TestVerifyBlockHash(t *testing.T) {
	block := request.RequireBlockInfo("0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103")
	VerifyBlockHash(block)
}

func TestDoubleSha256(t *testing.T) {
	node1 := "f5dd84f36b869e062a473b40a2ce283af6d40b235a461f6ca239535e236ce0b2"
	node2 := "f5dd84f36b869e062a473b40a2ce283af6d40b235a461f6ca239535e236ce0b2"
	newHash := append(reverse(String2ByteArray(node1)), reverse(String2ByteArray(node2))...)
	log.Printf("%x\n", newHash)
	log.Printf("%x", reverse(doubleSha256(newHash))) //shoule be 827f header
}
