package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"

	"hezzlTestTask/internal/controllers/itemController"
	"hezzlTestTask/internal/services/itemService"
)

func Start(c *ConfigJSON) error {

	postgres, err := c.Postgres.ConnectPostgres(context.Background())
	if err != nil {
		log.Fatalf("Start #1 \nError: %s \n", err.Error())
	}
	log.Println("postgres connected")

	//clickHouse, err := c.ClickHouse.ConnectClickHous()
	//if err != nil{
	//	log.Fatalf("Start #2 \nError: %s \n", err.Error())
	//}

	redis, err := c.Redis.Connect()
	if err != nil {
		log.Fatalf("Start #2 \nError: %s \n", err.Error())
	}
	log.Println("redis connected")

	nats, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Start #3 \nError: %s \n", err.Error())
	}
	log.Println("nats connected")
	broker := itemService.NewNatsBroker(nats)

	item := itemService.NewItem(
		itemService.NewPostgresDB(postgres),
		itemService.NewRedisCache(redis),
		broker,
		itemService.NewClickHouseDB(nil))

	itemController.HandlersRegister(itemController.NewController(item))
	log.Println("item handlers registration done")

	go func() {
		broker.Read()

	}()

	log.Println("start server")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port), nil)
	if err != nil {
		log.Printf("server error: %s", err)
	}

	return err
}
