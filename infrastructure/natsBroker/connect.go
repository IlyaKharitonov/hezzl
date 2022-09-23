package natsBroker

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func (nb *NatsConfig)Connect()(*nats.Conn, error){
	addr := nats.DefaultURL
	if nb.Host != "" && nb.Port != ""{
		addr = fmt.Sprintf("%s:%s", nb.Host, nb.Port)
	}
	conn, err := nats.Connect(addr)
	if err != nil{
		return nil, err
	}
	return conn, nil
}
