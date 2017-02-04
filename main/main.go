package main

import (
	"net/http"

	"github.com/labstack/echo"
)

var (
	e = echo.New()
)

func main() {
	RegisterApi()
	e.Start(":1001")
}

func RegisterApi() {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello epay")
	})

	apiGroup := e.Group("/api")
	v1 := apiGroup.Group("/v1")

	wx := v1.Group("/wx")
	al := v1.Group("/al")

	wx.POST("/pay", DirectPayWx)

	// wx.POST("/query", api.OrderQuery)

	// wx.POST("/refund", api.Refund)

	// wx.POST("/reverse", api.Reverse)

	al.POST("/pay", DirectPayAL)

	al.POST("/query", OrderQueryAL)

	al.POST("/refund", RefundAL)

	al.POST("/reverse", ReverseAL)

}
