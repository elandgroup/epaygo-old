package main

import (
	"epaygo/helper"
	"fmt"
	"net/http"

	"epaygo"

	"github.com/labstack/echo"
)

func DirectPay(c echo.Context) error {
	directPayDto := new(AlDirectPayDto)
	if err := c.Bind(directPayDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, APIResult{Success: false, Error: APIError{Code: "20002"}})
	}
	directPayDto.OutTradeNo = helper.Uuid32()
	fmt.Printf("%#v", directPayDto)

	payService := new(epaygo.AlPayService)

	directPayDtoP := structToMap(directPayDto)

	if result, err := payService.DirectPay(directPayDtoP); err != nil {
		return c.JSON(http.StatusOK, APIResult{Success: false, Error: APIError{Code: err.Error()}})
	} else {
		fmt.Println(result)
		//c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		return c.JSON(http.StatusOK, APIResult{Success: true, Result: result})
		//c.String(http.StatusOK, result)
	}

}
