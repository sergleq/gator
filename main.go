package main

import (
	"fmt"
	"gator/internal/cli"
	"gator/internal/config"
	"os"
)

func main() {
	cfg := config.Read()

	state := &cli.State{
		ConfigFile: &cfg,
	}

	cmds := &cli.Commands{
		Handlers: make(map[string]func(*cli.State, cli.Command) error),
	}
	cmds.Register("login", cli.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: укажите команду")
		os.Exit(1)
	}

	cmd := cli.Command{
		Name:     os.Args[1],
		Argument: os.Args[2:],
	}

	err := cmds.Run(state, cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}
