package model

import (
	"github.com/smister/go-ddd/demo3/domain/account"
)

type Account struct {
	ID       uint64  `json:"id"`
	Amount   float32 `json:"amount"`
	Province string  `json:"province"` // 省
	City     string  `json:"city"`     // 市
	District string  `json:"district"` // 区
	Address  string  `json:"address"`  // 地址
}

func (Account) TableName() string {
	return "account"
}

func AccountMDToDO(md *Account, bankcards []*BankCard) *account.Account {
	do := &account.Account{
		ID:     md.ID,
		Amount: md.Amount,
		Addr: &account.Address{
			Address:  md.Address,
			Province: md.Province,
			District: md.District,
			City:     md.City,
		},
		BankCards: make([]*account.BankCard, 0),
	}

	for _, bankcard := range bankcards {
		do.BankCards = append(do.BankCards, &account.BankCard{
			ID:       bankcard.ID,
			BankName: bankcard.BankName,
			Status:   bankcard.Status,
		})
	}

	return do
}
