package model

type Wallet struct {
	Name    string
	Address string
}

type ExistingWalletAddress struct {
	Wallets []Wallet
}

type TransactionDetail struct {
	Token        string `json:"token" bson:"token"`
	Amount       string `json:"amount" bson:"amount"`
	Chain        string `json:"chain" bson:"chain"`
	From         string `json:"from" bson:"from"`
	FromAddress  string `json:"from_address" bson:"from_address"`
	To           string `json:"to" bson:"to"`
	ToAddress    string `json:"to_address" bson:"to_address"`
	TokenAddress string `json:"token_address" bson:"token_address"`
	TokenScanner string `json:"token_scanner" bson:"token_scanner"`
}

type Transaction struct {
	Name              string            `json:"name,omitempty" bson:"name"`
	WalletAddress     string            `json:"wallet_address,omitempty" bson:"wallet_address"`
	Chain             string            `json:"chain,omitempty" bson:"chain"`
	Scanner           string            `json:"scanner,omitempty" bson:"scanner"`
	TransactionDetail TransactionDetail `json:"transaction_detail,omitempty" bson:"transaction_detail"`
	CurrentHash       string            `json:"current_hash,omitempty" bson:"current_hash"`
}
