package main

import (
	"coba-gin/router"
	"fmt"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml")
	viper.SetConfigName("app.config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
	}

	var port = viper.GetString("app.port")

	fmt.Println("port: ", ":"+port)

	router.StartServer().Run(":" + port)
}
