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

	v1 := e.Group("/v1")

	//wx := v1.Group("/wx")
	al := v1.Group("/al")

	// POST  HTTP/1.1
	// Host: /api/v1/wx/pay
	// Content-Type: application/json
	// {
	//     "AppId":"wx856df5e42a345096",
	//     "Body":"go test",
	//     "MchId":"1293941701",
	//     "TotalFee":"1",
	//     "key":"7911b23f732e60955b297f2021a2492f",
	//     "SpbillCreateIp":"10.202.101.43",
	//     "AuthCode":"130600782533911801"

	// }
	al.POST("/pay", DirectPay)
	// POST /api/v1/wx/query HTTP/1.1
	// Host: localhost:1001
	// Content-Type: application/json
	// {
	//     "AppId":"wx856df5e42a345096",
	//     "MchId":"1293941701",
	//     "key":"7911b23f732e60955b297f2021a2492f",
	//     "SpbillCreateIp":"10.202.101.43",
	//     "OutTradeNo":"82acc3e2e8c043528af963f72b873171"

	// }
	//wx.POST("/query", api.OrderQueryWx)
	// 	POST /api/v1/wx/refund HTTP/1.1
	// Host: localhost:1001
	// Content-Type: application/json
	// Cache-Control: no-cache
	// {
	//     "AppId":"wx856df5e42a345096",
	//     "MchId":"1293941701",
	//     "TotalFee":"1",
	//     "Key":"7911b23f732e60955b297f2021a2492f",
	//     "RefundFee":"1",
	//     "SpbillCreateIp":"10.202.101.43",

	//     "OutTradeNo":"82acc3e2e8c043528af963f72b873171",
	//         "OpUserId":"1293941701",

	//     "CertName":".\\cert\\apiclient_cert.pem",
	//     "CertKey":".\\cert\\apiclient_key.pem",
	//     "RootCa":".\\cert\\rootca.pem"

	// }
	//wx.POST("/refund", api.RefundWx)
	// POST /api/v1/wx/reverse HTTP/1.1
	// Host: localhost:1001
	// Content-Type: application/json

	// {
	//     "AppId":"wx856df5e42a345096",
	//     "MchId":"1293941701",
	//     "key":"7911b23f732e60955b297f2021a2492f",
	//     "SpbillCreateIp":"10.202.101.43",
	//     "OutTradeNo":"574c6284140441168d6dba726362738d",
	//     "CertName":".\\cert\\apiclient_cert.pem",
	//     "CertKey":".\\cert\\apiclient_key.pem",
	//     "RootCa":".\\cert\\rootca.pem"

	// }
	//wx.POST("/reverse", api.ReverseWx)

	// POST /api/v1/al/pay HTTP/1.1
	// Host: localhost:1001
	// Content-Type: application/json
	// {
	//     "AppId":"2015081700219294",
	//      "SellerPrivateKey":"MIICXQIBAAKBgQC7PyjMEuniN6BPn8oqzIZ6AO1NjSTO9R3adCCIwKfKIEoWXXM+tHDpktdPKSaAsWJPTNAGvEvtxOfzXib/EMXKqD0eUy5MatfpRjRdf1hJVimmfrb09Qx2j7CsKLy7nD23m4xubdYBwvkjMwt/L3JxB5D6qryW1wei/j1c+/OCxQIDAQABAoGAT7vGYJgRNf4f6qgNS4pKHTu10RcwPFyOOM7IZ9M5380+HyXuBB6MEjowKwpH1fcy+LepwaR+5KG7b5uBGY4H2ticMtdysBd9gLwnY4Eh4j7LCWE54HvELpeWXkWpFQdb/NQhcqMAGwYsTnRPdBqkrUmJBTYqEGkIlqCQ5vUJOCECQQDhe0KGmbq1RWp6TDvgpA2dUmlt2fdP8oNW8O7MvbDaQRduoZnVRTPYCDKfzFqpNXL1hAYgth1N0vzDnv3VoLcpAkEA1JcY+rLv5js1g5Luv8LaI5/3uOg0CW7fmh/LfGuz8k/OxASN+cAOUjPHrxtc5xn1zat4/bnV5GEdlOp/DhquPQJBAIV2Fsdi4M+AueiPjPWHRQO0jvDVjfwFOFZSn5YSRUa6NmtmPY6tumUJXSWWqKb1GwlVTuc3xBqXYsNLLUWwLhkCQQDJUJCiD0LohhdGEqUuSKnj5H9kxddJO4pZXFSI7UEJbJQDwcBkyn+FTm2BH+tZGZdQfVnlA89OJr0poOpSg+eNAkAKY85SR9KASaTiDBoPpJ8N805XEhd0Kq+ghzSThxL3fVtKUQLiCh7Yd8oMd/G5S3xWJHUXSioATT8uPRH2bOb/",
	//  	"AliPublicKey":"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB",
	//     "Subject":"go test",
	//     "TotalAmount":"0.01",
	//     "AuthCode":"282715674991437167"

	// }
	//al.POST("/pay", api.DirectPayAl)

	// POST /api/v1/al/query HTTP/1.1
	// Host: localhost:1001
	// Content-Type: application/json
	// {
	// 	"AppId":"2015081700219294",
	// 	"SellerPrivateKey":"MIICXQIBAAKBgQC7PyjMEuniN6BPn8oqzIZ6AO1NjSTO9R3adCCIwKfKIEoWXXM+tHDpktdPKSaAsWJPTNAGvEvtxOfzXib/EMXKqD0eUy5MatfpRjRdf1hJVimmfrb09Qx2j7CsKLy7nD23m4xubdYBwvkjMwt/L3JxB5D6qryW1wei/j1c+/OCxQIDAQABAoGAT7vGYJgRNf4f6qgNS4pKHTu10RcwPFyOOM7IZ9M5380+HyXuBB6MEjowKwpH1fcy+LepwaR+5KG7b5uBGY4H2ticMtdysBd9gLwnY4Eh4j7LCWE54HvELpeWXkWpFQdb/NQhcqMAGwYsTnRPdBqkrUmJBTYqEGkIlqCQ5vUJOCECQQDhe0KGmbq1RWp6TDvgpA2dUmlt2fdP8oNW8O7MvbDaQRduoZnVRTPYCDKfzFqpNXL1hAYgth1N0vzDnv3VoLcpAkEA1JcY+rLv5js1g5Luv8LaI5/3uOg0CW7fmh/LfGuz8k/OxASN+cAOUjPHrxtc5xn1zat4/bnV5GEdlOp/DhquPQJBAIV2Fsdi4M+AueiPjPWHRQO0jvDVjfwFOFZSn5YSRUa6NmtmPY6tumUJXSWWqKb1GwlVTuc3xBqXYsNLLUWwLhkCQQDJUJCiD0LohhdGEqUuSKnj5H9kxddJO4pZXFSI7UEJbJQDwcBkyn+FTm2BH+tZGZdQfVnlA89OJr0poOpSg+eNAkAKY85SR9KASaTiDBoPpJ8N805XEhd0Kq+ghzSThxL3fVtKUQLiCh7Yd8oMd/G5S3xWJHUXSioATT8uPRH2bOb/",
	// 	"AliPublicKey":"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB",
	// 	"OutTradeNo":"16e6354521094981aecfba7123685855"
	// }
	//al.POST("/query", api.OrderQueryAl)

	// POST /api/v1/al/refund HTTP/1.1
	// Host: localhost:1001
	// Content-Type: application/json
	// {
	// 	"AppId":"2015081700219294",
	// 	"SellerPrivateKey":"MIICXQIBAAKBgQC7PyjMEuniN6BPn8oqzIZ6AO1NjSTO9R3adCCIwKfKIEoWXXM+tHDpktdPKSaAsWJPTNAGvEvtxOfzXib/EMXKqD0eUy5MatfpRjRdf1hJVimmfrb09Qx2j7CsKLy7nD23m4xubdYBwvkjMwt/L3JxB5D6qryW1wei/j1c+/OCxQIDAQABAoGAT7vGYJgRNf4f6qgNS4pKHTu10RcwPFyOOM7IZ9M5380+HyXuBB6MEjowKwpH1fcy+LepwaR+5KG7b5uBGY4H2ticMtdysBd9gLwnY4Eh4j7LCWE54HvELpeWXkWpFQdb/NQhcqMAGwYsTnRPdBqkrUmJBTYqEGkIlqCQ5vUJOCECQQDhe0KGmbq1RWp6TDvgpA2dUmlt2fdP8oNW8O7MvbDaQRduoZnVRTPYCDKfzFqpNXL1hAYgth1N0vzDnv3VoLcpAkEA1JcY+rLv5js1g5Luv8LaI5/3uOg0CW7fmh/LfGuz8k/OxASN+cAOUjPHrxtc5xn1zat4/bnV5GEdlOp/DhquPQJBAIV2Fsdi4M+AueiPjPWHRQO0jvDVjfwFOFZSn5YSRUa6NmtmPY6tumUJXSWWqKb1GwlVTuc3xBqXYsNLLUWwLhkCQQDJUJCiD0LohhdGEqUuSKnj5H9kxddJO4pZXFSI7UEJbJQDwcBkyn+FTm2BH+tZGZdQfVnlA89OJr0poOpSg+eNAkAKY85SR9KASaTiDBoPpJ8N805XEhd0Kq+ghzSThxL3fVtKUQLiCh7Yd8oMd/G5S3xWJHUXSioATT8uPRH2bOb/",
	// 	"AliPublicKey":"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB",
	// 	"OutTradeNo":"16e6354521094981aecfba7123685855",
	// 	"RefundAmount":"0.01"
	// }
	//al.POST("/refund", api.RefundAl)

	// POST /api/v1/al/cancel HTTP/1.1
	// Host: localhost:1001
	// Content-Type: application/json
	// {
	//     "AppId":"2015081700219294",
	//      "SellerPrivateKey":"MIICXQIBAAKBgQC7PyjMEuniN6BPn8oqzIZ6AO1NjSTO9R3adCCIwKfKIEoWXXM+tHDpktdPKSaAsWJPTNAGvEvtxOfzXib/EMXKqD0eUy5MatfpRjRdf1hJVimmfrb09Qx2j7CsKLy7nD23m4xubdYBwvkjMwt/L3JxB5D6qryW1wei/j1c+/OCxQIDAQABAoGAT7vGYJgRNf4f6qgNS4pKHTu10RcwPFyOOM7IZ9M5380+HyXuBB6MEjowKwpH1fcy+LepwaR+5KG7b5uBGY4H2ticMtdysBd9gLwnY4Eh4j7LCWE54HvELpeWXkWpFQdb/NQhcqMAGwYsTnRPdBqkrUmJBTYqEGkIlqCQ5vUJOCECQQDhe0KGmbq1RWp6TDvgpA2dUmlt2fdP8oNW8O7MvbDaQRduoZnVRTPYCDKfzFqpNXL1hAYgth1N0vzDnv3VoLcpAkEA1JcY+rLv5js1g5Luv8LaI5/3uOg0CW7fmh/LfGuz8k/OxASN+cAOUjPHrxtc5xn1zat4/bnV5GEdlOp/DhquPQJBAIV2Fsdi4M+AueiPjPWHRQO0jvDVjfwFOFZSn5YSRUa6NmtmPY6tumUJXSWWqKb1GwlVTuc3xBqXYsNLLUWwLhkCQQDJUJCiD0LohhdGEqUuSKnj5H9kxddJO4pZXFSI7UEJbJQDwcBkyn+FTm2BH+tZGZdQfVnlA89OJr0poOpSg+eNAkAKY85SR9KASaTiDBoPpJ8N805XEhd0Kq+ghzSThxL3fVtKUQLiCh7Yd8oMd/G5S3xWJHUXSioATT8uPRH2bOb/",
	//  	"AliPublicKey":"MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDDI6d306Q8fIfCOaTXyiUeJHkrIvYISRcc73s3vF1ZT7XN8RNPwJxo8pWaJMmvyTn9N4HQ632qJBVHf8sxHi/fEsraprwCtzvzQETrNRwVxLO5jVmRGi60j8Ue1efIlzPXV9je9mkjzOmdssymZkh2QhUrCmZYI/FCEa3/cNMW0QIDAQAB",
	//     "OutTradeNo":"16e6354521094981aecfba7123685855"
	// }
	//al.POST("/reverse", api.ReverseAl)

}
