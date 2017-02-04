package epaygo

import (
	"epaygo/helper"
	"epaygo/wx"
	"errors"
	"net/http"
	"strings"

	"github.com/smallnest/goreq"
)

type WxPayService struct {
}

func (a *WxPayService) DirectPay(params map[string]string) (result string, err error) {

	wxPayData := a.BuildCommonparam(params)

	a.SetValue(wxPayData, wx.Body, params[wx.BodyMap])
	a.SetValue(wxPayData, wx.OutTradeNo, params[wx.OutTradeNoMap])
	a.SetValue(wxPayData, wx.TotalFee, params[wx.TotalFeeMap])
	a.SetValue(wxPayData, wx.AuthCode, params[wx.AuthCodeMap])
	a.SetValue(wxPayData, wx.DeviceInfo, params[wx.DeviceInfoMap])

	a.SetValue(wxPayData, wx.Detail, params[wx.DetailMap])
	a.SetValue(wxPayData, wx.Attach, params[wx.AttachMap])
	a.SetValue(wxPayData, wx.FeeType, params[wx.FeeTypeMap])
	a.SetValue(wxPayData, wx.GoodsTag, params[wx.GoodsTagMap])
	a.SetValue(wxPayData, wx.LimitPay, params[wx.LimitPayMap])

	a.SetValue(wxPayData, wx.Sign, wxPayData.MakeSign(params[wx.KeyMap]))

	xmlParam := wxPayData.ToXml()
	req, body, reqErr := goreq.New().Post(wx.MicroPay_Url).ContentType("xml").SendRawString(xmlParam).End()

	return a.ParseResult(req, body, reqErr, params[wx.KeyMap])

}

func (a *WxPayService) BuildCommonparam(params map[string]string) wx.WxPayData {
	wxPayData := wx.NewWxPayData()
	a.SetValue(*wxPayData, wx.SpbillCreateIp, params[wx.SpbillCreateIpMap])
	a.SetValue(*wxPayData, wx.AppId, params[wx.AppIdMap])
	a.SetValue(*wxPayData, wx.MchId, params[wx.MchIdMap])
	a.SetValue(*wxPayData, wx.SubAppId, params[wx.SubAppIdMap])
	a.SetValue(*wxPayData, wx.SubMchId, params[wx.SubMchIdMap])

	a.SetValue(*wxPayData, wx.NonceStr, helper.Uuid32())
	return *wxPayData
}

func (a *WxPayService) SetValue(wxPayData wx.WxPayData, key string, value string) {
	if len(strings.TrimSpace(value)) != 0 {
		wxPayData.SetValue(key, value)
	}
}

func (a *WxPayService) ParseResult(req goreq.Response, body string, reqErrs []error, key string) (result string, err error) {
	//serviceResult := ServiceResult{Result: nil, Success: ResultType.Unknown, Error: APIError{Code: 10004, Message: "", Details: nil}}
	wxResponse := wx.NewWxPayData()
	if err != nil {
		return "", reqErrs[0]
	}
	if req.StatusCode == http.StatusOK {
		if _, err := wxResponse.FromXml(body, key); err != nil {
			return "", errors.New("The request failed, please check whether the network is normal")
		}

		if wxResponse == nil {
			return "", errors.New("The request failed, please check whether the network is normal")
		} else {
			if len(wxResponse.GetValue(wx.ReturnCode)) == 0 || strings.ToUpper(wxResponse.GetValue(wx.ReturnCode)) != "SUCCESS" {
				return wxResponse.ToJson(), errors.New("The request failed, please check whether the network is normal")
			}
			if len(wxResponse.GetValue(wx.ResultCode)) != 0 && strings.ToUpper(wxResponse.GetValue(wx.ResultCode)) == "SUCCESS" {
				return wxResponse.ToJson(), nil
			} else {
				errCode := wxResponse.GetValue(wx.ErrCode)
				if errCode == wx.SystemError || errCode == wx.BankError || errCode == wx.UserPaying {
					return wxResponse.ToJson(), errors.New("result is unknown")
				} else {
					return wxResponse.ToJson(), errors.New(errCode)
				}
			}
		}
	} else {
		return "", errors.New("The request failed, please check whether the network is normal")
	}
	return "", errors.New("The request failed, please check whether the network is normal")
}
