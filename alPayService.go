package epaygo

import (
    "epaygo/al"
)
type AlPayService struct {
}

func (a *AlPayService) DirectPayAl(dto *(al.AlDirectPayDto)) (result string, err error) {
    return "111",nil
}

