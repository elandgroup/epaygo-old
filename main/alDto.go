package main

type AlPayBaseDto struct {
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
}

type AlDirectPayDto struct {
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
	AuthCode         string
	TotalAmount      string
	Subject          string
	StoreId          string

	SellerId     string
	TimeExpire   string
	ExtendParams string
}

type AlOrderQueryDto struct {
	AlPayBaseDto
	TradeNo string
}

type AlRefundDto struct {
	AlPayBaseDto
	TradeNo      string
	RefundAmount string
	OutRequestNo string
	RefundReason string
	StoreId      string
}

type AlReverseDto struct {
	AlPayBaseDto
	TradeNo string
}
