package main

import (
	"github.com/urfave/cli"
)

var (
	BlockHashFlag = &cli.StringFlag{
		Name:  "blockhash, bh",
		Value: "",
		Usage: "verify the block from `blockhash`",
	}
)
