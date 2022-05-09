package account

import (
	"errors"
	"github.com/smister/go-ddd/demo1/common/pkg/itool"
)

// domain/account/account.go
// 账号聚合根
type Account struct {
	ID     uint64  // ID
	Amount float32 // 金额
}

func NewAccount(amount float32) (*Account, error) {
	if amount < 0 {
		return nil, errors.New("账号金额不能小于0")
	}
	return &Account{
		ID:     itool.GenerateSId(),
		Amount: amount,
	}, nil
}

// 扣除金额
func (a *Account) DecreaseBalance(amount float32) error {
	if amount < 0 {
		return errors.New("扣除金额必须大于0")
	}
	if amount > a.Amount {
		return errors.New("账户余额不足")
	}
	a.Amount = a.Amount - amount
	return nil
}

// 增加金额
func (a *Account) IncreaseBalance(amount float32) error {
	if amount < 0 {
		return errors.New("增加余额必须大于0")
	}
	a.Amount = a.Amount + amount
	return nil
}
