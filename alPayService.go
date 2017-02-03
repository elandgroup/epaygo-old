package epaygo

import (
	"encoding/json"
	"epaygo/al"
	"epaygo/helper/cryptoHelper"
	"epaygo/helper/mapHelper"
	"fmt"
	"strings"
	"time"

	"github.com/smallnest/goreq"
)

type AlPayService struct {
}

func (a *AlPayService) DirectPay(params map[string]string) (result string, err error) {
	payData := *a.BuildCommonparam(params, al.PayCustom)

	bizContent := make(map[string]string)
	a.SetValue(&bizContent, al.Scence, al.BarCodeScenceCustom)
	a.SetValue(&bizContent, al.OutTradeNo, params[al.OutTradeNoMap])
	a.SetValue(&bizContent, al.AuthCode, params[al.AuthCodeMap])
	a.SetValue(&bizContent, al.TotalAmount, params[al.TotalAmountMap])
	a.SetValue(&bizContent, al.Subject, params[al.SubjectMap])

	a.SetValue(&bizContent, al.StoreId, params[al.StoreIdMap])
	a.SetValue(&bizContent, al.SellerId, params[al.SellerIdMap])
	a.SetValue(&bizContent, al.TimeExpire, params[al.TimeExpireMap])
	a.SetValue(&bizContent, al.ExtendParams, params[al.ExtendParamsMap])

	b, _ := json.Marshal(bizContent)
	payData[al.BizContent] = string(b)

	payData[al.Sign], _ = cryptoHelper.GetSha1Hash(mapHelper.SortedUrl(&payData), params[al.SellerPrivateKey])
	p, _ := json.Marshal(payData)
	_, body, _ := goreq.New().Post(al.OpenApi).ContentType("form").SendMapString(string(p)).End()
	//return a.ParseResponse(body, params[al.AliPublicKey], al.PayNode)
	return body, nil
}

func (a *AlPayService) BuildCommonparam(commonParams map[string]string, method string) *map[string]string {
	payData := make(map[string]string)
	a.SetValue(&payData, al.AppId, commonParams[al.AppIdMap])
	a.SetValue(&payData, al.Charset, al.CharsetCustom)
	a.SetValue(&payData, al.Method, method)
	a.SetValue(&payData, al.SignType, al.SignTypeCustom)
	t := time.Now()
	a.SetValue(&payData, al.TimeStamp, fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))

	a.SetValue(&payData, al.Version, al.VersionCustom)
	a.SetValue(&payData, al.SignType, al.SignTypeCustom)

	return &payData
}

func (a *AlPayService) SetValue(mapData *map[string]string, key string, value string) {
	if len(strings.TrimSpace(value)) != 0 {
		(*mapData)[key] = value
	}
}
