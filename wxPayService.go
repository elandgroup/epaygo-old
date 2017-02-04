package epaygo

// import (
// 	"epaygo/wx"
// 	"oc-api-go/config"
// 	"strings"

// 	"epay.api.biz.go/epay.api.go/common"
// 	"epay.api.biz.go/epay.api.go/common/smt"
// 	"epay.api.biz.go/epay.api.go/common/wxConst"
// 	"github.com/smallnest/goreq"
// )

// type WxPayService struct {
// }

// func (a *WxPayService) DirectPay(params map[string]string) (result string, err error) {

// 	wxPayData := a.BuildCommonparam(params)

// 	a.SetValue(wxPayData, wx.Body, params[wx.BodyMap])
// 	a.SetValue(wxPayData, wxConst.OutTradeNo, dto.OutTradeNo)
// 	a.SetValue(wxPayData, wxConst.TotalFee, dto.TotalFee)
// 	a.SetValue(wxPayData, wxConst.AuthCode, dto.AuthCode)
// 	a.SetValue(wxPayData, wxConst.DeviceInfo, dto.DeviceInfo)

// 	a.SetValue(wxPayData, wxConst.Detail, dto.Detail)
// 	a.SetValue(wxPayData, wxConst.Attach, dto.Attach)
// 	a.SetValue(wxPayData, wxConst.FeeType, dto.FeeType)
// 	a.SetValue(wxPayData, wxConst.GoodsTag, dto.GoodsTag)
// 	a.SetValue(wxPayData, wxConst.LimitPay, dto.LimitPay)

// 	a.SetValue(wxPayData, wxConst.Sign, wxPayData.MakeSign(dto.Key))

// 	xmlParam := wxPayData.ToXml()
// 	req, body, reqErr := goreq.New().Post(config.Config.WX.MicroPay_Url).ContentType("xml").SendRawString(xmlParam).End()
// 	smt.Debug.Println("MicroPay_Url:", config.Config.WX.MicroPay_Url)
// 	smt.Debug.Println(req, body, reqErr)

// 	return a.ParseResult(req, body, reqErr, dto.Key)

// }

// func (a *WxPayService) BuildCommonparam(params map[string]string) WxPayData {
// 	wxPayData := NewWxPayData()
// 	a.SetValue(*wxPayData, wxConst.SpbillCreateIp, wxpaybaseDto.SpbillCreateIp)
// 	a.SetValue(*wxPayData, wxConst.AppId, wxpaybaseDto.AppId)
// 	a.SetValue(*wxPayData, wxConst.MchId, wxpaybaseDto.MchId)
// 	a.SetValue(*wxPayData, wxConst.SubAppId, wxpaybaseDto.SubAppId)
// 	a.SetValue(*wxPayData, wxConst.SubMchId, wxpaybaseDto.SubMchId)

// 	a.SetValue(*wxPayData, wxConst.NonceStr, common.Uuid32())
// 	return *wxPayData
// }

// func (a *WxPayService) SetValue(wxPayData WxPayData, key string, value string) {
// 	if len(strings.TrimSpace(value)) != 0 {
// 		wxPayData.SetValue(key, value)
// 	}
// }
