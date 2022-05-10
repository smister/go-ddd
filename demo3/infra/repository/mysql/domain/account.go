package domain

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/smister/go-ddd/demo3/common/pkg/db"
	"github.com/smister/go-ddd/demo3/common/pkg/itool"
	"github.com/smister/go-ddd/demo3/common/vars"
	"github.com/smister/go-ddd/demo3/domain/account"
	"github.com/smister/go-ddd/demo3/infra/repository/mysql/model"
)

type AccountRepo struct {
}

func (a *AccountRepo) GetAccount(ctx context.Context, accountID uint64) (*account.Account, error) {
	dao := itool.GetDBWithContext(ctx)
	accountMD := &model.Account{}
	if err := dao.Where("id=?", accountID).First(accountMD).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// 关联银行卡
	bankCards := make([]*model.BankCard, 0)
	if err := dao.Where("account_id = ?", accountID).Find(&bankCards).Error; err != nil {
		return nil, err
	}

	return model.AccountMDToDO(accountMD, bankCards), nil
}

func (a *AccountRepo) Update(ctx context.Context, account *account.Account) error {
	dao := itool.GetDBWithContext(ctx)

	accountChanges := account.GetChanges()
	if len(accountChanges) > 0 { // 只更新变动的
		if err := dao.Model(&model.Account{}).Where("id = ?", account.ID).Updates(accountChanges).Error; err != nil {
			return err
		}
	}

	// 处理银行卡的更新
	if err := a.handleBankCards(ctx, account); err != nil {
		return err
	}

	return nil
}

func (a *AccountRepo) Translation(ctx context.Context, callback func(ctx context.Context) error) error {
	dao, err := db.NewDBEngine(vars.DatabaseSetting)
	if err != nil {
		return err
	}
	defer dao.Close()

	return dao.Transaction(func(tx *gorm.DB) error {
		ctx = itool.ContextWithDB(ctx, tx)
		return callback(ctx)
	})
}

// 处理银行卡
func (a *AccountRepo) handleBankCards(ctx context.Context, account *account.Account) error {
	// 获取当前账号的银行卡
	bankCardsIDs, err := a.getBankCardIDs(ctx, account.ID)
	if err != nil {
		return err
	}

	// 删除被删除的银行卡
	if err := a.deleteBankCards(ctx, bankCardsIDs, account); err != nil {
		return err
	}

	// 新增银行卡
	if err := a.addBankCards(ctx, bankCardsIDs, account); err != nil {
		return err
	}

	// 更新银行卡
	return a.updateBankCards(ctx, account)
}

// 删除银行卡
func (a *AccountRepo) deleteBankCards(ctx context.Context, bankCardIDs []string, account *account.Account) error {
	dao := itool.GetDBWithContext(ctx)
	currentBankCardMap := make(map[string]struct{})
	for _, bankCard := range account.BankCards {
		currentBankCardMap[bankCard.ID] = struct{}{}
	}

	deleteIDs := make([]string, 0)
	for _, bankCardID := range bankCardIDs {
		if _, exists := currentBankCardMap[bankCardID]; !exists {
			deleteIDs = append(deleteIDs, bankCardID)
		}
	}

	if len(deleteIDs) > 0 {
		return dao.Delete(&model.BankCard{}, deleteIDs).Error
	}
	return nil
}

// addBankCards 新增银行卡
func (a *AccountRepo) addBankCards(ctx context.Context, bankCardIDs []string, account *account.Account) error {
	dao := itool.GetDBWithContext(ctx)

	bankCardMap := make(map[string]struct{})
	for _, bankCardID := range bankCardIDs {
		bankCardMap[bankCardID] = struct{}{}
	}

	for _, bankCard := range account.BankCards {
		if _, exists := bankCardMap[bankCard.ID]; !exists {
			if err := dao.Create(&model.BankCard{
				ID:        bankCard.ID,
				AccountID: account.ID,
				BankName:  bankCard.BankName,
				Status:    bankCard.Status,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// updateBankCards 更新银行卡
func (a *AccountRepo) updateBankCards(ctx context.Context, account *account.Account) error {
	dao := itool.GetDBWithContext(ctx)
	for _, bankCard := range account.BankCards {
		bankCardChanges := bankCard.GetChanges()
		if len(bankCardChanges) > 0 {
			if err := dao.Model(&model.BankCard{}).Where("id = ?", bankCard.ID).Updates(bankCardChanges).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// getBankCardIDs 获取当前银行卡IDs
func (a *AccountRepo) getBankCardIDs(ctx context.Context, accountID uint64) ([]string, error) {
	dao := itool.GetDBWithContext(ctx)
	ids := make([]string, 0)
	if err := dao.Model(&model.BankCard{}).Where("account_id=?", accountID).Pluck("id", &ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}
