package main

import (
	//"context"
	"flag"
	"log"

	tgClient "bot/clients/telegram"
	event_consumer "bot/consumer/event-consumer"
	"bot/events/telegram"

	"bot/storage/files"
	//"bot/storage/sqlite"
)

const (
	tgBotHost         = "api.telegram.org"
	storagePath       = "files_storage"
	sqlitestoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	s := files.New(storagePath)
	/*s, err := sqlite.New(sqlitestoragePath)
	if err != nil {
		log.Fatal("can't connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can't init storage: ", err)
	}*/

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",                     //имя
		"",                                 //значение
		"token for access to telegram bot", //описание
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
