package app

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/smister/go-ddd/demo3/common/pkg/db"
	"github.com/smister/go-ddd/demo3/common/pkg/itool"
	"github.com/smister/go-ddd/demo3/common/vars"
)

type AccountRepo struct {
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
