package main

import (
	"fmt"
	"os"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

var g_cmd *commander.Commander

func init() {
	g_cmd = &commander.Commander{
		Name: os.Args[0],
		Commands: []*commander.Command{
			ex_make_cmd_cmd1(),
			ex_make_cmd_cmd2(),
		},
		Flag: flag.NewFlagSet("my-cmd", flag.ExitOnError),
		Commanders: []*commander.Commander{
			ex_make_cmd_subcmd1(),
			ex_make_cmd_subcmd2(),
		},
	}
}

func main() {
	err := g_cmd.Flag.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("**err**: %v\n", err)
		os.Exit(1)
	}

	args := g_cmd.Flag.Args()
	err = g_cmd.Run(args)
	if err != nil {
		fmt.Printf("**err**: %v\n", err)
		os.Exit(1)
	}

	return
}

// EOF
