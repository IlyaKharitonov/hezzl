package databases

//for pg and ch
type DBConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	DBName   string `json:"dbName"`
	SSLmode  string `json:"sslMode"`
	Debug    string `json:"debug"`
}
