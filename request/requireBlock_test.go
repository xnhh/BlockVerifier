package request

import (
	"fmt"
	"testing"
)

func TestRequireBlock(t *testing.T) {
	blockhash := "0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103"
	block := requireBlockInfo(blockhash)
	fmt.Println(block.Hash)
	fmt.Println(block.MerkleRoot)
}
