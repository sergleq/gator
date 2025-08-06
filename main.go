package main

import (
	"database/sql"
	"fmt"
	"gator/internal/cli"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/state"
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
	appstate := &state.State{
		DB:  dbQueries,
		CFG: &cfg,
	}
	cmds := &cli.Commands{
		Handlers: make(map[string]func(*state.State, cli.Command) error),
	}
	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerReset)
	cmds.Register("users", cli.HandlerUsers)
	cmds.Register("agg", cli.HandlerAgg)
	cmds.Register("addfeed", cli.MiddlewareLoggedIn(cli.HandlerAddFeed))
	cmds.Register("feeds", cli.HandlerFeeds)
	cmds.Register("follow", cli.HandlerFollow)
	cmds.Register("listfollow", cli.HandlerListFollows)
	cmds.Register("following", cli.HandlerFollowing)
	cmds.Register("unfollow", cli.MiddlewareLoggedIn(cli.HandlerUnfollow))

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: укажите команду")
		os.Exit(1)
	}
	cmd := cli.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	err1 := cmds.Run(appstate, cmd)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err1)
		os.Exit(1)
	}
}
