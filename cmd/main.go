package main

import (
	"flag"
	"fmt"

	"github.com/HazeyamaLab/go-crud/conf"
	"github.com/HazeyamaLab/go-crud/pkg/server"
)

var state = flag.String("s", "local", "local/prd")

func main() {

	flag.Parse()
	err := conf.SetUp(fmt.Sprintf("conf/env/%s.toml", *state))

	if err != nil{
		fmt.Printf("%s", err)
		return
	}

	conf.Init()
	server.Init()
}
