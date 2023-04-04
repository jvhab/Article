package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"grpc-article/internal/model"
	"strings"
	"testing"
)

func TestPostgresRepository(t *testing.T) {
	t.Run("Test create article", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")

		testArticle := &model.Article{
			Title:       "TestArticle",
			Description: "TestDescription",
			Body:        "TestBody",
			Counts:      10,
			TagList:     []string{"TestTag01", "TestTag02", "TestTag03"},
		}

		mock.ExpectExec("INSERT INTO article (title, description, body, counts, tag_list) VALUES ($1, $2, $3, $4, $5)").WithArgs(testArticle.Title, testArticle.Description, testArticle.Body, testArticle.Counts, strings.Join(testArticle.TagList, ",")).WillReturnResult(sqlmock.NewResult(1, 1))
		repo := NewPgRepo(sqlxDB)
		ctx := context.Background()
		result, err := repo.Create(ctx, testArticle)
		assert.NoError(t, err)
		assert.Equal(t, result, testArticle.Title)
	})

	t.Run("Test get article", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()
		sqlxDB := sqlx.NewDb(db, "sqlmock")

		testArticle := &model.Article{
			Title:       "TestArticle",
			Description: "TestDescription",
			Body:        "TestBody",
			Counts:      10,
			TagList:     []string{"TestTag01", "TestTag02", "TestTag03"},
		}
		rows := sqlmock.NewRows([]string{"title", "description", "body", "counts"}).AddRow("TestArticle", "TestDescription", "TestBody", 10)
		mock.ExpectQuery("SELECT title, description, body, counts, tag_list FROM article WHERE title = $1").WithArgs(testArticle.Title).WillReturnRows(rows)
		repo := NewPgRepo(sqlxDB)
		ctx := context.Background()
		article, err := repo.Get(ctx, testArticle.Title)
		assert.NoError(t, err)
		assert.Equal(t, article.Title, testArticle.Title)
	})

	t.Run("Test update article", func(t *testing.T) {

	})

	t.Run("Test delete article", func(t *testing.T) {

	})
}
