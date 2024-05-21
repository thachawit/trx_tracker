package utils

import (
	"fmt"
	"time"
	m3 "trx_tracker/adaptor/model"
	"trx_tracker/model"
)

func blockExplorer(chain string) string {
	explorer := map[string]string{
		"eth":       "https://etherscan.io/tx/",
		"polygon":   "https://polygonscan.com/tx/",
		"bsc":       "https://bscscan.com/tx/",
		"avalanche": "https://explorer.avax.network/search",
		"arbitrum":  "https://arbiscan.io/tx/",
		"base":      "https://basescan.org/tx/",
		"optimism":  "https://optimistic.etherscan.io/tx/",
	}
	return explorer[chain]

}

func SendGetTransactionMsg(transaction *m3.TransactionFromAPI, wallet string, chain string) string {
	layout := "2006-01-02T15:04:05"
	msgSetup := &model.SeeNewTransactionMessage{
		Summary:       transaction.Summary,
		Timestamp:     transaction.BlockTimestamp.Add(7 * time.Hour).Format(layout),
		TokenSymbol:   transaction.ERC20Transfers[0].TokenSymbol,
		Chain:         chain,
		BlockExplorer: blockExplorer(chain) + transaction.Hash,
	}
	msg := fmt.Sprintf("%s\n"+
		"%s\n"+
		"%s\n"+
		"Chain: %s\n"+
		"%s\n", msgSetup.Summary, msgSetup.Timestamp, msgSetup.TokenSymbol, msgSetup.Chain, msgSetup.BlockExplorer)

	return msg
}
