package blockverifier

import (
	"bytes"
	"fmt"
	"log"
	. "shawnH/blockverifier/model"
	. "shawnH/blockverifier/util"
)

func VerifyBlockHash(block *Block) error {
	if block == nil {
		return fmt.Errorf("Error: cannot verify block with nil\n")
	}

	// Verify merkle root
	hashes := make([]string, len(block.Transactions))
	for i := range hashes {
		hashes[i] = block.Transactions[i].Hash
		// log.Printf("the transaction with index %d has hash: %s\n", i, block.Transactions[i].Hash)
	}
	root := InitMerkleTree(hashes)
	fmt.Println()
	if bytes.Equal(root.hash, String2ByteArray(block.MerkleRoot)) {
		fmt.Printf("Step1: verify the merkle root hash, it is %s \n", block.MerkleRoot)
		fmt.Printf("Build the merkle root hash with transactions, the merkle tree is: (Use the first 2 bytes of hash to present) \n")
		PrintMerkleTree([]*Node{root}, 0)
	} else {
		// log.Fatalf("Error: the merkle root hash is %x, but the correct one is %s\n", root.hash, block.MerkleRoot)
		return fmt.Errorf("Error: the calculated merkle root hash is %x, but the correct one is %s\n", root.hash, block.MerkleRoot)
	}
	fmt.Printf("Verify the merkle root hash success. \n")
	fmt.Println()

	fmt.Printf("Step2: verify the block hash, it is %s \n", block.Hash)
	fmt.Printf("The block message:\n")

	//version
	v := IntToHex(int32(block.Version))
	log.Printf("Version: %d\n", block.Version)

	// previous block hash
	pbh := Reverse(String2ByteArray(block.PreviousBlockHash))
	log.Printf("Previous block hash: %s\n", block.PreviousBlockHash)

	// merkle root hash
	mrh := Reverse(String2ByteArray(block.MerkleRoot))
	log.Printf("Merkle root hash: %s\n", block.MerkleRoot)

	//time
	t := IntToHex(int32(block.Ts))
	log.Printf("Timestamp: %d\n", block.Ts)

	//bits
	b := IntToHex(int32(block.Bits))
	log.Printf("NBits: %d\n", block.Bits)

	//Nonce
	n := IntToHex(int32(block.Nonce))
	log.Printf("Nonce: %d\n", block.Nonce)

	// all := append(v, append(pbh, append(mrh, append(t, append(b, n...)...)...)...)...)
	all := bytes.Join(
		[][]byte{
			v,
			pbh,
			mrh,
			t,
			b,
			n},
		[]byte{},
	)

	res := Reverse(DoubleSha256(all))

	if bytes.Equal(res, String2ByteArray(block.Hash)) {
		fmt.Printf("Verify the block hash success, it is %x\n", res)
	} else {
		log.Fatalf("Error: the calculated block hash is %x, but the correct one is %s\n", res, block.Hash)
		return fmt.Errorf("Error: the calculated block hash is %x, but the correct one is %s\n", res, block.Hash)
	}
	//log.Printf("Step2: the block hash is true, it is %x\n", block.Hash)
	fmt.Println()

	return nil
}
