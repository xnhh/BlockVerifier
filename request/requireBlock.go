package request

import (
	"encoding/json"
	"fmt"
	"log"
	. "shawnH/blockverifier/model"
)

const url = "https://blockchain.info/rawblock/"

func RequireBlockInfo(blockhash string) *Block {
	url := fmt.Sprintf("%s%s", url, blockhash)
	// fmt.Printf(url)
	res, err := HttpGet(url)
	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}
	block := Block{}
	jsonErr := json.Unmarshal([]byte(res), &block)
	if jsonErr != nil {
		log.Fatalf(jsonErr.Error())
		return nil
	}

	// log.Printf("received block: %+v \n", block)

	return &block
}
