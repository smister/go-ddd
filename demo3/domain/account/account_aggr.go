package account

import (
	"errors"
	"github.com/smister/go-ddd/demo3/common/pkg/attr"
	"github.com/smister/go-ddd/demo3/common/pkg/itool"
)

// domain/account/account.go
// 账号聚合根
type Account struct {
	ID        uint64      `json:"id"`     // ID
	Amount    float32     `json:"amount"` // 金额
	Addr      *Address    // 值对象
	BankCards []*BankCard // 银行卡(实体)
	changes   attr.Attribute
}

func NewAccount(amount float32, province, city, district, address string) (*Account, error) {
	if amount < 0 {
		return nil, errors.New("账号金额不能小于0")
	}

	addr, err := NewAddress(province, city, district, address)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:     itool.GenerateSId(),
		Amount: amount,
		Addr:   addr,
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
	a.changes.Set("amount", a.Amount)
	return nil
}

// 增加金额
func (a *Account) IncreaseBalance(amount float32) error {
	if amount < 0 {
		return errors.New("增加余额必须大于0")
	}
	a.Amount = a.Amount + amount
	a.changes.Set("amount", a.Amount)
	return nil
}

// 更新账号地址
func (a *Account) UpdateAddress(province, city, district, address string) error {
	addr, err := NewAddress(province, city, district, address)
	if err != nil {
		return err
	}
	a.Addr = addr
	a.changes.Set("province", province)
	a.changes.Set("city", city)
	a.changes.Set("district", district)
	a.changes.Set("address", address)
	return nil
}

// 删除银行卡
func (a *Account) RemoveBankCard(bankNumber string) error {
	bankCards := make([]*BankCard, 0)
	for _, bankCard := range a.BankCards {
		if bankCard.ID != bankNumber {
			bankCards = append(bankCards, bankCard)
		}
	}

	if len(bankCards) == len(a.BankCards) {
		return errors.New("找不到对应银行卡")
	}

	a.BankCards = bankCards
	return nil
}

// 增加银行卡
func (a *Account) AddBankCard(bankNumber, bankName string) error {
	// 判断该银行是否已经存在
	for _, bankCard := range a.BankCards {
		if bankCard.ID == bankNumber {
			return errors.New("该银行卡已经存在")
		}
	}

	bankCard, err := NewBankCard(bankNumber, bankName)
	if err != nil {
		return err
	}

	a.BankCards = append(a.BankCards, bankCard)
	return nil
}

// 启用银行卡
func (a *Account) EnableBankCard(bankNumber string) error {
	opBankCard, err := a.getBankCard(bankNumber)
	if err != nil {
		return err
	}
	return opBankCard.Enable()
}

// 禁用银行卡
func (a *Account) DisableBankCard(bankNumber string) error {
	opBankCard, err := a.getBankCard(bankNumber)
	if err != nil {
		return err
	}
	return opBankCard.Disable()
}

// 读取账号内的bank
func (a *Account) getBankCard(bankNumber string) (*BankCard, error) {
	var opBankCard *BankCard
	for _, bankCard := range a.BankCards {
		if bankCard.ID == bankNumber {
			opBankCard = bankCard
			break
		}
	}

	if opBankCard == nil {
		return nil, errors.New("找不到银行卡")
	}

	return opBankCard, nil
}

// 获取变动c参数
func (a *Account) GetChanges() map[string]interface{} {
	changes := a.changes.Attributes()
	a.changes.Refresh()
	return changes
}
