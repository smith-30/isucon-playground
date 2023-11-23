package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echoのインスタンスを作成
	e := echo.New()

	// ルートのハンドラーを設定
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// サーバーを開始
	e.Start(":8080")
}