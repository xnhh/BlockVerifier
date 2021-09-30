package request

import (
	"encoding/json"
	"fmt"
	"log"
	. "shawnH/blockverifier"
)

const url = "https://blockchain.info/rawblock/"

func requireBlockInfo(blockhash string) *Block {
	url := fmt.Sprintf("%s%s", url, blockhash)
	// fmt.Printf(url)
	res, err := HttpGet(url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	block := Block{}
	jsonErr := json.Unmarshal([]byte(res), &block)
	if jsonErr != nil {
		log.Fatalf(jsonErr.Error())
	}

	return &block
}
