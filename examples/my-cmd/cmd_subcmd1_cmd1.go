package main

import (
	"fmt"

	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func ex_make_cmd_subcmd1_cmd1() *commander.Command {
	cmd := &commander.Command{
		Run:       ex_run_cmd_subcmd1_cmd1,
		UsageLine: "cmd1 [options]",
		Short:     "runs cmd1 and exits",
		Long: `
runs cmd1 and exits.

ex:
 $ my-cmd subcmd1 cmd1
`,
		Flag: *flag.NewFlagSet("my-cmd-subcmd1-cmd1", flag.ExitOnError),
	}
	cmd.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
	return cmd
}

func ex_run_cmd_subcmd1_cmd1(cmd *commander.Command, args []string) {
	name := "my-cmd-subcmd1-" + cmd.Name()
	quiet := cmd.Flag.Lookup("q").Value.Get().(bool)
	fmt.Printf("%s: hello from subcmd1-cmd1 (quiet=%v)\n", name, quiet)
}

// EOF
