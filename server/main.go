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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MainPage())
	e.POST("/yamabiko", handler.YamabikoAPI())

	// サーバー起動
	e.Start(":8000")
}
