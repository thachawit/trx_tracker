package constant

// var (
// 	Chains = [7]string{"eth", "polygon", "bsc", "avalanche", "arbitrum", "base", "optimism"}
// 	// GetTransactionFromWalletURL = "https://deep-index.moralis.io/api/v2.2/wallets/"
// )

type MoralisConstant struct {
	Address [3]string
	Chains  [7]string
}

func Chains() *MoralisConstant {
	return &MoralisConstant{
		Address: [3]string{"0x53c58e975f3f72162fc0509b9742c9b55e24a599", "0x0d0707963952f2fba59dd06f2b425ace40b492fe", "0x46f4a6bf81a41d2b5ada059d43b959b8e6068fe2"},
		Chains:  [7]string{"eth", "polygon", "bsc", "avalanche", "arbitrum", "base", "optimism"}}
}
