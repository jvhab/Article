package repository

import (
	"context"
	"grpc-article/internal/model"
)

type PostgresRepository interface {
	Create(context.Context, *model.Article) (string, error)
	Get(context.Context, string) (*model.Article, error)
	Update(context.Context, *model.Article) (*model.Article, error)
	Delete(context.Context, string) (string, error)
}
