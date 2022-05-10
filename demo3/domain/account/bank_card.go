package account

import (
	"errors"
	"github.com/smister/go-ddd/demo3/common/pkg/attr"
)

// domain/account/bank_cardgo
// 银行卡实体
type BankCard struct {
	ID       string `json:"id"`        // 银行卡号
	BankName string `json:"bank_name"` // 银行名称
	Status   bool   `json:"status"`    // 开启状态
	changes  attr.Attribute
}

func NewBankCard(bankNumber string, bankName string) (*BankCard, error) {
	if bankNumber == "" {
		return nil, errors.New("银行号码不能为空")
	}

	if bankName == "" {
		return nil, errors.New("银行不能为空")
	}

	return &BankCard{
		ID:       bankNumber,
		BankName: bankName,
		Status:   true,
	}, nil
}

// 启动银行卡
func (c *BankCard) Enable() error {
	c.Status = true
	c.changes.Set("status", true)
	return nil
}

// 禁用银行卡
func (c *BankCard) Disable() error {
	c.Status = false
	c.changes.Set("status", false)
	return nil
}

// 获取变动参数
func (c *BankCard) GetChanges() map[string]interface{} {
	changes := c.changes.Attributes()
	c.changes.Refresh()
	return changes
}
