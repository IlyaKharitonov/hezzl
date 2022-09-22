package itemService

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type natsBroker struct {
	conn *nats.Conn
}

func NewNatsBroker(conn *nats.Conn) *natsBroker {
	return &natsBroker{conn: conn}
}

func (nb *natsBroker) Send(log string) {

}

func (nb *natsBroker) Read(stat statistic) {
	ec, err := nats.NewEncodedConn(nb.conn, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("(nb *natsBroker)Read #1 \nError: %s \n", err.Error())
	}
	defer ec.Close()

	logs := make([]string, 0, 10)
	chanStr := make(chan string)
	ec.BindRecvChan("log_all", chanStr)

	for {
		req := <-chanStr
		logs = append(logs, req)
		if len(logs) == 10 {

			err := stat.Insert(logs)
			if err != nil {
				log.Fatalf("crashed nats %s", err.Error())
			}
			fmt.Println("запись в кликхаус")
			logs = make([]string, 0, 10)
		}
	}
}
