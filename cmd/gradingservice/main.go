package main

import (
	"LabUser/grades"
	"LabUser/registry"
	"LabUser/service"
	"context"
	"fmt"
	stdlog "log"
)

func main() {
	host, port := "localhost", "6000"
	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  "http://" + host + ":" + port,
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
	<-ctx.Done()

	fmt.Println("Shutting down grading service.")
}