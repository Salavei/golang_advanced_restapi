package author

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, author *Author) error
	FindAll(ctx context.Context) ([]Author, error)
	FindOne(ctx context.Context, name string) (Author, error)
	Update(ctx context.Context, author Author) error
	Delete(ctx context.Context, name string) error
}
