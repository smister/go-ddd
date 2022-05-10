package repository

import "context"

type AccountRepoIFace interface {
	Translation(ctx context.Context, callback func(ctx context.Context) error) error
}
