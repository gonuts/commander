// Copyright 2012 The Go-Commander Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Based on the original work by The Go Authors:
// Copyright 2011 The Go Authors.  All rights reserved.

// commander helps creating command line programs whose arguments are flags,
// commands and subcommands.
package commander

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/gonuts/flag"
)

// A Commander holds the configuration for the command line tool.
type Commander struct {
	// Name is the command name, usually the executable's name.
	Name string
	// Short is a short description of the commander
	Short string
	// Commands is the list of commands supported by this commander program.
	Commands []*Command
	// Flag is a set of flags for the whole commander. It should not be
	// changed after Run() is called.
	Flag *flag.FlagSet
	// Parent is the parent commander of this commander
	Parent *Commander
	// Commanders is the list of sub-commanders supported by this commander program.
	Commanders []*Commander
}

// Run executes the commander using the provided arguments. The command
// matching the first argument is executed and it receives the remaining
// arguments.
func (c *Commander) Run(args []string) error {
	if c == nil {
		return fmt.Errorf("Called Run() on a nil Commander")
	}
	// setup hierarchy...
	for _, cmd := range c.Commanders {
		cmd.Parent = c
	}

	if c.Flag == nil {
		c.Flag = flag.NewFlagSet(c.Name, flag.ExitOnError)
	}
	if c.Flag.Usage == nil {
		c.Flag.Usage = func() {
			if err := c.usage(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
	if !c.Flag.Parsed() {
		if err := c.Flag.Parse(args); err != nil {
			return fmt.Errorf("Commander.Main flag parsing failure: %v", err)
		}
	}
	if len(args) < 1 {
		if err := c.usage(); err != nil {
			return err
		}
		return fmt.Errorf("Not enough arguments provided")
	}

	if args[0] == "help" {
		return c.help(args[1:])
	}

	// first, try a sub-commander
	for _, cmd := range c.Commanders {
		n := cmd.Name
		if n == args[0] {
			return cmd.Run(args[1:])
		}
	}

	// then, try an internal command
	for _, cmd := range c.Commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			return nil
		}
	}

	// then try out an external one
	bin, err := exec.LookPath(c.FullName() + "-" + args[0])
	if err == nil {
		cmd := exec.Command(bin, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	// TODO: try an alias
	//...

	return fmt.Errorf("unknown subcommand %q\nRun 'help' for usage.\n", args[0])
}

func (c *Commander) usage() error {
	err := tmpl(os.Stderr, usageTemplate, c)
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// help implements the 'help' command.
func (c *Commander) help(args []string) error {
	if len(args) == 0 {
		return c.usage()
	}
	if len(args) != 1 {
		return fmt.Errorf("usage: %v help command\n\nToo many arguments given.\n", c.Name)
	}

	arg := args[0]

	for _, cmd := range c.Commanders {
		n := cmd.Name
		if strings.HasPrefix(n, c.Name+"-") {
			n = n[len(c.Name+"-"):]
		}
		if n == arg {
			return cmd.help(args[1:])
		}
	}

	for _, cmd := range c.Commands {
		if cmd.Name() == arg {
			c := struct {
				*Command
				ProgramName string
			}{cmd, c.Name}
			return tmpl(os.Stdout, helpTemplate, c)
		}
	}

	return fmt.Errorf("Unknown help topic %#q.  Run '%v help'.\n", arg, c.Name)
}

// FullName returns the full name of the commander, prefixed with its parent commanders, if any.
func (c *Commander) FullName() string {
	n := c.Name
	if c.Parent != nil {
		n = c.Parent.FullName() + "-" + n
	}
	return n
}

// A Command is an implementation of a subcommand.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)

	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Usage prints the usage details to the standard error output.
func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
}

// FlagOptions returns the flag's options as a string
func (c *Command) FlagOptions() string {
	var buf bytes.Buffer
	c.Flag.SetOutput(&buf)
	fmt.Fprintf(&buf, "\noptions:\n")
	if c.Flag.Usage != nil {
		c.Flag.Usage()
	} else {
		c.Flag.PrintDefaults()
	}
	return string(buf.Bytes())
}

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as importpath.
func (c *Command) Runnable() bool {
	return c.Run != nil
}

var usageTemplate = `Usage:

	{{.Name}} command [arguments]

The commands are:
{{range .Commands}}{{if .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}
{{range .Commanders}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "{{$.Name}} help [command]" for more information about a command.

Additional help topics:
{{range .Commands}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "{{.Name}} help [topic]" for more information about that topic.

`

var helpTemplate = `{{if .Runnable}}Usage: {{.ProgramName}} {{.UsageLine}}

{{end}}{{.Long | trim}}
{{.FlagOptions}}
`

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) error {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	return t.Execute(w, data)
}
