package book

import (
	"context"
	"fmt"
	"github.com/Salavei/golang_advanced_restapi/internal/author"
	"github.com/Salavei/golang_advanced_restapi/internal/book"
	"github.com/Salavei/golang_advanced_restapi/pkg/client/postgresql"
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, book *book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindAll(ctx context.Context) (u []book.Book, err error) {
	q := `select id, name from book`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	books := make([]book.Book, 0)

	for rows.Next() {
		var bk book.Book

		err = rows.Scan(&bk.ID, &bk.Name)
		if err != nil {
			return nil, err
		}

		sq := `SELECT author.id, author.name 
			   FROM mm_book_author as ba
			   JOIN author ON author.id = ba.author_id
			   WHERE book_id = $1`
		authorRows, err := r.client.Query(ctx, sq, bk.ID)
		if err != nil {
			return nil, err
		}
		var authors = make([]author.Author, 0)
		for authorRows.Next() {
			var at author.Author

			err = authorRows.Scan(&at.ID, &at.Name)
			if err != nil {
				return nil, err
			}
			authors = append(authors, at)
		}
		bk.Authors = authors
		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *repository) Update(ctx context.Context, name book.Book) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) FindOne(ctx context.Context, id string) (book.Book, error) {
	q := `select ba.b_id, ba.name, array_agg(ba.a_id)
		  from
			(select b.id as b_id, b.name, a.id as a_id
			from mm_book_author
			inner join author as a on a.id = mm_book_author.author_id
			inner join book as b on b.id = mm_book_author.book_id
			where b.id = $1
			) as ba
		  group by ba.b_id, ba.name`
	var bk book.Book
	if err := r.client.QueryRow(context.Background(), q, id).Scan(&bk.ID, &bk.Name); err != nil {
		return book.Book{}, err
	}
	return bk, nil

}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
