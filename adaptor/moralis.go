package adaptor

import (
	"errors"
	"fmt"
	"trx_tracker/adaptor/model"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type MoralisAdaptor struct {
	client *resty.Client
}

func NewMoralisAdaptor(client *resty.Client) OnChainTracking {
	return &MoralisAdaptor{
		client: client,
	}
}

func (a *MoralisAdaptor) GetTransaction(address string, chain string) (*model.TransactionsFromAPI, error) {

	url := fmt.Sprintf("https://deep-index.moralis.io/api/v2.2/wallets/%s/history?chain=%s&order=DESC", address, chain)
	result := model.TransactionsFromAPI{}
	apiKey := viper.GetString("app.api-key")
	fmt.Println(apiKey)

	resp, err := a.client.R().
		SetHeader("Accept", "application/json").
		SetHeader("X-API-Key", apiKey).
		SetResult(&result).
		SetQueryParams(map[string]string{
			"order": "DESC",
		}).
		Get(url)

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New(resp.Status())
	}

	return &result, nil
}
