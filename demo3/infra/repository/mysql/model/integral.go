package model

import (
	"github.com/smister/go-ddd/demo3/domain/integral"
)

type Integral struct {
	ID       uint64  `json:"id"`
	Integral float32 `json:"integral"`
}

func (Integral) TableName() string {
	return "integral"
}

func IntegralMDToDO(md *Integral) *integral.Integral {
	do := &integral.Integral{
		ID:       md.ID,
		Integral: md.Integral,
	}

	return do
}
