package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"grpc-article/config"
	"grpc-article/pkg/db/postgres"
	"log"
)

func main() {
	//ctx := context.Background()

	cfg := new(config.Config)
	err := cleanenv.ReadConfig("config/config.yaml", cfg)
	fmt.Println(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	psqlDB, err := postgres.NewPostgresConn(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer psqlDB.Close()
}
