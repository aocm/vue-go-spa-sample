package handler

import (
	"fmt"
	"net/http"

	"github.com/aocm/vue-go-spa-sample/server/infra/accessor"
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
		saveMessge(param.Message)
		return c.JSON(http.StatusOK, map[string]interface{}{"message": param.Message})
	}
}

// 一覧取得（20件まで、新しい順）
func GetHistory() echo.HandlerFunc {
	var dbmap = accessor.ConnectDB(accessor.MysqlAccessor{})
	return func(c echo.Context) error {
		var histroy []accessor.Message
		_, err := dbmap.Select(&histroy, "select * from message order by id desc limit 20")
		if err != nil {
			fmt.Println("error! %v", err)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"history": histroy})
	}
}

func saveMessge(text string) bool {
	var dbmap = accessor.ConnectDB(accessor.MysqlAccessor{})
	var txmap, err = accessor.StartTransaction(accessor.MysqlAccessor{}, dbmap)
	message := accessor.Message{
		Text: text,
	}
	err = txmap.Insert(&message)
	defer dbmap.Db.Close()

	fmt.Println(err)
	if err == nil {
		txmap.Commit()
		return true
	} else {
		txmap.Rollback()
		return false
	}
}

func OptionsCheck() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
