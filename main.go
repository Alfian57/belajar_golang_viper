package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	config := viper.New()
	config.SetConfigFile(".env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	gopath := config.GetString("GOPATH")
	fmt.Println(gopath)
}
