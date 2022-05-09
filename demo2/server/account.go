package server

import (
	"github.com/gin-gonic/gin"
	"github.com/smister/go-ddd/demo2/app"
	"github.com/smister/go-ddd/demo2/common/pkg/itool"
	"github.com/smister/go-ddd/demo2/domain/account"
	"github.com/smister/go-ddd/demo2/infra/repository/mysql/domain"
)

type AccountServer struct {
	AccountAppService *app.AccountService
}

func NewAccountServer() *AccountServer {
	return &AccountServer{
		AccountAppService: app.NewAccountService(account.NewAccountService(&domain.AccountRepo{})),
	}
}

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
