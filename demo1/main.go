package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jeffcail/jsnowflake"
	"github.com/smister/go-ddd/demo1/common/pkg/db"
	"github.com/smister/go-ddd/demo1/common/pkg/event"
	"github.com/smister/go-ddd/demo1/common/pkg/itool"
	"github.com/smister/go-ddd/demo1/common/vars"
	"github.com/smister/go-ddd/demo1/server"
	"log"
	"net/http"
	"time"
)

func init() {
	if err := setupDB(); err != nil {
		log.Fatal(err)
	}

	if err := setupSnowflake(); err != nil {
		log.Fatal(err)
	}

	if err := setupEvent(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
		defer cancel()

		dao, err := db.NewDBEngine(vars.DatabaseSetting)
		if err != nil {
			log.Fatal(err)
		}
		defer dao.Close()
		ctx = itool.ContextWithDB(ctx, dao)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	r.POST("/account/transfer", func(c *gin.Context) {
		srv := server.NewAccountServer()
		if err := srv.TransferAccounts(c); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":   -1,
				"errorMsg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{"status": 1})
	})

	r.Run()
}

func setupDB() (err error) {
	vars.DatabaseSetting = &db.DatabaseSettingS{
		DBType:       "mysql",
		UserName:     "root",
		Password:     "123456",
		Host:         "127.0.0.1:3306",
		DBName:       "test",
		Charset:      "utf8mb4",
		ParseTime:    true,
		MaxIdleConns: 10,
		MaxOpenConns: 30,
	}
	return nil
}

func setupSnowflake() (err error) {
	vars.Snowflake, err = jsnowflake.NewMachine(1)
	return
}

func setupEvent() (err error) {
	vars.EventPublisher = &event.Event{}
	return
}
