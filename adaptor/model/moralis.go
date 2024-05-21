package model

import "time"

type ERC20Transfer struct {
	TokenName        string `json:"token_name"`
	TokenSymbol      string `json:"token_symbol"`
	TokenLogo        string `json:"token_logo"`
	TokenDecimals    string `json:"token_decimals"`
	FromAddress      string `json:"from_address"`
	FromAddressLabel string `json:"from_address_label"`
	ToAddress        string `json:"to_address"`
	ToAddressLabel   string `json:"to_address_label"`
	Address          string `json:"address"`
	LogIndex         int    `json:"log_index"`
	Value            string `json:"value"`
	PossibleSpam     bool   `json:"possible_spam"`
	VerifiedContract bool   `json:"verified_contract"`
	Direction        string `json:"direction"`
	ValueFormatted   string `json:"value_formatted"`
}

// Transaction represents a transaction
type TransactionFromAPI struct {
	Hash                     string          `json:"hash"`
	Nonce                    string          `json:"nonce"`
	TransactionIndex         string          `json:"transaction_index"`
	FromAddress              string          `json:"from_address"`
	FromAddressLabel         interface{}     `json:"from_address_label"`
	ToAddress                string          `json:"to_address"`
	ToAddressLabel           interface{}     `json:"to_address_label"`
	Value                    string          `json:"value"`
	Gas                      string          `json:"gas"`
	GasPrice                 string          `json:"gas_price"`
	Input                    string          `json:"input"`
	ReceiptCumulativeGasUsed string          `json:"receipt_cumulative_gas_used"`
	ReceiptGasUsed           string          `json:"receipt_gas_used"`
	ReceiptStatus            string          `json:"receipt_status"`
	BlockTimestamp           time.Time       `json:"block_timestamp"`
	BlockNumber              string          `json:"block_number"`
	BlockHash                string          `json:"block_hash"`
	NFTTransfers             []interface{}   `json:"nft_transfers"`
	ERC20Transfers           []ERC20Transfer `json:"erc20_transfers"`
	MethodLabel              interface{}     `json:"method_label"`
	NativeTransfers          []interface{}   `json:"native_transfers"`
	Summary                  string          `json:"summary"`
	PossibleSpam             bool            `json:"possible_spam"`
	Category                 string          `json:"category"`
}
type TransactionsFromAPI struct {
	Result []TransactionFromAPI
}

// "result": [
//     {
//       "hash": "0x1fd0e8111346998ce22918f7bdfc34aa3116eae2e72eaf3e7d3a39f077ced4ac",
//       "nonce": "737",
//       "transaction_index": "76",
//       "from_address": "0x4890ae58845cee70b04f7b54dd2d4bbfd8b697d4",
//       "from_address_label": null,
//       "to_address": "0xe0d0cf872c09968f1554730042ee680354dec994",
//       "to_address_label": null,
//       "value": "0",
//       "gas": "93191",
//       "gas_price": "7360680861",
//       "input": "",
//       "receipt_cumulative_gas_used": "5461471",
//       "receipt_gas_used": "73738",
//       "receipt_status": "1",
//       "block_timestamp": "2024-04-20T10:48:59.000Z",
//       "block_number": "19696135",
//       "block_hash": "0x53ff2fb1116a439b96387ef2983f809eac7fcd67506f6c5e5fb1859801dda5a3",
//       "nft_transfers": [],
//       "erc20_transfers": [
//         {
//           "token_name": "Tether USD",
//           "token_symbol": "USDT",
//           "token_logo": "https://logo.moralis.io/0x1_0xdac17f958d2ee523a2206206994597c13d831ec7_0b0d126af6c744c185e112a2c8dc1495",
//           "token_decimals": "6",
//           "from_address": "0xe0d0cf872c09968f1554730042ee680354dec994",
//           "from_address_label": null,
//           "to_address": "0x53c58e975f3f72162fc0509b9742c9b55e24a599",
//           "to_address_label": null,
//           "address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
//           "log_index": 116,
//           "value": "1300000000",
//           "possible_spam": false,
//           "verified_contract": true,
//           "direction": "receive",
//           "value_formatted": "1300"
//         }
//       ],
//       "method_label": null,
//       "native_transfers": [],
//       "summary": "Received 1,300 USDT from 0xe0...c994",
//       "possible_spam": false,
//       "category": "token receive"
//     },

type CompareFromAPI struct {
	// Name              string               `json:"name,omitempty" bson:"name"`
	// WalletAddress     string               `json:"wallet_address,omitempty" bson:"wallet_address"`
	// Chain             string               `json:"chain,omitempty" bson:"chain"`
	// Scanner           string               `json:"scanner,omitempty" bson:"scanner"`
	// TransactionDetail compareDetailFromAPI `json:"transaction_detail,omitempty" bson:"transaction_detail"`
	CurrentHash string `json:"current_hash,omitempty" bson:"current_hash"`
}
