package author

import (
	"context"
	"errors"
	"fmt"
	"github.com/Salavei/golang_advanced_restapi/internal/author"
	"github.com/Salavei/golang_advanced_restapi/pkg/client/postgresql"
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
	"github.com/jackc/pgconn"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", "")
}

func (r *repository) Create(ctx context.Context, author *author.Author) error {
	r.logger.Debug("create author")
	q := `
		INSERT INTO author (name) 
		VALUES ($1) 
		RETURNING id
		`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(newErr)
			return nil
		}
		return err
	}
	return nil

}

func (r *repository) FindAll(ctx context.Context) ([]author.Author, error) {
	q := `SELECT id, name FROM author`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)

	for rows.Next() {
		var ath author.Author

		err = rows.Scan(&ath.ID, &ath.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, ath)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, name string) (author.Author, error) {
	q := `SELECT id, name FROM author WHERE name = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	var ath author.Author
	if err := r.client.QueryRow(ctx, q, name).Scan(&ath.ID, &ath.Name); err != nil {
		return author.Author{}, err
	}
	return ath, nil
}

func (r *repository) Update(ctx context.Context, author author.Author) error {
	q := `UPDATE FROM author ........`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, name string) error {
	q := `DELETE FROM author WHERE name = $1`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	//
	//result, err := r.client.Query(ctx, q, name)
	//fmt.Println(result.)
	//fmt.Println(err)
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
