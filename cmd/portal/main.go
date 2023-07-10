package main

import (
	"Go_Amoeba/log"
	"Go_Amoeba/portal"
	"Go_Amoeba/registry"
	"Go_Amoeba/service"
	"context"
	"fmt"
	stdlog "log"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stdlog.Fatal(err)
	}
	host, port := "localhost", "15000"
	serviceAddress := "http://" + host + ":" + port
	r := registry.Registration{
		ServiceName: registry.PortalService,
		ServiceURL:  serviceAddress,
		RequiredServices: []registry.ServiceName{
			registry.GradingService,
			registry.LogService,
		},
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL:     serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		portal.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	if GradingProvider, err := registry.GetProvider(registry.GradingService); err == nil {
		fmt.Printf("Grading service found at: %s\n", GradingProvider)
	}
	<-ctx.Done()
	fmt.Println("Shutting down portal.")
}
