package main

import (
	"LabUser/log"
	"LabUser/registry"
	"LabUser/service"
	"context"
	"fmt"
	stdlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  "http://" + host + ":" + port,
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
