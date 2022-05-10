package app

import (
	"context"
	"github.com/smister/go-ddd/demo3/app/repository"
	"github.com/smister/go-ddd/demo3/common/vars"
	"github.com/smister/go-ddd/demo3/domain/account"
	"github.com/smister/go-ddd/demo3/domain/integral"
)

// 账号应用服务
type AccountService struct {
	accountDS   account.AccountServiceIFace
	integralDS  integral.IntegralServiceIFace
	accountRepo repository.AccountRepoIFace
}

func NewAccountService(accountDS account.AccountServiceIFace, integralDS integral.IntegralServiceIFace, accountRepo repository.AccountRepoIFace) *AccountService {
	return &AccountService{
		accountDS:   accountDS,
		integralDS:  integralDS,
		accountRepo: accountRepo,
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

// 余额购买积分
func (as *AccountService) BuyIntegral(ctx context.Context, accountID, integralID uint64, amount float32) error {
	if err := as.accountRepo.Translation(ctx, func(newCtx context.Context) error {
		// 扣除账号金额
		if err := as.accountDS.DecreaseBalance(newCtx, accountID, amount); err != nil {
			return err
		}

		// 增加积分金额
		if err := as.integralDS.IncreaseIntegral(newCtx, integralID, amount); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	// 发布事件
	return vars.EventPublisher.Publish(ctx, "account", "buy", map[string]interface{}{
		"accountID":  accountID,
		"integralID": integralID,
		"amount":     amount,
	})
}
