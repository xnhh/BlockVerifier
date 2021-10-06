package main

import (
	"log"

	"shawnH/blockverifier/request"

	. "shawnH/blockverifier/util"

	"github.com/cbergoon/merkletree"
)

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	x string
}

//CalculateHash hashes the values of a TestContent
func (t TestContent) CalculateHash() ([]byte, error) {
	return String2ByteArray(t.x), nil
}

//Equals tests for equality of two Contents
func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}

func main() {
	//Build list of Content to build tree
	block := request.RequireBlockInfo("0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103")
	var list []merkletree.Content
	for _, t := range block.Transactions {
		list = append(list, TestContent{x: t.Hash})
	}

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Printf("%x\n", mr)

	//Verify the entire tree (hashes for each node) is valid
	vt, err := t.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	vc, err := t.VerifyContent(list[0])
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Verify Content: ", vc)

	//String representation
	// log.Println(t)
}
