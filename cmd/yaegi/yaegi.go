package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/traefik/yaegi/interp"
)

const (
	Extract = "extract"
	Help    = "help"
	Run     = "run"
	Test    = "test"
	Version = "version"
)

var version = "devel"

func main() {
	var cmd string
	var err error
	var exitCode int

	log.SetFlags(log.Lshortfile)

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case Extract:
		err = extractCmd(os.Args[2:])
	case Help, "-h", "--help":
		err = help(os.Args[2:])
	case Run:
		err = run(os.Args[2:])
	case Test:
		err = test(os.Args[2:])
	case Version:
		fmt.Println(version)
	default:

		cmd = Run
		err = run(os.Args[1:])
	}

	if err != nil && !errors.Is(err, flag.ErrHelp) {
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s: %w", cmd, err))
		if p, ok := err.(interp.Panic); ok {
			fmt.Fprintln(os.Stderr, string(p.Stack))
		}
		exitCode = 1
	}
	os.Exit(exitCode)
}
