package main

import (
	"github.com/aocm/vue-go-spa-sample/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.POST("/yamabiko", handler.YamabikoAPI())
	e.OPTIONS("/yamabiko", handler.OptionsCheck())

	// サーバー起動
	e.Start(":8000")
}
