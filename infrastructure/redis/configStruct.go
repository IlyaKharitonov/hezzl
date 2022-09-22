package redis

type RedisConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
	NameDB   int    `json:"name"`
}
