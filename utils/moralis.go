package utils

import "trx_tracker/adaptor/model"

func CheckSpamResult(fromAPI *model.TransactionsFromAPI) *model.TransactionFromAPI {
	for i, transaction := range fromAPI.Result {
		if transaction.PossibleSpam == false {
			return &fromAPI.Result[i]
		}
	}
	return nil
}
