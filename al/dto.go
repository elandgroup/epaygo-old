package al

type AlPayBaseDto struct {
	AppId            string
	SellerPrivateKey string
	AliPublicKey     string
	OutTradeNo       string
	ALAuthToken      string
}

type AlDirectPayDto struct {
	AlPayBaseDto
	AuthCode    string
	TotalAmount string
	Subject     string
	StoreId     string

	SellerId     string
	TimeExpire   string
	ExtendParams string
}
