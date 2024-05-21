package scheduler

import (
	"log"
	"trx_tracker/adaptor"
	"trx_tracker/constant"
	"trx_tracker/utils"
)

type SchedulerService struct {
	cache   adaptor.Cache
	moralis adaptor.OnChainTracking
	mongo   adaptor.NoSQL
	bot     adaptor.TelegramBot
}

func NewSchedulerService(cache adaptor.Cache, moralis adaptor.OnChainTracking, mongo adaptor.NoSQL, bot adaptor.TelegramBot) *SchedulerService {
	return &SchedulerService{
		cache:   cache,
		moralis: moralis,
		mongo:   mongo,
		bot:     bot,
	}
}

func (a *SchedulerService) GetTransaction() error {
	apiDetails := constant.Chains()
	for _, address := range apiDetails.Address {
		// Iterate over each chain
		for _, chain := range apiDetails.Chains {
			log.Println(address, chain)
			//get new transaction
			fromAPI, err := a.moralis.GetTransaction(address, chain)
			if err != nil {
				log.Println("from api: ", err)
				return err
			}
			// means there is no wallet address in this cache

			result := utils.CheckSpamResult(fromAPI)
			if result == nil {
				// done nothing, skip to next chain.
				continue
			}
			if len(result.ERC20Transfers) == 0 {
				continue

			}
			//means the entire existing transaction that has happended are all spam.
			dict, found := a.cache.Get(address)
			if !found {
				a.cache.UpdateTransaction(dict, address, chain, result.Hash)
				a.bot.SendMessage(utils.SendGetTransactionMsg(result, address, chain))
			}
			newTransaction := a.cache.CompareCacheValue(dict, chain, result.Hash)
			if newTransaction == false {
				continue
			}

			//add cache then send message since it is new transaction never existed in cache before add it then send message
			a.cache.UpdateTransaction(dict, address, chain, result.Hash)
			a.bot.SendMessage(utils.SendGetTransactionMsg(result, address, chain))
		}
	}
	// means there is no new transaction happened, send nothing.
	return nil
}
