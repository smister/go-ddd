package account

import "context"

// AccountRepoIFace 任务仓储接口定义
type AccountRepoIFace interface {
	Translation(ctx context.Context, callback func(ctx context.Context) error) error
	GetAccount(ctx context.Context, accountID uint64) (*Account, error)
	Update(ctx context.Context, account *Account) error
}
