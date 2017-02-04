package main

import (
	"epaygo"
	"epaygo/helper"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func DirectPayWx(c echo.Context) error {
	directPayDto := new(WxDirectPayDto)
	if err := c.Bind(directPayDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Message: "A required parameter is missing or doesn't have the right format" + "directPayDto"}})
	}
	directPayDto.OutTradeNo = helper.Uuid32()
	fmt.Printf("%#v", directPayDto)

	wxPayService := new(epaygo.WxPayService)

	dtoP := structToMap(directPayDto)

	if result, err := wxPayService.DirectPay(dtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		fmt.Println(result)
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}
