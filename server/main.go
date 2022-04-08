package main

import (
	"github.com/aocm/vue-go-spa-sample/server/handler"
	"github.com/aocm/vue-go-spa-sample/server/infra/accessor"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()
	InitDb()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.POST("/yamabiko", handler.YamabikoAPI())
	e.GET("/yamabiko", handler.GetHistory())
	e.OPTIONS("/yamabiko", handler.OptionsCheck())

	// サーバー起動
	e.Start(":8000")
}

func InitDb() {
	accessor.AccessDB(accessor.MysqlAccessor{})
}
