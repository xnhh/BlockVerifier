package blockVerifier

type Block struct {
	Hash string `json:"hash"`

	Version int16 `json:"ver"`

	PreviousBlockHash string `json:"prev_block"`

	MerkleRoot string `json:"mrkl_root"`

	Ts int64 `json:"time"`

	Bits int64 `json:"bits"`

	NextBlockHashes []string `json:"next_block"`

	Fee int64 `json:"fee"`

	Nonce int64 `json:"nonce"`

	TransactionsNum int16 `json:"n_tx"`

	BlockSize int64 `json:"size"`

	BlockIndex int64 `json:"block_index"`

	InMainChain bool `json:"main_chain"`

	Height int64 `json:"height"`

	Weight int64 `json:"weight"`

	Transactions []struct {
		Hash string `json:"nonce"`

		Version int16 `json:"ver"`

		InNum int16 `json:"vin_sz"`

		OutNum int16 `json:"vout_sz"`

		Weight int64 `json:"weight"`

		Fee int64 `json:"fee"`

		RelayedBy string `json:"relayed_by"`

		LockTs int64 `json:"lock_time"`

		TxIndex int64 `json:"tx_index"`

		DoubleSpend bool `json:"double_spend"`

		Ts int64 `json:"time"`

		BlockIndex int64 `json:"block_index"`

		BlockHeight int64 `json:"block_height"`

		Inputs []struct {
			Sequence int64 `json:"sequence"`

			Witness string `json:"witness"`

			Script string `json:"script"`

			Index int64 `json:"index"`

			PreviousOut struct {
				Spent bool `json:"spent"`

				Script string `json:"script"`

				SpendingOutpoints []struct {
					TxIndex int64 `json:"tx_index"`

					N int64 `json:"n"`
				} `json:"spending_outpoints"`

				TxIndex int64 `json:"tx_index"`

				Value int64 `json:"value"`

				Address string `json:"addr"`

				N int64 `json:"n"`

				Type int64 `json:"type"`
			} `json:"prev_out"`
		} `json:"inputs"`

		Outputs []struct {
			Spent bool `json:"spent"`

			Script string `json:"script"`

			SpendingOutpoints []struct {
				TxIndex int64 `json:"tx_index"`

				N int64 `json:"n"`
			} `json:"spending_outpoints"`

			TxIndex int64 `json:"tx_index"`

			Value int64 `json:"value"`

			Address string `json:"addr"`

			N int64 `json:"n"`

			Type int64 `json:"type"`
		} `json:"out"`
	} `json:"tx"`
}
