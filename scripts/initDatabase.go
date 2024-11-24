package main

import (
	"CoreBaseGo/internal/infrastructure/config"
	"CoreBaseGo/internal/infrastructure/persistence"
)

func main() {
	config.InitConfig("./configs/", "env", ".env")
	persistence.InitDatabase()
}
