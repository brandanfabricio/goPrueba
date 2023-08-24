package main

import (
	"api/server"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background() 

	serverDoneChan := make(chan os.Signal,1)
	signal.Notify(serverDoneChan,os.Interrupt,syscall.SIGTERM)
	server := server.New(":3000")
	go server.ListenAndServe()

fmt.Println("Server start port 3000")
<- serverDoneChan
server.Shutdown(ctx)
log.Println("server stop ")

}
