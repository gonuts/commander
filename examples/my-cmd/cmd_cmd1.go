package main

import (
	"flag"
	"fmt"

	"github.com/gonuts/commander"
)

var cmd_cmd1 = &commander.Command{
	Run:       ex_run_cmd_cmd1,
	UsageLine: "cmd1 [options]",
	Short:     "runs cmd1 and exits",
	Long: `
runs cmd1 and exits.

ex:
$ my-cmd cmd1
`,
	Flag: *flag.NewFlagSet("my-cmd-cmd1", flag.ExitOnError),
}

func init() {
	cmd_cmd1.Flag.Bool("q", true, "only print error and warning messages, all other output will be suppressed")
}

func ex_run_cmd_cmd1(cmd *commander.Command, args []string) error {
	name := "my-cmd-" + cmd.Name()
	quiet := cmd.Lookup("q").(bool)
	fmt.Printf("%s: hello from cmd1 (quiet=%v)\n", name, quiet)
	return nil
}

// EOF
