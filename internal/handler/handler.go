package handler

import (
	"context"
	"grpc-article/internal/controller"
	"grpc-article/internal/model"
	"grpc-article/proto"
)

type Handler struct {
	ctrl *controller.Controller
	proto.UnimplementedArticleServiceServer
}

func NewHandler(ctrl *controller.Controller) *Handler {
	return &Handler{
		ctrl: ctrl,
	}
}

func (h *Handler) CreateArticle(ctx context.Context, request *proto.CreateArticleRequest) (*proto.CreateArticleResponse, error) {
	article := request.GetArticle()
	temp := &model.Article{
		Title:       article.GetTitle(),
		Description: article.GetDescription(),
		Body:        article.GetBody(),
		Counts:      int(article.GetCounts()),
		TagList:     article.GetTagList(),
	}
	name, err := h.ctrl.Create(ctx, temp)
	if err != nil {
		return &proto.CreateArticleResponse{}, err
	}
	return &proto.CreateArticleResponse{Title: name}, nil
}

func (h *Handler) GetArticle(ctx context.Context, request *proto.GetArticleRequest) (*proto.GetArticleResponse, error) {
	return nil, nil
}

func (h *Handler) UpdateArticle(ctx context.Context, request *proto.UpdateArticleRequest) (*proto.UpdateArticleResponse, error) {
	return nil, nil
}

func (h *Handler) DeleteArticle(ctx context.Context, request *proto.DeleteArticleRequest) (*proto.DeleteArticleResponse, error) {
	return nil, nil
}
