package main

import (
	"github.com/peter9207/unischeme/server"
	"github.com/spf13/viper"
)

func main() {

	viper.AutomaticEnv()
	s := server.New(viper.GetString("NAME"), viper.GetString("URL"))
	s.Start()
}
