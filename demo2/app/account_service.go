package app

import (
	"context"
	"github.com/smister/go-ddd/demo2/common/vars"
	"github.com/smister/go-ddd/demo2/domain/account"
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

// 转账
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

// 更新账号地址
func (as *AccountService) UpdateAddress(ctx context.Context, accountID uint64, province, city, district, address string) error {
	if err := as.accountDS.UpdateAddress(ctx, accountID, province, city, district, address); err != nil {
		return err
	}

	// 发布事件
	return vars.EventPublisher.Publish(ctx, "account", "update-address", map[string]interface{}{
		"accountID": accountID,
	})
}

// 增加银行卡
func (as *AccountService) AddBankCard(ctx context.Context, accountID uint64, bankNumber, bankName string) error {
	if err := as.accountDS.AddBankCard(ctx, accountID, bankNumber, bankName); err != nil {
		return err
	}

	// 发布事件
	return vars.EventPublisher.Publish(ctx, "account", "add-bank-card", map[string]interface{}{
		"accountID":  accountID,
		"bankNumber": bankNumber,
	})
}

// 删除银行卡
func (as *AccountService) RemoveBankCard(ctx context.Context, accountID uint64, bankNumber string) error {
	if err := as.accountDS.RemoveBankCard(ctx, accountID, bankNumber); err != nil {
		return err
	}

	// 发布事件
	return vars.EventPublisher.Publish(ctx, "account", "remove-bank-card", map[string]interface{}{
		"accountID":  accountID,
		"bankNumber": bankNumber,
	})
}

// 启用银行卡
func (as *AccountService) EnableBankCard(ctx context.Context, accountID uint64, bankNumber string) error {
	return as.accountDS.EnableBankCard(ctx, accountID, bankNumber)
}

// 禁用银行卡
func (as *AccountService) DisableBankCard(ctx context.Context, accountID uint64, bankNumber string) error {
	return as.accountDS.DisableBankCard(ctx, accountID, bankNumber)
}
