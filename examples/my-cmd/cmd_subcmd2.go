package main

import (
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func ex_make_cmd_subcmd2() *commander.Commander {
	cmd := &commander.Commander{
		Name:  "subcmd2",
		Short: "subcmd2 sub-commander. does subcmd2 thingies",
		Commands: []*commander.Command{
			ex_make_cmd_subcmd2_cmd1(),
			ex_make_cmd_subcmd2_cmd2(),
		},
		Flag: flag.NewFlagSet("my-cmd-subcmd2", flag.ExitOnError),
	}
	return cmd
}

// EOF
