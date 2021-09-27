package main

import (
	"github.com/urfave/cli"
)

func verifyCommand() cli.Command {
	command := cli.Command{
		Name:  "verify",
		Usage: "verify a btc block",

		Action: verify,
	}
	return command
}

func verify(ctx *cli.Context) error {

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
