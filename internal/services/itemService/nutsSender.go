package itemService

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type natsSender struct {
	*nats.Conn
}

func NewNatsSender(conn *nats.Conn) *natsSender {
	return &natsSender{conn}
}

func (nb *natsSender) Send(logMes string) error{
	err := nb.Conn.Publish("log.all", []byte(logMes))
	if err != nil {
		return fmt.Errorf("(nb *natsSender) Send #1 \n Error:%s \n", err.Error())
	}


	//ec, err := nats.NewEncodedConn(nb.Conn, nats.JSON_ENCODER)
	//if err != nil {
	//	return fmt.Errorf("(nb *natsSender) Send #1 \n Error:%s \n", err.Error())
	//}
	//defer ec.Close()
	//
	//chanStr := make(chan string)
	//err = ec.BindSendChan("log_all", chanStr)
	//if err != nil {
	//	return fmt.Errorf("(nb *natsSender) Send #2 \n Error:%s \n", err.Error())
	//}
	//
	//chanStr <- logMes

	return nil
}
