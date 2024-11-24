package main

import (
	"CoreBaseGo/internal/infrastructure/config"
	"CoreBaseGo/internal/infrastructure/http"
	"CoreBaseGo/internal/infrastructure/logging"
)

func init() {
	config.InitConfig("./configs/", "env", ".env")
	config.SetGlobalEnv()
	//utils.ConnectToRedis()
	logging.InitLogger()
}

func main() {
	http.StartServer()
}
