package configs

import (
	"github.com/spf13/viper"
	"log"
)

const (
	configFilePath = "../../configs"
	configFileName = "config"
)

type Configuration struct {
	Database DatabaseConfig
	Server ServerConfig
}

type DatabaseConfig struct {
	Driver string
	Host string
	Name string
	Port int
	User string
	Password string
	MigrationsPath string
}

type ServerConfig struct {
	Port int
}

var configuration Configuration

func init()  {
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configFilePath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalln("The config.yml was not unmarshalled correctly")
	}
}

func GetInitConfig() Configuration {
	return configuration
}

//func (dbConfig DatabaseConfig) GetDriver() string  {
//	return dbConfig.Driver
//}
//
//func (dbConfig DatabaseConfig) GetHost() string  {
//	return dbConfig.Host
//}
//
//func (dbConfig DatabaseConfig) GetPort() int  {
//	return dbConfig.Port
//}
//
//func (dbConfig DatabaseConfig) GetName() string  {
//	return dbConfig.Name
//}
//
//func (dbConfig DatabaseConfig) GetUser() string  {
//	return dbConfig.User
//}
//
//func (dbConfig DatabaseConfig) GetPassword() string  {
//	return dbConfig.Password
//}

