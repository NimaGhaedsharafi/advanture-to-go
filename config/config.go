package config

type DBConfig struct {
	Name string
	Host string
	Port string
	User string
	Pass string
}

var PSConfig = DBConfig{
	Name: "contacts",
	Host: "127.0.0.1",
	Port: "5432",
	User: "testing",
	Pass: "testing",
}
