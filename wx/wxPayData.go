package wx

// import (
// 	"encoding/xml"
// 	"epaygo/helper"

// 	"strings"

// 	"encoding/json"

// 	"errors"
// )

// type WxPayData struct {
// 	xmlMap helper.XmlMap
// }

// func NewWxPayData() *WxPayData {
// 	wxPayData := WxPayData{}
// 	wxPayData.xmlMap = make(XmlMap)
// 	return &wxPayData
// }

// func (w *WxPayData) SetValue(key string, value string) {
// 	w.xmlMap[key] = value
// }

// func (w *WxPayData) GetValue(key string) string {
// 	return w.xmlMap[key]
// }
// func (w *WxPayData) RemoveKey(key string) {
// 	delete(w.xmlMap, key)
// }

// func (w *WxPayData) IsSet(key string) bool {
// 	_, ok := w.xmlMap[key]
// 	return !ok
// }

// func (w *WxPayData) ToXml() string {
// 	if len(w.xmlMap) == 0 {
// 		return ""
// 	}
// 	if x, err := xml.MarshalIndent(w.xmlMap, "", " "); err == nil {
// 		return string(x)
// 	}
// 	return ""
// }

// func (w *WxPayData) ToUrl() string {
// 	return SortedUrlWX(&(w.xmlMap))
// }
// func (w *WxPayData) ToJson() string {

// 	v, _ := json.Marshal(w.xmlMap)

// 	return string(v)
// }

// func (w *WxPayData) MakeSign(key string) string {
// 	str := w.ToUrl()
// 	str += "&key=" + key
// 	return GetMD5Hash(str)
// }

// func (w *WxPayData) CheckSign(key string) (isSign bool, err error) {
// 	if !w.IsSet(Sign) {
// 		return false, errors.New("10001")
// 	} else if len(strings.TrimSpace(w.GetValue(Sign))) == 0 {
// 		return false, errors.New("10002")
// 	} else if w.GetValue(Sign) == w.MakeSign(key) {
// 		return true, nil
// 	}
// 	return false, errors.New("10002")

// }

// func (w *WxPayData) FromXml(xml string, key string) (xmlMap XmlMap, errKey error) {

// 	if xmlMap, err := XmlToWxMap(xml); err == nil {
// 		w.xmlMap = xmlMap
// 	} else {
// 		return nil, err
// 	}
// 	return w.xmlMap, nil

// }
