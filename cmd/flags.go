package main

import (
	"github.com/urfave/cli"
)

var (
	BlockHashFlag = &cli.StringFlag{
		Name:  "blockhash, bh",
		Usage: "the block hash to get a block",
	}
)
