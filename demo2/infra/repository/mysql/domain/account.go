package domain

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/smister/go-ddd/demo2/common/pkg/db"
	"github.com/smister/go-ddd/demo2/common/pkg/itool"
	"github.com/smister/go-ddd/demo2/common/vars"
	"github.com/smister/go-ddd/demo2/domain/account"
	"github.com/smister/go-ddd/demo2/infra/repository/mysql/model"
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
	if err := dao.Model(&model.Account{}).Where("id = ?", account.ID).Updates(map[string]interface{}{
		"amount":   account.Amount,
		"province": account.Addr.Province,
		"district": account.Addr.District,
		"city":     account.Addr.City,
		"address":  account.Addr.Address,
	}).Error; err != nil {
		return err
	}

	// 处理银行卡的更新

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
