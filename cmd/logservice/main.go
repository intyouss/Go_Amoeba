package main

import (
	"Go_Amoeba/log"
	"Go_Amoeba/registry"
	"Go_Amoeba/service"
	"context"
	"fmt"
	stdlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "14000"
	serviceAddress := "http://" + host + ":" + port
	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL:     serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service.")
}
