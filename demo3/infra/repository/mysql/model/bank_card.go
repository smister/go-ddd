package model

type BankCard struct {
	ID        string `json:"id"`         // 银行卡号
	AccountID uint64 `json:"account_id"` // 账号ID
	BankName  string `json:"bank_name"`  // 银行名称
	Status    bool   `json:"status"`     // 开启状态
}

func (BankCard) TableName() string {
	return "bank_card"
}
