package server

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"google.golang.org/grpc"
	"grpc-article/config"
	"grpc-article/internal/controller"
	"grpc-article/internal/handler"
	"grpc-article/internal/repository"
	"grpc-article/pkg/db/postgres"
	"grpc-article/proto"
	"log"
	"net"
)

func Run() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen")
	}
	cfg := &config.Config{}
	err = cleanenv.ReadConfig("config/config.yaml", cfg)
	fmt.Println(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	psqlDB, err := postgres.NewPostgresConn(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer psqlDB.Close()
	rep := repository.NewPgRepo(psqlDB)
	ctrl := controller.NewContorller(rep)
	hdlr := handler.NewHandler(ctrl)
	server := grpc.NewServer()
	proto.RegisterArticleServiceServer(server, hdlr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve")
	}
}
