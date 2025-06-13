package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapsttructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapsttructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// read server configuration
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("Security::", viper.GetString("security.jwt.key"))

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v\n", err)
	}

	fmt.Println("Config port::", config.Server.Port)

	for _, db := range config.Database {
		fmt.Printf("database User::%s, Password::%s, Host::%s\n", db.User, db.Password, db.Host)
	}
}
