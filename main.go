package main

import (
	"database/sql"
	"fmt"
	"gator/internal/cli"
	"gator/internal/config"
	"gator/internal/database"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Read()
	// Открываем подключение к БД
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка подключения к БД: %v\n", err)
		os.Exit(1)
	}
	// Проверяем соединение
	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "БД не отвечает: %v\n", err)
		os.Exit(1)
	}
	// Инициализируем sqlc Queries
	dbQueries := database.New(db)
	// Инициализируем state
	state := &cli.State{
		DB:  dbQueries,
		CFG: &cfg,
	}
	cmds := &cli.Commands{
		Handlers: make(map[string]func(*cli.State, cli.Command) error),
	}
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: укажите команду")
		os.Exit(1)
	}

	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err1 := cmds.Run(state, cmd)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err1)
		os.Exit(1)
	}
}
