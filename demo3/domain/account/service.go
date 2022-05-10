package account

import (
	"context"
	"errors"
)

// 账户领域服务接口
type AccountServiceIFace interface {
	TransferAccounts(ctx context.Context, sourceAccountID uint64, destAccountID uint64, amount float32) error
	UpdateAddress(ctx context.Context, accountID uint64, province, city, district, address string) error
	// 增加银行卡
	AddBankCard(ctx context.Context, accountID uint64, bankNumber, bankName string) error
	// 移除银行卡
	RemoveBankCard(ctx context.Context, accountID uint64, bankNumber string) error
	// 启用银行卡
	EnableBankCard(ctx context.Context, accountID uint64, bankNumber string) error
	// 禁用银行卡
	DisableBankCard(ctx context.Context, accountID uint64, bankNumber string) error
	// 扣除余额
	DecreaseBalance(ctx context.Context, accountID uint64, amount float32) error
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

// 更新地址
func (ds *AccountService) UpdateAddress(ctx context.Context, accountID uint64, province, city, district, address string) error {
	// 读取更新账号
	account, err := ds.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	if err := account.UpdateAddress(province, city, district, address); err != nil {
		return err
	}

	// 存储账号聚合
	return ds.accountRepo.Update(ctx, account)
}

// 增加银行卡
func (ds *AccountService) AddBankCard(ctx context.Context, accountID uint64, bankNumber, bankName string) error {
	// 读取更新账号
	account, err := ds.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	if err := account.AddBankCard(bankNumber, bankName); err != nil {
		return err
	}

	// 存储账号聚合
	return ds.accountRepo.Update(ctx, account)
}

// 移除银行卡
func (ds *AccountService) RemoveBankCard(ctx context.Context, accountID uint64, bankNumber string) error {
	// 读取更新账号
	account, err := ds.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	if err := account.RemoveBankCard(bankNumber); err != nil {
		return err
	}

	// 存储账号聚合
	return ds.accountRepo.Update(ctx, account)
}

// 启用银行卡
func (ds *AccountService) EnableBankCard(ctx context.Context, accountID uint64, bankNumber string) error {
	// 读取更新账号
	account, err := ds.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	if err := account.EnableBankCard(bankNumber); err != nil {
		return err
	}

	// 存储账号聚合
	return ds.accountRepo.Update(ctx, account)
}

// 禁用银行卡
func (ds *AccountService) DisableBankCard(ctx context.Context, accountID uint64, bankNumber string) error {
	// 读取更新账号
	account, err := ds.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	if err := account.DisableBankCard(bankNumber); err != nil {
		return err
	}

	// 存储账号聚合
	return ds.accountRepo.Update(ctx, account)
}

// 扣除账号金额
func (ds *AccountService) DecreaseBalance(ctx context.Context, accountID uint64, amount float32) error {
	if amount <= 0 {
		return errors.New("扣除金额不能小于0")
	}

	// 读取账号A
	accountAggr, err := ds.accountRepo.GetAccount(ctx, accountID)
	if err != nil {
		return err
	}

	// 减少账号A的金额
	if err := accountAggr.DecreaseBalance(amount); err != nil {
		return err
	}

	return ds.accountRepo.Update(ctx, accountAggr)
}
