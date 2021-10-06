package main

import (
	"log"
	. "shawnH/blockverifier"
	"shawnH/blockverifier/request"

	"github.com/urfave/cli"
)

func verifyCommand() cli.Command {
	command := cli.Command{
		Name:  "verify",
		Usage: "verify a btc block by blockhash",
		Flags: []cli.Flag{
			BlockHashFlag,
		},
		Action: verify,
	}
	return command
}

func verify(ctx *cli.Context) {
	blockHash := ctx.String(BlockHashFlag.Name)
	block := request.RequireBlockInfo(blockHash)
	if err := VerifyBlockHash(block); err != nil {
		log.Fatalf(err.Error())
		return
	}

}

func main() {
	app := cli.NewApp()
	app.Usage = "Verify a block in BTC, use verify command and -w with the blockhash to verify."

	app.Commands = []cli.Command{}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "Language for the greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from `FILE`",
		},
	}

}
