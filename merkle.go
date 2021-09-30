package blockverifier

import (
	"crypto/sha256"
)

func doubleSha256(origin []byte) []byte {
	h := sha256.New()
	mid := sha256.Sum256(origin)
	h.Write(mid[:])
	res := sha256.Sum256(h.Sum(nil))
	return res[:]
}

type Node struct {
	hash                []byte
	left, right, parent *Node
}

func initMerkleTree(hashes []string) *Node {
	length := len(hashes)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return &Node{
			hash: []byte(hashes[0]),
		}
	}

	layer := make([]*Node, length)
	for i, _ := range layer {
		layer[i] = &Node{
			hash: []byte(hashes[i]),
		}
	}

	for len(layer) > 1 {
		layer = getMerkleRoot(layer)
	}

	return layer[0]
}

func getMerkleRoot(nodes []*Node) []*Node {
	i := 0
	length := len(nodes)
	res := make([]*Node, (length+1)/2)
	for i < length {
		var newHash []byte
		newNode := &Node{
			left: nodes[i],
		}
		if i+1 == length {
			newHash = append(nodes[i].hash, nodes[i].hash...)
			newNode.right = nodes[i]
		} else {
			newHash = append(nodes[i].hash, nodes[i+1].hash...)
			newNode.right = nodes[i+1]
		}
		newNode.hash = newHash
		nodes[i].parent = newNode
		nodes[i+1].parent = newNode

		res[i/2] = newNode
		i += 2
	}
	return res
}
