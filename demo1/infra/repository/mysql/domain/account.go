package domain

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/smister/go-ddd/demo1/common/pkg/db"
	"github.com/smister/go-ddd/demo1/common/pkg/itool"
	"github.com/smister/go-ddd/demo1/common/vars"
	"github.com/smister/go-ddd/demo1/domain/account"
	"github.com/smister/go-ddd/demo1/infra/repository/mysql/model"
)

type AccountRepo struct {
}

func (a *AccountRepo) GetAccount(ctx context.Context, accountID uint64) (*account.Account, error) {
	dao := itool.GetDBWithContext(ctx)
	accountMD := &model.Account{}
	if err := dao.Where("id=?", accountID).First(accountMD).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return model.AccountMDToDO(accountMD), nil
}

func (a *AccountRepo) Update(ctx context.Context, account *account.Account) error {
	dao := itool.GetDBWithContext(ctx)
	return dao.Model(&model.Account{}).Where("id = ?", account.ID).Updates(map[string]interface{}{
		"amount": account.Amount,
	}).Error
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
