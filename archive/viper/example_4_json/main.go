package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	var config ConfigExample

	viper.Unmarshal(&config) // Viper insert into config

	fmt.Println(config.CompilerOptions.Target)
}

type ConfigExample struct {
	CompilerOptions struct {
		Module string `json:"module"`
		Target string `json:"target"`
	} `json:"compilerOptions"`
	Exclude []string `json:"exclude"`
}
