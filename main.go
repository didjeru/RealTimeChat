package main

import (
	"fmt"
	"goPat/realtimechat/server"
	"log"
)

func main() {

	serv := server.New()
	fmt.Println("Realtimechat server is running ...")
	if err := serv.Start(); err != nil {
		log.Fatalf("start error: %v", err)
	}

}
