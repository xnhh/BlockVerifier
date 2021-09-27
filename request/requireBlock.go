package request

import (
	"fmt"
)

const url = "https://blockchain.info/rawblock/"

func requireBlockInfo(blockhash string) *Block {
	url := fmt.Sprintf("%s%s", url, blockhash)
	fmt.Printf(url)
	return nil

}
