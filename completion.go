// Copyright 2021 The Go-Commander Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package commander

import (
	"flag"
	"os"
	"path/filepath"
	"strconv"

	"github.com/posener/complete"
)

// CompleterFrom creates a new tab-completer from the provided command.
func CompleterFrom(cmdr *Command) *complete.Command {
	const user = true
	return makeDefaultCompleter(cmdr, user)
}

func makeDefaultCompleter(cmdr *Command, user bool) *complete.Command {
	o := complete.Command{
		Sub:         make(complete.Commands),
		GlobalFlags: make(complete.Flags),
	}
	for _, sub := range cmdr.Subcommands {
		o.Sub[sub.Name()] = *sub.Complete
	}
	cmdr.Flag.VisitAll(func(f *flag.Flag) {
		k := "-" + f.Name
		switch {
		case user:
			// early circuit-breaker: if the user invoked this function,
			// cmdr.Parent may not be correctly setup.
			// better put that flag as a non-global one.
			o.Flags[k] = complete.PredictNothing
		case cmdr.Parent == nil:
			o.GlobalFlags[k] = complete.PredictNothing
		default:
			o.Flags[k] = complete.PredictNothing
		}
	})

	return &o
}

// completion returns true when [TAB]-completion is requested.
func completion() bool {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return false
	}

	switch name := filepath.Base(shell); name {
	case "bash":
		pt := os.Getenv("COMP_POINT")
		if pt == "" {
			return false
		}
		_, err := strconv.Atoi(pt)
		if err != nil {
			return false
		}
		return true

	case "csh", "tcsh":
		// TODO
	case "fish":
		// TODO
	case "zsh":
		// TODO
	}

	return false
}
