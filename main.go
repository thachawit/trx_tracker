package main

import (
	"log"
	"strings"
	"time"
	"trx_tracker/adaptor"
	"trx_tracker/scheduler"

	"github.com/go-co-op/gocron/v2"
	"github.com/go-resty/resty/v2"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	tele "gopkg.in/telebot.v3"
)

func initViper() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func initTelebot() *tele.Bot {
	token := viper.GetString("telegram.api-key")
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return bot

}

func initGoCron() gocron.Scheduler {
	schedule, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	return schedule
}

func gracefullyShutdown() {

}

func initCache() *cache.Cache {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	return c
}

func main() {
	initViper()
	bot := initTelebot()
	schedule := initGoCron()

	rest := resty.New()
	cache := adaptor.NewCacheAdaptor()
	moralis := adaptor.NewMoralisAdaptor(rest)
	mongoDB := adaptor.NewNoSQLAdaptor(&mongo.Client{}) //
	telegramBot := adaptor.NewTelegramBotAdaptor(bot)
	s := scheduler.NewSchedulerService(cache, moralis, mongoDB, telegramBot)

	log.Println("Service is up and running.")

	job, err := schedule.NewJob(
		gocron.DurationJob(
			60*time.Minute,
		),
		gocron.NewTask(
			func() {
				s.GetTransaction()
			},
		),
	)

	log.Println(job.ID())

	if err != nil {
		log.Fatal(err)
	}
	schedule.Start()

	defer bot.Start()

}
