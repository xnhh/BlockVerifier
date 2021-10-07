package blockverifier

import (
	"shawnH/blockverifier/request"
	"testing"
)

func TestVerifyBlockHash(t *testing.T) {
	block := request.RequireBlockInfo("0000000000000bae09a7a393a8acded75aa67e46cb81f7acaa5ad94f9eacd103")
	VerifyBlockHash(block)
}
