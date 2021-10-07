package blockverifier

import (
	. "shawnH/blockverifier/util"
)

type Node struct {
	hash                []byte
	left, right, parent *Node
}

func InitMerkleTree(hashes []string) *Node {
	length := len(hashes)
	if length == 0 {
		return nil
	}

	if length == 1 {
		return &Node{
			hash: String2ByteArray(hashes[0]),
		}
	}

	layer := make([]*Node, length)
	for i := range layer {
		layer[i] = &Node{
			hash: String2ByteArray(hashes[i]),
		}
		// log.Printf("the transaction with index %d has hash: %s\n", i, hashes[i])
	}
	// log.Printf("the transaction with index 0 has hash: %s\n", hashes[0])
	for len(layer) > 1 {
		// log.Printf("layer size: %d, and the first node hash: %x, the last node has hash: %x\n", len(layer), layer[0].hash, layer[len(layer)-1].hash)
		layer = getMerkleRoot(layer)
	}

	return layer[0]
}

func getMerkleRoot(nodes []*Node) []*Node {
	i := 0
	length := len(nodes)
	res := make([]*Node, (length+1)/2)
	for i < length {
		var new2Hash []byte
		newNode := &Node{
			left: nodes[i],
		}
		if i+1 == length {
			Reverse(nodes[i].hash)
			new2Hash = append(nodes[i].hash, nodes[i].hash...)
			Reverse(nodes[i].hash)
			newNode.right = nodes[i]
		} else {
			new2Hash = append(Reverse(nodes[i].hash), Reverse(nodes[i+1].hash)...)
			Reverse(nodes[i].hash)
			Reverse(nodes[i+1].hash)
			newNode.right = nodes[i+1]
			nodes[i+1].parent = newNode
		}
		newNode.hash = Reverse(DoubleSha256(new2Hash))
		nodes[i].parent = newNode

		res[i/2] = newNode
		i += 2
	}
	return res
}
