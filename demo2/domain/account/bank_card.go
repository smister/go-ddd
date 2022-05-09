package account

import "errors"

// domain/account/bank_cardgo
// 银行卡实体
type BankCard struct {
	ID       string `json:"id"`        // 银行卡号
	BankName string `json:"bank_name"` // 银行名称
	Status   bool   `json:"status"`    // 开启状态
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
	return nil
}

// 禁用银行卡
func (c *BankCard) Disable() error {
	c.Status = false
	return nil
}
