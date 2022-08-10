package config

import (
	"fmt"
	"os"
)

type APIConfig struct {
	URL string
}
type DBConfig struct {
	DataSourceName string
	Env            string
}

type FilePathConfig struct {
	FilePath string
}

type Config struct {
	APIConfig
	DBConfig
	FilePathConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")

	env := os.Getenv("ENV")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	c.DBConfig = DBConfig{
		DataSourceName: dsn,
		Env:            env,
	}
	c.APIConfig = APIConfig{
		URL: api,
	}
	c.FilePathConfig = FilePathConfig{
		FilePath: os.Getenv("FILE_PATH"),
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	fmt.Println("filepath:", cfg.FilePath)
	return cfg
}
