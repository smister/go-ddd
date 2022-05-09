package itool

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/smister/go-ddd/demo2/common/vars"
)

func GenerateSId() uint64 {
	return uint64(vars.Snowflake.GenerateId())
}

func ContextWithDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, "db", db)
}

func GetDBWithContext(ctx context.Context) *gorm.DB {
	dbI := ctx.Value("db")
	if db, ok := dbI.(*gorm.DB); ok {
		return db
	}

	return nil
}
