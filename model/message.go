package model

import "time"

type SeeNewTransactionMsg struct {
	Summary       string
	WalletAddress string
	TokenName     string
	TokenSymbol   string
	Direction     string
	Amount        string
	Chain         string
	Timestamp     time.Time
	Txscanner     string
}

type SeeNewTransactionMessage struct {
	Summary       string
	Timestamp     string
	TokenSymbol   string `json:"token_symbol"`
	Chain         string
	BlockExplorer string
}

//"https://deep-index.moralis.io/api/v2.2/
//0x1f9090aaE28b8a3dCeaDf281B0F12828e676c326?chain=eth&order=DESC"

//https://deep-index.moralis.io/api/v2.2/:address
