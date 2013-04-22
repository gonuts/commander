commander
============

``commander`` is a spin off of [golang](http://golang.org) ``go tool`` infrastructure to provide commands and sub-commands.

A ``commander.Commander`` has a ``Commands`` field holding ``[]*commander.Command`` subcommands, referenced by name from the command line and an optional ``[]*commander.Commander`` ``Commanders`` field for nested commanders.

So a ``Commander`` can have sub commanders.

An example is provided by the [hwaf](https://github.com/mana-fwk/hwaf) command.

So you can have, /e.g./
```sh
$ mycmd action1 [options...]
$ mycmd subcmd1 action1 [options...]
```

## Documentation
Is available on [godoc](http://godoc.org/github.com/gonuts/commander)

## Installation
Is performed with the usual:
```sh
$ go get github.com/gonuts/commander
```

## Example

See the simple ``my-cmd`` example command for how this all hangs
together [there](http://github.com/gonuts/commander/blob/master/examples/my-cmd/main.go):

```sh
$ my-cmd cmd1        
my-cmd-cmd1: hello from cmd1 (quiet=true)

$ my-cmd cmd1 -q
my-cmd-cmd1: hello from cmd1 (quiet=true)

$ my-cmd cmd1 -q=0
my-cmd-cmd1: hello from cmd1 (quiet=false)

$ my-cmd cmd2     
my-cmd-cmd2: hello from cmd2 (quiet=true)

$ my-cmd subcmd1 cmd1
my-cmd-subcmd1-cmd1: hello from subcmd1-cmd1 (quiet=true)

$ my-cmd subcmd1 cmd2
my-cmd-subcmd1-cmd2: hello from subcmd1-cmd2 (quiet=true)

$ my-cmd subcmd2 cmd1
my-cmd-subcmd2-cmd1: hello from subcmd2-cmd1 (quiet=true)

$ my-cmd subcmd2 cmd2
my-cmd-subcmd2-cmd2: hello from subcmd2-cmd2 (quiet=true)

$ my-cmd help
Usage:

	my-cmd command [arguments]

The commands are:

    cmd1        runs cmd1 and exits
    cmd2        runs cmd2 and exits

    subcmd1     subcmd1 sub-commander. does subcmd1 thingies
    subcmd2     subcmd2 sub-commander. does subcmd2 thingies

Use "my-cmd help [command]" for more information about a command.

Additional help topics:


Use "my-cmd help [topic]" for more information about that topic.


$ my-cmd help subcmd1
Usage:

	subcmd1 command [arguments]

The commands are:

    cmd1        runs cmd1 and exits
    cmd2        runs cmd2 and exits


Use "subcmd1 help [command]" for more information about a command.

Additional help topics:


Use "subcmd1 help [topic]" for more information about that topic.

```


## TODO

- automatically generate the bash/zsh/csh autocompletion lists


