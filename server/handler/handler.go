package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// YamabikoParam は /yamabiko が受けとるJSONパラメータを定義します。
type YamabikoParam struct {
	Message string `json:"message"`
}

// YamabikoAPI は /api/hello のPost時のJSONデータ生成処理を行います。
func YamabikoAPI() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := new(YamabikoParam)
		if err := c.Bind(param); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"hello": param.Message})
	}
}
