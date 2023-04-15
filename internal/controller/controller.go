package controller

import (
	"context"
	"grpc-article/internal/model"
	"grpc-article/internal/repository"
)

type Controller struct {
	repo repository.PostgresRepository
}

func NewContorller(repo repository.PostgresRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) Create(ctx context.Context, article *model.Article) (string, error) {
	title, err := c.repo.Create(ctx, article)
	return title, err
}

func (c *Controller) Get(ctx context.Context, title string) (*model.Article, error) {
	article, err := c.repo.Get(ctx, title)
	return article, err
}

func (c *Controller) Update(ctx context.Context, articleNow *model.Article) (*model.Article, error) {
	article, err := c.repo.Update(ctx, articleNow)
	return article, err
}

func (c *Controller) Delete(ctx context.Context, title string) (string, error) {
	titleDel, err := c.repo.Delete(ctx, title)
	return titleDel, err
}
