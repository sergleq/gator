package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {

	cfg := config.Read()
	cfg.SetUser("Serg")

	updtCfg := config.Read()

	fmt.Println(updtCfg.DBURL, updtCfg.CurrentUser)
}
