package http

import (
	routes "CoreBaseGo/internal/interfaces/rest/importer"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

// StartServer run a gin server
func StartServer() {
	router := gin.Default()
	trustedProxies := []string{viper.GetString("PROXIES")}
	if len(trustedProxies) > 0 {
		err := router.SetTrustedProxies(trustedProxies)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
	}
	routes.V1(router)
	fmt.Println("Starting server on port: " + viper.GetString("PORT"))
	log.Fatal(router.Run(":" + viper.GetString("PORT")))
}
