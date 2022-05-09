package account

import (
	"context"
	"errors"
)

// 账户领域服务接口
type AccountServiceIFace interface {
	TransferAccounts(ctx context.Context, sourceAccountID uint64, destAccountID uint64, amount float32) error
}

// 账户领域服务
type AccountService struct {
	accountRepo AccountRepoIFace
}

func NewAccountService(accountRepo AccountRepoIFace) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

// 转账
func (ds *AccountService) TransferAccounts(ctx context.Context, sourceAccountID uint64, destAccountID uint64, amount float32) error {
	if amount <= 0 {
		return errors.New("转账金额不能小于0")
	}

	// 读取账号A
	sourceAccount, err := ds.accountRepo.GetAccount(ctx, sourceAccountID)
	if err != nil {
		return err
	}

	// 读取账号B
	destAccount, err := ds.accountRepo.GetAccount(ctx, destAccountID)
	if err != nil {
		return err
	}

	// 减少账号A的金额
	if err := sourceAccount.DecreaseBalance(amount); err != nil {
		return err
	}

	// 增加账号B的金额
	if err := destAccount.IncreaseBalance(amount); err != nil {
		return err
	}

	// 开启事务
	return ds.accountRepo.Translation(ctx, func(newCtx context.Context) error {
		//  存储账号聚合A
		if err := ds.accountRepo.Update(newCtx, sourceAccount); err != nil {
			return err
		}

		//  存储账号聚合B
		if err := ds.accountRepo.Update(newCtx, destAccount); err != nil {
			return err
		}

		return nil
	})
}
