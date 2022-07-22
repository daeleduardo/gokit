package config

import (
	"sync"
)

//Application Configuration properties.
type Config struct {
	DbConn	string
	CacheConn string
	GinMode string
	Port int
	JwtToken string
	JwtTimeout int64
}

func DefaultConfig() Config {
	return Config  {
		DbConn	: "user:password@tcp(localhost:3306)/database",
		CacheConn : "username:password@localhost:6379/0",
		GinMode :	"debug",
		Port :	8080,
		JwtToken :	"UltraMegaPowerSecretKey",
		JwtTimeout : 3600000000000, /*1 Hour*/
	}
}

var conf *Config

var lock = &sync.Mutex{}

//Return a singleton config struct
func GetConfig () *Config {
	lock.Lock()
	defer lock.Unlock()
	if conf == nil {
		DefaultConf := DefaultConfig()
		conf = &DefaultConf
	}
	return conf
}

func (c Config) Init()  {
	conf = &c
}
