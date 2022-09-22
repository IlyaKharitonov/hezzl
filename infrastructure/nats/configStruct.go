package nats

type NatsConfig struct {
	//Conn *nats.Conn
	Host string `json:"host"`
	Port string `json:"port"`
}
