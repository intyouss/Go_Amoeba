package main

import (
	"Go_Amoeba/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	registry.MonitorHeartbeat()
	// 服务注册地址
	http.Handle("/services", &registry.RegistryService{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Refistry service started.Press any key to stop")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()

	fmt.Println("Shutting down registry service.")
}
