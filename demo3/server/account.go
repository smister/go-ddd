package server

import (
	"github.com/gin-gonic/gin"
	"github.com/smister/go-ddd/demo3/app"
	"github.com/smister/go-ddd/demo3/common/pkg/itool"
	"github.com/smister/go-ddd/demo3/domain/account"
	"github.com/smister/go-ddd/demo3/domain/integral"
	repoApp "github.com/smister/go-ddd/demo3/infra/repository/mysql/app"
	"github.com/smister/go-ddd/demo3/infra/repository/mysql/domain"
)

type AccountServer struct {
	AccountAppService *app.AccountService
}

func NewAccountServer() *AccountServer {
	return &AccountServer{
		AccountAppService: app.NewAccountService(
			account.NewAccountService(&domain.AccountRepo{}),
			integral.NewIntegralService(&domain.IntegralRepo{}),
			&repoApp.AccountRepo{},
		),
	}
}

// 转账
func (s *AccountServer) TransferAccounts(ctx *gin.Context) error {
	sourceAccountID := itool.StringToUint64(ctx.PostForm("source_account_id"))
	destAccountID := itool.StringToUint64(ctx.PostForm("dest_account_id"))
	amount := itool.StringToFloat32(ctx.PostForm("amount"))

	return s.AccountAppService.TransferAccounts(ctx.Request.Context(), sourceAccountID, destAccountID, amount)
}

// 更新账号地址
func (s *AccountServer) UpdateAddress(ctx *gin.Context) error {
	accountID := itool.StringToUint64(ctx.PostForm("account_id"))
	province := ctx.PostForm("province")
	city := ctx.PostForm("city")
	district := ctx.PostForm("district")
	address := ctx.PostForm("address")

	return s.AccountAppService.UpdateAddress(ctx.Request.Context(), accountID, province, city, district, address)
}

// 添加银行卡
func (s *AccountServer) AddBankCard(ctx *gin.Context) error {
	accountID := itool.StringToUint64(ctx.PostForm("account_id"))
	bankNumber := ctx.PostForm("bank_number")
	bankName := ctx.PostForm("bank_name")

	return s.AccountAppService.AddBankCard(ctx.Request.Context(), accountID, bankNumber, bankName)
}

// 移除银行卡
func (s *AccountServer) RemoveBankCard(ctx *gin.Context) error {
	accountID := itool.StringToUint64(ctx.PostForm("account_id"))
	bankNumber := ctx.PostForm("bank_number")

	return s.AccountAppService.RemoveBankCard(ctx.Request.Context(), accountID, bankNumber)
}

// 启用银行卡
func (s *AccountServer) EnableBankCard(ctx *gin.Context) error {
	accountID := itool.StringToUint64(ctx.PostForm("account_id"))
	bankNumber := ctx.PostForm("bank_number")

	return s.AccountAppService.EnableBankCard(ctx.Request.Context(), accountID, bankNumber)
}

// 禁用银行卡
func (s *AccountServer) DisableBankCard(ctx *gin.Context) error {
	accountID := itool.StringToUint64(ctx.PostForm("account_id"))
	bankNumber := ctx.PostForm("bank_number")

	return s.AccountAppService.DisableBankCard(ctx.Request.Context(), accountID, bankNumber)
}

// 购买积分
func (s *AccountServer) BuyIntegral(ctx *gin.Context) error {
	accountID := itool.StringToUint64(ctx.PostForm("account_id"))
	integralID := itool.StringToUint64(ctx.PostForm("integral_id"))
	amount := itool.StringToFloat32(ctx.PostForm("amount"))

	return s.AccountAppService.BuyIntegral(ctx.Request.Context(), accountID, integralID, amount)
}
