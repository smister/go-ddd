package integral

import (
	"context"
	"errors"
)

type IntegralServiceIFace interface {
	IncreaseIntegral(ctx context.Context, integralID uint64, integral float32) error
}

type IntegralService struct {
	integralRepo IntegralRepoIFace
}

func NewIntegralService(integralRepo IntegralRepoIFace) *IntegralService {
	return &IntegralService{
		integralRepo: integralRepo,
	}
}

// 增加积分
func (ds *IntegralService) IncreaseIntegral(ctx context.Context, integralID uint64, integral float32) error {
	if integral <= 0 {
		return errors.New("增加积分不能小于0")
	}

	// 读取积分聚合
	integralAggr, err := ds.integralRepo.GetIntegral(ctx, integralID)
	if err != nil {
		return err
	}

	// 增加积分
	if err := integralAggr.IncreaseIntegral(integral); err != nil {
		return err
	}

	// 积分聚合落库
	return ds.integralRepo.Update(ctx, integralAggr)
}
