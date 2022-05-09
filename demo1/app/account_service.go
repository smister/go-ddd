package app

import (
	"context"
	"github.com/smister/go-ddd/demo1/common/vars"
	"github.com/smister/go-ddd/demo1/domain/account"
)

// 账号应用服务
type AccountService struct {
	accountDS account.AccountServiceIFace
}

func NewAccountService(accountDS account.AccountServiceIFace) *AccountService {
	return &AccountService{
		accountDS: accountDS,
	}
}

func (as *AccountService) TransferAccounts(ctx context.Context, sourceAccountID uint64, destAccountID uint64, amount float32) error {
	if err := as.accountDS.TransferAccounts(ctx, sourceAccountID, destAccountID, amount); err != nil {
		return err
	}

	// 发布事件，处理短信通知
	return vars.EventPublisher.Publish(ctx, "account", "transfer", map[string]interface{}{
		"sourceAccountID": sourceAccountID,
		"destAccountID":   destAccountID,
		"amount":          amount,
	})
}
