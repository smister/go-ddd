package account

import "errors"

// 地址值对象
type Address struct {
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 市
	District string `json:"district"` // 区
	Address  string `json:"address"`  // 地址
}

func NewAddress(province, city, district, address string) (*Address, error) {
	if province == "" || city == "" || district == "" || address == "" {
		return nil, errors.New("地址不能为空")
	}

	return &Address{
		Province: province,
		City:     city,
		District: district,
		Address:  address,
	}, nil
}

// Equals 判断2个地址是否相等
func (addr *Address) Equals(compareAddr *Address) bool {
	if addr.Province == compareAddr.Province && addr.City == compareAddr.City &&
		addr.District == compareAddr.District && addr.Address == compareAddr.Address {
		return true
	}
	return false
}
