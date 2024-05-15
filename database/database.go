package database

import (
	"log"
	"os"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*influxdb3.Client, error) {
	// Carrega o arquivo .env que está um diretório acima
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro carregando o arquivo .env")
	}

	url := os.Getenv("INFLUXDB_URL")
	token := os.Getenv("INFLUXDB_TOKEN")
	database := os.Getenv("INFLUXDB_DATABASE")

	influxdb3Client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:     url,
		Token:    token,
		Database: database,
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
		return nil, err
	}

	return influxdb3Client, nil
}
