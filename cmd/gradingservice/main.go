package main

import (
	"Go_Amoeba/grades"
	"Go_Amoeba/log"
	"Go_Amoeba/registry"
	"Go_Amoeba/service"
	"context"
	"fmt"
	stdlog "log"
)

func main() {
	host, port := "localhost", "16000"
	serviceAddress := "http://" + host + ":" + port
	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL:     serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatalln(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()

	fmt.Println("Shutting down grading service.")
}
