package repository

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"grpc-article/internal/model"
	"strings"
)

type pgRepo struct {
	db *sqlx.DB
}

func NewPgRepo(db *sqlx.DB) PostgresRepository {
	return &pgRepo{
		db: db,
	}
}

func (pg *pgRepo) Create(ctx context.Context, article *model.Article) (string, error) {
	query := `INSERT INTO article (title, description, body, counts, tag_list) VALUES ($1, $2, $3, $4, $5)`
	result, err := pg.db.ExecContext(ctx, query, article.Title, article.Description, article.Body, article.Counts, strings.Join(article.TagList, ","))
	if err != nil {
		return "", errors.Wrap(err, "articleRepo.Create.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", errors.Wrap(err, "articleRepo.Create.RowsAffected")
	}
	if rowsAffected != 1 {
		return "", errors.Wrap(err, "articleRepo.Create.RowsAffected")
	}
	return article.Title, nil
}

func (pg *pgRepo) Get(ctx context.Context, title string) (*model.Article, error) {
	query := `SELECT title, description, body, counts, tag_list FROM article WHERE title = $1`
	article := &model.Article{}
	err := pg.db.QueryRowxContext(ctx, query, title).StructScan(article)
	if err != nil {
		return &model.Article{}, errors.Wrap(err, "articleRepo.Get.QueryRowxContext")
	}
	return article, nil
}

func (pg *pgRepo) Update(ctx context.Context, article *model.Article) (*model.Article, error) {
	query := `UPDATE article SET description = COALESCE(NULLIF($1, ''), description),
                   body = COALESCE(NULLIF($2, ''), body),
                   counts = COALESCE(NULLIF($3, ''), counts),
                   tag_list = COALESCE(NULLIF($4, ''), tag_list) 
               WHERE title = $5
                   `
	updateArticle := &model.Article{}
	err := pg.db.GetContext(ctx, updateArticle, query, article.Description, article.Body, article.Counts, strings.Join(article.TagList, ","))
	if err != nil {
		return &model.Article{}, errors.Wrap(err, "articleRepo.Update.GetContext")
	}
	return updateArticle, nil
}

func (pg *pgRepo) Delete(ctx context.Context, title string) (string, error) {
	query := `DELETE FROM article WHERE title = $1`
	result, err := pg.db.ExecContext(ctx, query, title)
	if err != nil {
		return "", errors.Wrap(err, "articleRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", errors.Wrap(err, "articleRepo.Delete.RowsAffected")
	}
	if rowsAffected != 1 {
		return "", errors.Wrap(err, "article.Delete.RowsAffected")
	}
	return title, nil
}
