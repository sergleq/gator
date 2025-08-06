package state

import (
	"gator/internal/config"
	"gator/internal/database"
)

type State struct {
	DB  *database.Queries
	CFG *config.Config
}
