package request

import (
	"testing"
)

func TestHttpGetBlockInfo(t *testing.T) {
	url := "https://blockchain.info/rawblock/0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103"
	res, err := HttpGet(url)
	if err != nil {
		t.Logf("Error: %s", err.Error())
		t.Fail()
	} else {
		t.Log(res)
	}

}
