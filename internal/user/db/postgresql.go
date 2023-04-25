package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/Salavei/golang_advanced_restapi/internal/user"
	"github.com/Salavei/golang_advanced_restapi/pkg/client/postgresql"
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, user user.User) error {

	r.logger.Debug("create author")
	q := ``
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	if err := r.client.QueryRow(ctx, q, user.Username).Scan(&user.ID); err != nil {
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

func (r *repository) FindAll(ctx context.Context) (u []user.User, err error) {

	q := ``
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]user.User, 0)

	for rows.Next() {
		var usr user.User

		err = rows.Scan(&usr.ID, &usr.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, usr)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (u user.User, err error) {

	return u, nil
}

func (r *repository) Update(ctx context.Context, user user.User) error {

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Storage {
	return &repository{
		client: client,
		logger: logger,
	}
}
