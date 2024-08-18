package main

import "fmt"
import "github.com/spf13/viper"

func main() {
	fmt.Println("HW")
	viper.SetConfigFile("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	fmt.Println("Name:", viper.GetString("name"))
	fmt.Println("Version:", viper.GetString("version"))

}
