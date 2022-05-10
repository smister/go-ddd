package domain

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/smister/go-ddd/demo3/common/pkg/itool"
	"github.com/smister/go-ddd/demo3/domain/integral"
	"github.com/smister/go-ddd/demo3/infra/repository/mysql/model"
)

type IntegralRepo struct {
}

func (i *IntegralRepo) GetIntegral(ctx context.Context, integralID uint64) (*integral.Integral, error) {
	dao := itool.GetDBWithContext(ctx)
	integralMD := &model.Integral{}
	if err := dao.Where("id=?", integralID).First(integralMD).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return model.IntegralMDToDO(integralMD), nil
}

func (i *IntegralRepo) Update(ctx context.Context, integral *integral.Integral) error {
	dao := itool.GetDBWithContext(ctx)

	integralChanges := integral.GetChanges()
	if len(integralChanges) > 0 { // 只更新变动的
		if err := dao.Model(&model.Integral{}).Where("id = ?", integral.ID).Updates(integralChanges).Error; err != nil {
			return err
		}
	}
	return nil
}
