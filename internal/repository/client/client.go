package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-article/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()
	c := proto.NewArticleServiceClient(conn)
	ctx := context.Background()
	tmp := &proto.Article{
		Title: "TestArticle",
	}
	tmpArticle := &proto.CreateArticleRequest{
		Article: tmp,
	}
	r, err := c.CreateArticle(ctx, tmpArticle)
	if err != nil {
		log.Fatalf("Dont add new article")
	}
	fmt.Println(r.GetTitle())

}
