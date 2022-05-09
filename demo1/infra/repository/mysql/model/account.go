package model

import (
	"github.com/smister/go-ddd/demo1/common/pkg/itool"
	"github.com/smister/go-ddd/demo1/domain/account"
)

type Account struct {
	ID     uint64  `json:"id"`
	Amount float32 `json:"amount"`
}

func (Account) TableName() string {
	return "account"
}

func AccountMDToDO(md *Account) *account.Account {
	do := &account.Account{}
	itool.Convert(md, do)
	return do
}
