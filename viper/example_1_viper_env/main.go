package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	fmt.Println(viper.Get("PORT"))
}

