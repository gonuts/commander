package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func ex_make_cmd_subcmd1() *commander.Commander {
	cmd := &commander.Commander{
		Name:  "subcmd1",
		Short: "subcmd1 sub-commander. does subcmd1 thingies",
		Commands: []*commander.Command{
			ex_make_cmd_subcmd1_cmd1(),
			ex_make_cmd_subcmd1_cmd2(),
		},
		Flag: flag.NewFlagSet("my-cmd-subcmd1", flag.ExitOnError),
	}
	return cmd
}

// EOF
