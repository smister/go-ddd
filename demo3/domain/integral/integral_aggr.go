package integral

import (
	"errors"
	"github.com/smister/go-ddd/demo3/common/pkg/attr"
)

// 积分聚合根
type Integral struct {
	ID       uint64  // ID
	Integral float32 // 积分
	changes  attr.Attribute
}

// 扣除积分
func (i *Integral) DecreaseIntegral(integral float32) error {
	if integral < 0 {
		return errors.New("扣除积分必须大于0")
	}
	if integral > i.Integral {
		return errors.New("账户积分不足")
	}
	i.Integral = i.Integral - integral
	i.changes.Set("integral", i.Integral)
	return nil
}

// 增加积分
func (i *Integral) IncreaseIntegral(integral float32) error {
	if integral < 0 {
		return errors.New("增加积分必须大于0")
	}
	i.Integral = i.Integral + integral
	// 追加amount的调整属性
	i.changes.Set("integral", i.Integral)
	return nil
}

// 获取变动c参数
func (i *Integral) GetChanges() map[string]interface{} {
	changes := i.changes.Attributes()
	i.changes.Refresh()
	return changes
}
