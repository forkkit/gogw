package main

import ( 
	"flag"

	"gogw/logger"
	"gogw/config"
	"gogw/server"
	"gogw/client"
)

var cfgFile = flag.String("c", "cfg.json", "config file")
var role = flag.String("r", "server", "role: server/client")
var logLevel = flag.String("l", "info", "log level: info/debug")

func main(){
	logger.LEVEL = logger.INFO

	logger.Info("gogw start")
	flag.Parse()

	cfg, err := config.NewConfig(*cfgFile)
	if err != nil {
		logger.Error(err)
		return
	}

	if *logLevel == "debug" {
		logger.LEVEL = logger.DEBUG
	}

	if *role == "server" {
		server := server.NewServer(cfg.Server.ServerAddr, cfg.Server.TimeoutSecond)
		server.Start()
	}

	if *role == "client" {
		client := client.NewClient(
			cfg.Client.ServerAddr, 
			cfg.Client.SourceAddr, 
			cfg.Client.ToPort, 
			cfg.Client.Direction, 
			cfg.Client.Protocol, 
			cfg.Client.Description,
		)
		client.Start()
	}
}