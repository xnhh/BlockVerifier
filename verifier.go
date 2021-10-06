package blockverifier

import (
	"bytes"
	"fmt"
	"log"
	. "shawnH/blockverifier/model"
	. "shawnH/blockverifier/util"
)

func VerifyBlockHash(block *Block) error {
	hashes := make([]string, len(block.Transactions))
	for i := range hashes {
		hashes[i] = block.Transactions[i].Hash
		// log.Printf("the transaction with index %d has hash: %s\n", i, block.Transactions[i].Hash)
	}
	root := InitMerkleTree(hashes)
	if bytes.Equal(root.hash, String2ByteArray(block.MerkleRoot)) {
		log.Printf("Success, the merkle root hash is true, the tree is _\n")
		return nil
	} else {
		// log.Fatalf("Error: the merkle root hash is %x, but the correct one is %s\n", root.hash, block.MerkleRoot)
		return fmt.Errorf("Error: the merkle root hash is %x, but the correct one is %s\n", root.hash, block.MerkleRoot)
	}
}
