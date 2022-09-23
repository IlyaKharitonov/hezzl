package natsBroker

import (
	//"context"
	"fmt"
	"log"
	"strings"

	ch "github.com/leprosus/golang-clickhouse"

	"github.com/nats-io/nats.go"
)

type natsReader struct {
	n*nats.Conn
	ch *ch.Conn
}

func NewReader(conn *nats.Conn, chConn *ch.Conn) *natsReader {
	return &natsReader{conn, chConn}
}

func(nr *natsReader)Read()error{
	go func() {
		logs := make([]string, 0, 10)

		_, err := nr.n.Subscribe("log.all", func(m *nats.Msg){
			logs = append(logs, string(m.Data))
			fmt.Println(string(m.Data))
			if len(logs) == 10 {

				err := nr.Insert(logs)
				if err != nil {
					log.Fatalf("crashed nats %s", err.Error())
				}
				logs = make([]string, 0, 10)
			}
		})
		if err != nil {
			log.Fatalf("(nb *natsBroker)Read #2 \nError: %s \n", err.Error())
		}
	}()
	return nil
}

func (nr *natsReader)Insert(logs []string)error {
	var args = make([]string, 0, 10)
	for _,el := range logs{
		args = append(args, fmt.Sprintf("('%s')",el))
	}

	err := nr.ch.Exec(fmt.Sprintf(
		`INSERT INTO logs (message) VALUES %s`,
		strings.Join(args, ",")))
	if err != nil {
		return err
	}
	fmt.Println("Записал в кликхаус")

	return nil

}