package cli

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	"gator/internal/rssfeed"
	"time"

	"github.com/google/uuid"
)

type State struct {
	DB  *database.Queries
	CFG *config.Config
}
type Command struct {
	Name string
	Args []string
}

func HandlerAgg(s *State, cmd Command) error {
	feed, err := rssfeed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("ошибка загрузки ленты: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось получить пользователей: %w", err)
	}
	for _, u := range users {
		line := "* " + u.Name
		if u.Name == s.CFG.CurrentUser {
			line += " (current)"
		}
		fmt.Println(line)
	}
	return nil
}

func HandlerReset(s *State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось очистить базу: %w", err)
	}
	fmt.Println("Все пользователи удалены.")
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("регистрация требует имя пользователя")
	}

	username := cmd.Args[0]

	// создаём пользователя
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	_, err := s.DB.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	// обновляем конфиг
	err = s.CFG.SetUser(username)
	if err != nil {
		return fmt.Errorf("ошибка обновления конфигурации: %w", err)
	}

	// лог и подтверждение
	fmt.Printf("Пользователь %q создан и установлен текущим\n", username)
	fmt.Printf("DEBUG: %+v\n", newUser)

	return nil
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("the login handler expects a single argument, the username")
	}
	username := cmd.Args[0]
	// Проверяем, существует ли пользователь в БД
	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		// Ошибка — пользователь не найден
		return fmt.Errorf("пользователь %q не найден", username)
	}

	// Если найден — записываем в конфиг
	err = s.CFG.SetUser(username)
	if err != nil {
		return fmt.Errorf("ошибка записи в конфиг: %w", err)
	}

	fmt.Println("User has been set")
	return nil
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}
