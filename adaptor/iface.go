package adaptor

import "trx_tracker/adaptor/model"

type OnChainTracking interface {
	GetTransaction(address string, chain string) (*model.TransactionsFromAPI, error)
}

type NoSQL interface {
	// AddWalletAddress(address string, name string) error
	// ShowExistingWalletAddress() (model.ExistingWalletAddress, error)
	// RemoveWalletAddress(address string) (model.ExistingWalletAddress, error)
	ReturnTrxHashFromWallet(address string) (string, error)
}

type TelegramBot interface {
	SendMessage(message string)
}

type Cache interface {
	Get(walletAddress string) (map[string]string, bool)
	CompareCacheValue(val map[string]string, symbol string, hash string) bool
	UpdateTransaction(val map[string]string, walletAddress string, symbol string, hash string)
}
