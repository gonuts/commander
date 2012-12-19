go-commander
============

``go-commander`` is a spin off of (golang)[http://golang.org] ``go tool`` infrastructure to provide commands and sub-commands.

A ``commander.Commander`` has a ``Commands`` field holding ``[]*commander.Command`` subcommands, referenced by name from the command line and an optional ``[]*commander.Commander`` ``Commanders`` field for nested commanders.

So a ``Commander`` can have sub commanders.

An example is provided by the (hwaf)[https://github.com/mana-fwk/hwaf] command.

So you can have, /e.g./
```sh
$ mycmd action1 [options...]
$ mycmd subcmd1 action1 [options...]
```

## Documentation
Is available on (godoc)[http://godoc.org/github.com/sbinet/go-commander]

## Installation
Is performed with the usual:
```sh
$ go get github.com/sbinet/go-commander
```

## TODO

- automatically generate the bash/zsh/csh autocompletion lists


