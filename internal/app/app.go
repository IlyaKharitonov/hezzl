package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"hezzlTestTask/config"
	"hezzlTestTask/internal/controllers"
)

func Start(c *config.ConfigJSON)error{
	ctx := context.Background()

	err := c.Postgres.Connect()
	if err != nil {
		fmt.Errorf("app.Start #1 \nError: %s \n", err.Error())
	}
	defer c.Postgres.DB.Close(ctx)
	log.Println("connected with postgres")


	//err = c.ClickHouse.Connect()
	//if err != nil {
	//	fmt.Errorf("app.Start #2 \nError: %s \n", err.Error())
	//}
	//defer c.ClickHouse.DB.Close()
	//log.Println("connected with clickhouse")

	err = c.Redis.Connect()
	if err != nil {
		fmt.Errorf("app.Start #2 \nError: %s \n", err.Error())
	}
	log.Println("connected with redis")

	controllers.HandlersRegister()
	log.Println("handlers registration done")

	log.Println("start server")
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port), nil)
	if err != nil {
		log.Printf("server error: %s", err)
	}

	return err
}
