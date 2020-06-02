package conf


import (
	"log"
	"runtime"

	"github.com/BurntSushi/toml"
)

var conf Config


type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User  string
	Password string
	IP string
}

func SetUp(file string) error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	_, err := toml.DecodeFile(file, &conf)
	if err != nil{
		log.Println(err)
		return err
	}

	return nil
}